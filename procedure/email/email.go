package email

import (
	"fmt"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
)

var client *sendgrid.Client

func InitEmailClient() {
	key := os.Getenv("SENDGRID_API_KEY")
	fmt.Println(key)
	client = sendgrid.NewSendClient(key)
}

func SendRecoveryEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(auth_db.AuthDBInterface).GetUserCredModel()
	token := c.DataDomain.(model.JwtInterface).GetJwtAuth()
	from := mail.NewEmail("Dealmaker Admin", "jiayi.zhang@wustl.edu")
	subject := "YOUR PASSWORD RECOVERY LINK"

	toEmail := data.LoginName+"@wustl.edu"
	c.Infow("sending email to", toEmail)
	to := mail.NewEmail(data.LoginName, toEmail)
	plainTextContent := "Hi, "+data.LoginName+"You will be able to use this TOKEN to reset your password:" + token.Token
	htmlContent := "Hi, "+data.LoginName+"You will be able to use this TOKEN to reset your password:" + token.Token

	c.Infow("email", plainTextContent)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	response, err := client.Send(message)
	if err != nil {
		c.Errorw("Send email error", err)
		return http.StatusInternalServerError
	} else {
		c.Infow("status",response.StatusCode,
			"body",response.Body,"header",response.Headers)
	}
	return http.StatusOK
}

func SendActivationEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(auth_db.AuthDBInterface).GetUserCredModel()
	token := c.DataDomain.(model.JwtInterface).GetJwtAuth()
	from := mail.NewEmail("Dealmaker Admin", "jiayi.zhang@wustl.edu")
	subject := "Activate your account!"

	toEmail := data.LoginName+"@wustl.edu"
	to := mail.NewEmail(data.LoginName, toEmail)
	activationLink := "http://437.cpp.moe/auth/user/activate?token="+token.Token

	plainTextContent := "Hi, "+data.LoginName+"\nclick link to activate your account:" + activationLink
	htmlContent := "Hi, "+data.LoginName+"\nclick link to activate your account:" + activationLink
	c.Infow("email", plainTextContent)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	response, err := client.Send(message)
	if err != nil {
		c.Errorw("Send email error", err)
		return http.StatusInternalServerError
	} else {
		c.Infow("status",response.StatusCode,
			"body",response.Body,"header",response.Headers)
	}
	return http.StatusOK
}