package factory

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/slice"
)

var Factory *streamline.Factory

func init() {
	Factory = streamline.New()
}

func BuildStreamlines() {
	userLoginSl := Factory.NewStreamline("/auth/user/login", streamline.Context{
		Logger:   log.GlobalLogger,
		Action:   "login",
		Resource: "user",
	})
	userLoginSl.Add("Authenticator", slice.Authenticator)
}
