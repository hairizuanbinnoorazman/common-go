package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Slack struct {
	SlackToken string
	ChannelID  string
}

func (s Slack) SendMessage(ctx context.Context, client *http.Client, message string) error {
	url := "https://slack.com/api/chat.postMessage"
	type postMessageBody struct {
		Token   string `json:"token"`
		Channel string `json:"channel"`
		Text    string `json:"Text"`
	}
	body := postMessageBody{Token: s.SlackToken, Channel: s.ChannelID, Text: message}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bodyJSON)
	_, err = client.Post(url, "application/json", reader)
	if err != nil {
		return err
	}
	return nil
}
