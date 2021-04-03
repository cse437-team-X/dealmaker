package email

import (
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type BuildRecoverEmailInterface interface {
	GetJwtAuth() *model.JwtAuth
	GetCredUser() *model.CredUser
	GetEmailContent() *model2.EmailContent
}

func (w *WorkerInstance) BuildRecoverEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BuildRecoverEmailInterface).GetCredUser()
	token := c.DataDomain.(BuildRecoverEmailInterface).GetJwtAuth()

	email := c.DataDomain.(BuildRecoverEmailInterface).GetEmailContent()

	toEmail := data.LoginName+"@wustl.edu"
	c.Infow("sending email to", toEmail)
	link := "http://437.cpp.moe:8080/auth/user/recover?token="+token.Token

	content := "Hi, "+data.LoginName+"\nclick link to recover your account:" + link
	email.Title = "RECOVER PASSWORD"
	email.Body = content
	email.To = toEmail
	email.Recipient = data.LoginName
	return http.StatusOK
}

