package email

import (
	"fmt"
	"github.com/dealmaker/shared/auth"
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

func SendEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(auth.CredUserInterface).GetCredUser()
	token := c.DataDomain.(auth.JwtInterface).GetJwtAuth()
	from := mail.NewEmail("Dealmaker Admin", "jiayi.zhang@wustl.edu")
	subject := "YOUR PASSWORD RECOVERY LINK"

	toEmail := data.LoginName+"@wustl.edu"
	c.Infow("sending email to", toEmail)
	to := mail.NewEmail(data.LoginName, toEmail)
	plainTextContent := "Hi, "+data.LoginName+"You will be able to use this TOKEN to reset your password:" + token.Token
	htmlContent := "Hi, "+data.LoginName+"You will be able to use this TOKEN to reset your password:" + token.Token
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
