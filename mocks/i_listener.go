// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	mock "github.com/stretchr/testify/mock"

	tgchain "github.com/egnd/go-tgchain"
)

// IListener is an autogenerated mock type for the IListener type
type IListener struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *IListener) Add(_a0 tgchain.EventType, _a1 ...tgchain.IEventHandler) tgchain.IListener {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 tgchain.IListener
	if rf, ok := ret.Get(0).(func(tgchain.EventType, ...tgchain.IEventHandler) tgchain.IListener); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tgchain.IListener)
		}
	}

	return r0
}

// Listen provides a mock function with given fields: _a0, _a1
func (_m *IListener) Listen(_a0 context.Context, _a1 tgbotapi.UpdatesChannel) {
	_m.Called(_a0, _a1)
}
