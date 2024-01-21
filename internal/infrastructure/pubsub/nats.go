/* Copyright (C) Fedir Petryk */

package pubsub

import (
	"time"

	"github.com/nats-io/nats.go"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewNatsConn(url string, logger *zap.SugaredLogger) (*nats.Conn, error) {
	logger.Debug("connecting to nats")

	maxRetries := 10
	var err error
	var nc *nats.Conn

	for {
		if maxRetries == 0 {
			return nil, errors.New("error connecting to redis: maximum attempts reached")
		}

		nc, err = nats.Connect(url)
		if err != nil {
			logger.Infow("failed to connect to nats", "err", err)
			maxRetries--
			time.Sleep(5 * time.Second)
			continue
		}

		status := nc.IsConnected()
		if !status {
			logger.Infow("failed to connect to nats: not connected")
			maxRetries--
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}

	logger.Debug("connecting to nats done")

	return nc, nil
}
