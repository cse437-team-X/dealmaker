package slice

import (
	"gitee.com/fat_marmota/streamline"
	"net/http"
)

type AuthenticatorInterface interface {
	BaseInterface
	GetHashedPassword() string
	GetUsername() string
}

func Login(c *streamline.ConveyorBelt) int {
	data := c.DataPanel.(AuthenticatorInterface)
	d1 := data.GetHashedPassword()
	d2 := data.GetUsername()

	if queryUsernamePassword(d2, d1) != true {
		return http.StatusForbidden
	}

	c.Logger.Debugw("Login",
		"username",d1,
			"hashed_pw",d2)
	return http.StatusOK
}

func queryUsernamePassword(username, hashpw string) bool {
	if username == "admin" && hashpw == "admin" {
		return true
	}
	return false
}

func Authenticator(c *streamline.ConveyorBelt) int {
	data := c.DataPanel.(BaseInterface)
	session := data.GetSessionId()
	c.Logger.Debugw("Authenticator", "sessionId", session)
	return http.StatusOK
}
