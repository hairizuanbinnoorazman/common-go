package email

import "context"

type Service interface {
	SendEmail(ctx context.Context, email EmailItem) error
}
