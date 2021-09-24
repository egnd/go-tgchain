package tgchain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type IListener interface {
	Add(EventType, ...IEventHandler) IListener
	Listen(context.Context, tgbotapi.UpdatesChannel)
}

// Listener is struct which is listening and handling events.
type Listener struct {
	handlers map[EventType][]IEventHandler
	warnUpd  WarnUpd
}

// NewListener constructor for Listener struct.
func NewListener(warnUpd WarnUpd) *Listener {
	return &Listener{
		warnUpd:  warnUpd,
		handlers: make(map[EventType][]IEventHandler),
	}
}

// Add adds decorators for handling specific Telegram event.
func (b *Listener) Add(event EventType, handlers ...IEventHandler) IListener {
	b.handlers[event] = append(b.handlers[event], handlers...)

	return b
}

func (b *Listener) buildChains() map[EventType]IEventHandler {
	chains := map[EventType]IEventHandler{}

	for event, handlers := range b.handlers {
		chains[event] = nil

		for i := len(handlers) - 1; i >= 0; i-- {
			if chains[event] == nil {
				chains[event] = handlers[i]
			} else {
				handlers[i].Decorate(chains[event])
				chains[event] = handlers[i]
			}
		}
	}

	return chains
}

// Listen starts listening incoming messages in channel.
func (b *Listener) Listen(ctx context.Context, updChan tgbotapi.UpdatesChannel) {
	handlers := b.buildChains()

	for update := range updChan { // @TODO: add workers pool
		event := GetEventFrom(update)

		if item, ok := handlers[event]; ok {
			item.Handle(context.WithValue(ctx, CtxEventKey, event), update)
			continue
		}

		if b.warnUpd != nil {
			b.warnUpd("unexpected event", &update)
		}
	}
}