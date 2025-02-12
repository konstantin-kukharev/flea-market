package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"flea-market/cmd/releaser/settings"
	"flea-market/internal/logger"
	"flea-market/internal/roundtripper"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Pool struct {
	cli          *http.Client
	log          *logger.Logger
	telegram     *settings.Telegram
	poolTimeout  time.Duration
	poolAddr     string
	lastUpdateID int
	reportChan   chan<- UpdatesMessage
}

// NewPool создает новый пул.
//
// ctx - контекст.
// token - токен.
// pd - интервал опроса telegram.
// r - таймаут запроса с сообщением.
//
// Возвращает новый пул.
func NewPool(t *settings.Telegram, pd time.Duration, rc chan<- UpdatesMessage, lu int, l *logger.Logger) *Pool {
	var rt http.RoundTripper
	rt = http.DefaultTransport
	rt = roundtripper.NewCompress(rt)

	return &Pool{
		cli: &http.Client{
			Transport: rt,
			Timeout:   pd,
		},
		log:          l,
		telegram:     t,
		poolAddr:     fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", t.Token),
		poolTimeout:  pd,
		reportChan:   rc,
		lastUpdateID: lu,
	}
}

type RequestGetUpdates struct {
	Offset  int      `json:"offset,omitempty"`
	Limit   int      `json:"limit,omitempty"`
	Timeout int      `json:"timeout,omitempty"`
	Updates []string `json:"allowed_updates,omitempty"`
}

type UpdatesMessage struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Language  string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date     int    `json:"date"`
	Text     string `json:"text"`
	Entities []struct {
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		Type   string `json:"type"`
	} `json:"entities,omitempty"`
}

type ResponseGetUpdates struct {
	OK     bool `json:"ok"`
	Result []struct {
		UpdateID int            `json:"update_id"`
		Message  UpdatesMessage `json:"message"`
	} `json:"result"`
}

func (p *Pool) Run(ctx context.Context) error {
	defer close(p.reportChan)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			u, err := p.request(ctx, p.lastUpdateID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if len(u.Result) == 0 {
				continue
			}
			for _, res := range u.Result {
				p.lastUpdateID = res.UpdateID + 1
				p.reportChan <- res.Message
			}
		}
	}
}

func (p *Pool) request(ctx context.Context, offset int) (updates *ResponseGetUpdates, err error) {
	updates = new(ResponseGetUpdates)
	pl := &RequestGetUpdates{Offset: offset}
	b, err := json.Marshal(pl)
	if err != nil {
		p.log.ErrorCtx(ctx, err.Error())
		return updates, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, p.poolAddr, bytes.NewBuffer(b))
	if err != nil {
		p.log.ErrorCtx(ctx, err.Error())
		return updates, err
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := p.cli.Do(request)
	if err != nil {
		p.log.ErrorCtx(ctx, err.Error())
		return updates, err
	}

	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		p.log.ErrorCtx(ctx, "error", logger.Field{Key: "error", Value: err.Error()})
		return updates, err
	}
	updates = new(ResponseGetUpdates)
	err = json.Unmarshal(b, updates)

	return updates, err
}
