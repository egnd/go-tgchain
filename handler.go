package tgchain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// EventHandler is event handler interface.
type IEventHandler interface {
	GetName() string
	GetDescription() string
	Decorate(IEventHandler)
	Next(context.Context, tgbotapi.Update) error
	Handle(context.Context, tgbotapi.Update) error
}

type AbstractHandler struct {
	next IEventHandler
}

func (h *AbstractHandler) GetName() string { return "" }

func (h *AbstractHandler) GetDescription() (descr string) { return "" }

func (h *AbstractHandler) Decorate(next IEventHandler) {
	h.next = next
}

func (h *AbstractHandler) Next(ctx context.Context, upd tgbotapi.Update) error {
	if h.next == nil {
		return nil
	}

	return h.next.Handle(ctx, upd)
}

func (h *AbstractHandler) ReplyToMsg(msg *tgbotapi.Message, text string, api ITgAPI) (err error) {
	if msg == nil || len(text) == 0 {
		return
	}

	resp := tgbotapi.NewMessage(msg.Chat.ID, text)
	resp.ReplyToMessageID = msg.MessageID

	_, err = api.Send(resp)

	return
}
