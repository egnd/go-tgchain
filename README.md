# go-tgchain

[![Go Reference](https://pkg.go.dev/badge/github.com/egnd/go-tgchain.svg)](https://pkg.go.dev/github.com/egnd/go-tgchain)
[![Go Report Card](https://goreportcard.com/badge/github.com/egnd/go-tgchain)](https://goreportcard.com/report/github.com/egnd/go-tgchain)
[![Coverage](https://gocover.io/_badge/github.com/egnd/go-tgchain)](https://gocover.io/github.com/egnd/go-tgchain)
[![Pipeline](https://github.com/egnd/go-tgchain/actions/workflows/pipeline.yml/badge.svg)](https://github.com/egnd/go-tgchain/actions?query=workflow%3APipeline)

DEPRECATED use https://github.com/egnd/go-toolbox/tree/master/tgchain instead

Chained wrapper for handling tg updates

### Example
```go
package main

import (
	"log"
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/egnd/go-tgchain"
)

func main() {
	api, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := api.GetUpdatesChan(u)

	tgchain.NewListener(nil, nil)
		.Add(tgchain.EventMessage, 
			// @TODO: your handlers, children of IEventHandler
		)
		.Listen(context.Background(), updates)
}

```
