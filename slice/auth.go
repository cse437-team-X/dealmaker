package slice

import (
	"gitee.com/fat_marmota/streamline"
)

type AuthInterface interface {
	GetHashedPassword() string
	GetUsername() string
	GetBaseTime() int64
	GetBaseLogId() string
}

func Authenticator(ctx *streamline.Context, di interface{}) error {
	data := di.(AuthInterface)
	d1 := data.GetHashedPassword()
	d2 := data.GetUsername()
	t := data.GetBaseTime()
	id := data.GetBaseLogId()
	ctx.Logger.Debugf("%v %v %v %v", d1,d2,t,id)
	return nil
}