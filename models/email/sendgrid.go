package email

import (
	"context"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	APIKey string
}

func (p SendGrid) SendEmail(ctx context.Context, email EmailItem) error {
	from := mail.NewEmail("", email.From)
	subject := email.Subject
	to := mail.NewEmail("", email.To)
	plainTextContent := email.Message
	htmlContent := email.HTMLMessage
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(p.APIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}
