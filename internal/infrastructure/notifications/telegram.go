/* Copyright (C) Fedir Petryk */

package notifications

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

type TeleBot interface {
	Send(msg string) error
}

type Telebot struct {
	api       *tele.Bot
	channelId int64
}

func NewTelebot(token string, channelId int64) (TeleBot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	teleCl, err := tele.NewBot(pref)
	if err != nil {
		return dummyTelebot{}, err
	}

	return &Telebot{
		api:       teleCl,
		channelId: channelId,
	}, nil
}

func (tb *Telebot) Send(msg string) error {
	chat, err := tb.api.ChatByID(tb.channelId)
	if err != nil {
		return err
	}

	_, err = tb.api.Send(chat, msg, &tele.SendOptions{})
	if err != nil {
		return err
	}

	return nil
}

type dummyTelebot struct {
}

func (tb dummyTelebot) Send(msg string) error {
	return nil
}
