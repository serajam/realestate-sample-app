/* Copyright (C) Fedir Petryk */

package datastore

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"runtime"
	"time"

	"github.com/uptrace/bun/extra/bundebug"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var sqlMigrations embed.FS

func NewPostgresConn(ctx context.Context, dsn string, enableQueryDebug bool, logger *zap.SugaredLogger) (*bun.DB, error) {
	logger.Debug("connecting to postgres database")
	sqldb := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
			pgdriver.WithTimeout(5*time.Second),
			pgdriver.WithDialTimeout(1*time.Second),
			pgdriver.WithReadTimeout(5*time.Second),
			pgdriver.WithWriteTimeout(5*time.Second),
		),
	)

	dbConn := bun.NewDB(sqldb, pgdialect.New())
	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)

	maxRetries := 10
	for {
		if maxRetries == 0 {
			return nil, errors.New("error connecting to db: maximum attempts reached")
		}

		ctxTime, cncl := context.WithTimeout(ctx, time.Second*1)
		defer cncl()

		if err := dbConn.PingContext(ctxTime); err != nil {
			logger.Infow("failed to ping database", "err", err)
			maxRetries--
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}

	if enableQueryDebug {
		//	debug dbConn queries
		dbConn.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	logger.Debug("connecting to postgres done")

	return dbConn, nil
}

func ApplyMigrations(ctx context.Context, conn *bun.DB, cleanup bool, logger *zap.SugaredLogger) error {
	migrations := migrate.NewMigrations()
	err := migrations.Discover(sqlMigrations)
	if err != nil {
		return fmt.Errorf("migrations discover: %w", err)
	}

	migrator := migrate.NewMigrator(
		conn,
		migrations,
		migrate.WithMarkAppliedOnSuccess(true),
	)

	if err := migrator.Init(ctx); err != nil {
		return fmt.Errorf("migrator init: %w", err)
	}

	if cleanup {
		logger.Debug("migrations: cleaning up database")
		migrs, err := migrator.AppliedMigrations(ctx)
		if err != nil {
			return fmt.Errorf("migrator rollback: %w", err)
		}

		for range migrs {
			if res, err := migrator.Rollback(ctx); err != nil {
				return fmt.Errorf("migrator rollback: %w, %s", err, res.String())
			}
		}

	}

	res, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migrator migrate: %w, %s", err, res.String())
	}

	if res.IsZero() {
		logger.Info("migrations: there are no new migrations to run (database is up to date)")
		return nil
	}

	logger.Infow(
		"migrations: migrated", "version", res.String(),
	)

	return nil
}
