package email

import (
	"github.com/dealmaker/procedure/email/model"
	"github.com/itzmeerkat/streamline"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
)

type SendEmailInterface interface {
	GetEmailContent() *model.EmailContent
}

func (w *WorkerInstance) SendEmail(c *streamline.ConveyorBelt) int {
	email := c.DataDomain.(SendEmailInterface).GetEmailContent()

	to := mail.NewEmail(email.Recipient, email.To)
	plainTextContent := email.Body
	htmlContent := email.Body
	c.Infow("email", plainTextContent)
	from := mail.NewEmail("Dealmaker Admin", "jiayi.zhang@wustl.edu")

	message := mail.NewSingleEmail(from, email.Title, to, plainTextContent, htmlContent)

	response, err := w.client.Send(message)
	if err != nil {
		c.Errorw("Send email error", err)
		return http.StatusInternalServerError
	} else {
		c.Infow("status",response.StatusCode,
			"body",response.Body,"header",response.Headers)
	}
	return http.StatusOK
}