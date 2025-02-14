package main

import (
	"context"
	"flea-market/cmd/releaser/application/telegram"
	"flea-market/cmd/releaser/settings"
	"flea-market/internal/graceful"
	"flea-market/internal/logger"
	"fmt"
	"os"
	"time"
)

func main() {
	conf := settings.NewConfig()
	err := conf.WithEnv()
	if err != nil {
		panic("telegram data is empty")
	}

	l := logger.NewLogger(logger.Debug)
	ctx := context.Background()
	ctx = l.WithContextFields(ctx,
		logger.Field{Key: "pid", Value: os.Getpid()},
	)
	defer l.Sync()

	if conf.Telegram.Token == "" || conf.Telegram.Admin == "" {
		panic("telegram data is empty")
	}

	lastUpdate := 0
	r := make(chan telegram.ButtonQuery)
	p := telegram.NewPool(conf.Telegram, conf.PoolTimeout, r, lastUpdate, l)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				l.InfoCtx(ctx, "reader context", logger.Field{Key: "message", Value: ctx.Err()})
				return
			case msg, ok := <-r:
				if !ok {
					return
				}
				go NewMessage(ctx, msg, l)
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}(ctx)

	gs := graceful.NewGracefulShutdown(ctx, 1*time.Second)
	gs.AddTask(p)
	err = gs.Wait()
	l.ErrorCtx(ctx, err.Error())
}

func NewMessage(ctx context.Context, msg telegram.ButtonQuery, l *logger.Logger) {
	// if msg.Data != "" {}
	// if msg.From.ID == admin
	// if msg.Message.Chat.ID == admin && msg.Message.Chat.Type == "private"
	l.InfoCtx(ctx, fmt.Sprintf("new message ID %d release tag %s", msg.Message.ID, msg.Data))
}
