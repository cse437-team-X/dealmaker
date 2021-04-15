package email

import (
	"fmt"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type BuildContactEmailInterface interface {
	//GetCredUser() *model.CredUser
	GetJwtAuth() *model.JwtAuth
	GetContactEmailInfo() *model2.ContactEmailInfo
	GetEmailContent() *model2.EmailContent
}

const contentFmt = `Hi %v,
%v is interested in your item %v, please send an email to the address and get contact.
Disclaimer: We do not hold responsibility any further from this point.
`

func (w *WorkerInstance) BuildContactEmail(c *streamline.ConveyorBelt) int {
	jwtData := c.DataDomain.(BuildContactEmailInterface).GetJwtAuth()
	email := c.DataDomain.(BuildContactEmailInterface).GetEmailContent()

	contact := c.DataDomain.(BuildContactEmailInterface).GetContactEmailInfo()

	toUser := w.FuncGetCredUser(&contact.To)
	c.Debugw("touser", toUser)
	toEmail := toUser.LoginName+"@wustl.edu"

	email.Title = "Contact request for " + contact.Item.Title
	email.Body = fmt.Sprintf(contentFmt, toUser.LoginName, jwtData.TokenClaim.LoginName, contact.Item.Title)
	email.To = toEmail
	email.Recipient = jwtData.TokenClaim.LoginName
	return http.StatusOK
}

