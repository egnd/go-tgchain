package tgchain_test

import (
	"context"
	"testing"

	"github.com/egnd/go-tgchain"
	"github.com/egnd/go-tgchain/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func Test_AbstractHandler(t *testing.T) {
	ctx := context.Background()
	upd := tgbotapi.Update{}

	mHandler := &mocks.IEventHandler{}
	mHandler.On("Handle", ctx, upd).Return(nil)

	h := tgchain.AbstractHandler{}
	assert.EqualValues(t, "", h.GetName())
	assert.EqualValues(t, "", h.GetDescription())
	h.Next(ctx, upd)

	h.Decorate(mHandler)
	h.Next(ctx, upd)

	mHandler.AssertExpectations(t)
}
