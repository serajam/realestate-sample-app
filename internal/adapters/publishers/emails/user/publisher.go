/* Copyright (C) Fedir Petryk */

package user

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Publisher struct {
	conn   *nats.Conn
	logger *zap.SugaredLogger
}

func NewPublisher(conn *nats.Conn, logger *zap.SugaredLogger) Publisher {
	return Publisher{conn: conn, logger: logger}
}

func (p Publisher) publish(ch string, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		p.logger.Error(err.Error())
		return
	}

	err = p.conn.Publish(ch, bytes)
	if err != nil {
		p.logger.Errorw("publish failed", "err", err.Error())
	}
}
