package factory

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/item_upload"
	"github.com/dealmaker/shared/auth"
	"github.com/dealmaker/shared/base"
)

var Factory *streamline.Factory

func init() {
	Factory = streamline.New()
}

func BuildStreamlines() {
	itemUpload := Factory.NewStreamline("/item/upload", "upload", "item")
	itemUpload.Add("rua", item_upload.InsertItem)

	signup := Factory.NewStreamline("/auth/user/signup", "signup", "user")
	signup.Add("validate signup req", auth.ValidateSignUp)
	signup.Add("insert to db", auth_db.InsertUser)


	login := Factory.NewStreamline("/auth/user/login", "login", "user")
	login.Add("check_username_pw", auth.ValidateCredUser)
	login.Add("get user form db", auth_db.GetUser)
	login.Add("sign_token", auth.SignToken)



	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", base.BaseRequestFiller)
		//v.InsertFront("Authenticator", slice.Authenticator)
	}
}