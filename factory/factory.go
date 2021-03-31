package factory

import (
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/email"
	"github.com/dealmaker/procedure/item_get"
	"github.com/dealmaker/procedure/item_upload"
	"github.com/dealmaker/shared/auth"
	"github.com/dealmaker/shared/base"
	"github.com/itzmeerkat/streamline"
	"time"
)

var Factory *streamline.Factory

func init() {
	Factory = streamline.New()
}

func BuildStreamlines() {
	itemGet := Factory.NewStreamline("/item/get", "get", "item")
	itemGet.Add("query items", item_get.QueryItem)

	itemUpload := Factory.NewStreamline("/item/upload", "upload", "item")
	itemUpload.Add("val", auth.Validate)
	itemUpload.Add("rua", item_upload.InsertItem)

	signup := Factory.NewStreamline("/auth/user/signup", "signup", "user")
	signup.Add("insert to db", auth_db.InsertUser)
	signup.Add("sign_token", auth.SignTokenWithTokenExpireTime(time.Hour * 24))
	signup.Add("send activation email", email.SendActivationEmail)


	login := Factory.NewStreamline("/auth/user/login", "login", "user")
	login.Add("get user form db", auth_db.GetUser)
	login.Add("sign_token", auth.SignTokenWithTokenExpireTime(60 * time.Minute))

	recoverPw := Factory.NewStreamline("/auth/user/recover", "recover", "user")
	recoverPw.Add("get user form db", auth_db.GetUser)
	recoverPw.Add("set recover", auth.SetRecover)
	recoverPw.Add("sign_token", auth.SignTokenWithTokenExpireTime(time.Hour * 24))
	recoverPw.Add("send email", email.SendRecoveryEmail)

	activeUser := Factory.NewStreamline("/auth/user/activate", "activate", "user")
	activeUser.Add("val", auth.Validate)
	activeUser.Add("get user form db", auth_db.ActiveUser)
	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", base.BaseRequestFiller)
		//v.InsertFront("Authenticator", slice.Authenticator)
	}
}