package chat

import (
	"context"
	"net/http"

	hangouts "google.golang.org/api/chat/v1"
)

type Hangouts struct {
	Parent string
}

func (h Hangouts) SendMessage(ctx context.Context, client *http.Client, message string) error {
	serviceChat, err := hangouts.New(client)
	if err != nil {
		return err
	}
	spaceMessageService := hangouts.NewSpacesMessagesService(serviceChat)
	msg := hangouts.Message{
		Text: message,
	}
	createCall := spaceMessageService.Create(h.Parent, &msg)
	createCall.Do()
	return nil
}
