package email

import "context"

type MailGun struct{}

func (m MailGun) SendEmail(ctx context.Context, email EmailItem) error {
	return nil
}
