package slice

import (
	"gitee.com/fat_marmota/streamline"
)

type AuthReqInterface interface {
	GetHashedPassword() string
	GetUsername() string
	BaseReqInterface
}

type AuthRespInterface interface {
	SetMessage(v string)
	BaseRespInterface
}

func Authenticator(self *streamline.Streamline, in interface{}, out interface{}) error {
	data := in.(AuthReqInterface)
	d1 := data.GetHashedPassword()
	d2 := data.GetUsername()

	t := data.GetBaseTime()

	resp := out.(AuthRespInterface)
	id := resp.GetBaseLogId()

	if queryUsernamePassword(d2, d1) == true {
		resp.SetBaseCode(200)
	} else {
		resp.SetBaseCode(503)
	}

	self.Logger.Debugf("%v %v %v %v", d1,d2,t,id)
	return nil
}

func queryUsernamePassword(username, hashpw string) bool {
	if username == "admin" && hashpw == "admin" {
		return true
	}
	return false
}