package email

type EmailItem struct {
	To          string
	From        string
	Subject     string
	CC          string
	BCC         string
	Message     string
	HTMLMessage string
}
