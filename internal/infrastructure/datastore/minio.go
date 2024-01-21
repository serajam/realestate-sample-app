/* Copyright (C) Fedir Petryk */

package datastore

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func NewMinioConn(ctx context.Context, endpoint, accessKeyID, secretAccessKey string, useSSL bool, logger *zap.SugaredLogger) (
	*minio.Client,
	error,
) {
	minioClient, err := minio.New(
		endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		},
	)
	if err != nil {
		return nil, err
	}

	cancel, err := minioClient.HealthCheck(1 * time.Second)
	defer cancel()
	if err != nil {
		return nil, err
	}

	maxRetries := 10
	for {
		select {
		case <-ctx.Done():
		default:
			if maxRetries == 0 {
				return nil, errors.New("error connecting to redis: maximum attempts reached")
			}

			if !minioClient.IsOnline() {
				logger.Info("minio is offline, retrying")
				maxRetries--
				time.Sleep(5 * time.Second)
				continue
			}

			return minioClient, nil
		}
	}
}
