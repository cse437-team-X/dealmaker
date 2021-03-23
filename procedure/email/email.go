package email

import (
	"fmt"
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/shared/auth"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
)

var client *sendgrid.Client

func InitEmailClient() {
	//key := os.Getenv("SENDGRID_API_KEY")
	//fmt.Println(key)
	client = sendgrid.NewSendClient("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")
}

func SendEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(auth.CredUserInterface).GetCredUser()
	token := c.DataDomain.(auth.JwtInterface).GetJwtAuth()
	from := mail.NewEmail("Dealmaker Admin", "jiayi.zhang@wustl.edu")
	subject := "YOUR PASSWORD RECOVERY LINK"
	to := mail.NewEmail(data.LoginName, data.LoginName)
	plainTextContent := "Your token here:" + token.Token
	htmlContent := "Your token here:" + token.Token
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	response, err := client.Send(message)
	if err != nil {
		c.Errorw("Send email error", err)
		return http.StatusInternalServerError
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return http.StatusOK
}
