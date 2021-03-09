package factory

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/slice"
)

var Factory *streamline.Factory

func init() {
	Factory = streamline.New()
}

func BuildStreamlines() {
	//userLoginSl := Factory.NewStreamline("/auth/user/login", "login", "user")
	//userLoginSl.Add("Login", slice.Login)
	//itemUpload := Factory.NewStreamline("/item/upload", "upload", "item")
	//itemUpload.Add("rua", slice.Item)

	signup := Factory.NewStreamline("/auth/user/signup", "signup", "user")
	signup.Add("add_user", slice.SignUp)

	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", slice.BaseRequestFiller)
		//v.InsertFront("Authenticator", slice.Authenticator)
	}
}