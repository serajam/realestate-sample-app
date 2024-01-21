/* Copyright (C) Fedir Petryk */

package datastore

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRedisConn(ctx context.Context, addr, pwd string, logger *zap.SugaredLogger) (*redis.Client, error) {
	logger.Debug("connecting to redis database")
	rdb := redis.NewClient(
		&redis.Options{
			Addr:                  addr,
			Password:              pwd,
			DB:                    0, // use default DB
			Protocol:              3, // specify 2 for RESP 2 or 3 for RESP 3
			MaxRetries:            100,
			ContextTimeoutEnabled: true,
		},
	)

	maxRetries := 10
	for {
		if maxRetries == 0 {
			return nil, errors.New("error connecting to redis: maximum attempts reached")
		}

		ctxTime, cncl := context.WithTimeout(ctx, time.Second*1)
		defer cncl()

		status := rdb.Conn().Ping(ctxTime)
		if status.Err() != nil {
			logger.Infow("failed to ping redis", "err", status.Err())
			maxRetries--
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}

	logger.Debug("connecting to redis done")

	return rdb, nil
}
