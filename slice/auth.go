package slice

import (
	"gitee.com/fat_marmota/streamline"
)

type AuthenticatorInterface interface {
	BaseInterface
	GetHashedPassword() string
	GetUsername() string
}

func Authenticator(c *streamline.ConveyorBelt) error {
	data := c.DataPanel.(AuthenticatorInterface)
	d1 := data.GetHashedPassword()
	d2 := data.GetUsername()

	t := data.GetBaseTime()

	if queryUsernamePassword(d2, d1) == true {
		data.SetBaseCode(200)
	} else {
		data.SetBaseCode(502)
	}

	c.Logger.Debugf("%v %v %v", d1,d2,t)
	return nil
}

func queryUsernamePassword(username, hashpw string) bool {
	if username == "admin" && hashpw == "admin" {
		return true
	}
	return false
}