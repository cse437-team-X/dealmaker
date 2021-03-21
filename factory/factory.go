package factory

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/item_get"
	"github.com/dealmaker/procedure/item_upload"
	"github.com/dealmaker/shared/auth"
	"github.com/dealmaker/shared/base"
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
	//signup.Add("validate signup req", auth.ValidateSignUp)
	signup.Add("insert to db", auth_db.InsertUser)


	login := Factory.NewStreamline("/auth/user/login", "login", "user")
	//login.Add("check_username_pw", auth.ValidateCredUser)
	login.Add("get user form db", auth_db.GetUser)
	login.Add("sign_token", auth.SignTokenWithTokenExpireTime(10 * time.Minute))

	recoverPw := Factory.NewStreamline("/auth/user/recover", "recover", "user")
	//recoverPw.Add("check_username_pw", auth.ValidateCredUser)
	recoverPw.Add("get user form db", auth_db.GetUser)
	recoverPw.Add("sign_token", auth.SignTokenWithTokenExpireTime(time.Hour * 24))


	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", base.BaseRequestFiller)
		//v.InsertFront("Authenticator", slice.Authenticator)
	}
}