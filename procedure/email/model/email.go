package model

type EmailContent struct {
	To        string
	Recipient string
	Title     string
	Body      string
}
func (e *EmailContent) GetEmailBody() *EmailContent {
	return e
}
