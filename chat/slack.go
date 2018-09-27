package chat

import (
	"context"
	"net/http"
	"net/url"
)

type Slack struct {
	SlackToken string
	ChannelID  string
}

func (s Slack) SendMessage(ctx context.Context, client *http.Client, message string) error {
	slackUrl := "https://slack.com/api/chat.postMessage"
	resp, err := client.PostForm(slackUrl, url.Values{
		"token":   []string{s.SlackToken},
		"channel": []string{s.ChannelID},
		"text":    []string{message},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
