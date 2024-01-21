/* Copyright (C) Fedir Petryk */

package user

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Subscriber struct {
	subscriptions []*nats.Subscription
	conn          *nats.Conn
	logger        *zap.SugaredLogger

	service EmailService
}

func NewSubscriber(conn *nats.Conn, service EmailService, logger *zap.SugaredLogger) (Subscriber, error) {
	sb := Subscriber{conn: conn, logger: logger, service: service, subscriptions: []*nats.Subscription{}}

	sub, err := conn.Subscribe(
		deactivated, func(msg *nats.Msg) {
			sb.Deactivated(msg.Data)
		},
	)

	if err != nil {
		return Subscriber{}, err
	}

	sb.subscriptions = append(sb.subscriptions, sub)

	sub, err = conn.Subscribe(
		signup, func(msg *nats.Msg) {
			sb.Signup(msg.Data)
		},
	)

	if err != nil {
		return Subscriber{}, err
	}

	sb.subscriptions = append(sb.subscriptions, sub)

	return sb, nil
}

func (s Subscriber) Shutdown() {
	for _, sub := range s.subscriptions {
		err := sub.Drain()
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
}
