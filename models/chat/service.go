package chat

import (
	"context"
	"net/http"
)

type Service interface {
	SendMessage(ctx context.Context, client *http.Client, message string) error
}
