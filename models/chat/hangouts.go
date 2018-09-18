package chat

import (
	"context"
	"net/http"
)

type Hangouts struct{}

func (h Hangouts) SendMessage(ctx context.Context, client *http.Client, message string) error {
	return nil
}
