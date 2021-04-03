package factory

import (
	"github.com/dealmaker/dal"
	"github.com/dealmaker/procedure/email"
	"github.com/dealmaker/procedure/item"
	"github.com/dealmaker/shared/auth"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/itzmeerkat/streamline"
	"time"
)

var Factory *streamline.Factory
var authInstance *auth.WorkerInstance
var itemInstance *item.WorkerInstance
var emailInstance *email.WorkerInstance
func init() {
	Factory = streamline.New()
	authInstance = auth.WorkerInstance{
		FuncGetCredUser:        dal.GetUser,
		FuncInsertCredUser:     dal.InsertUser,
		FuncUpdateCredUser:     dal.UpdateUser,
		InvalidTokenForgetTime: time.Minute * 65,
		TokenExpireTimes:        make(map[string]time.Duration),
	}.Init()

	itemInstance = item.WorkerInstance{
		FuncGetItem:    dal.GetItem,
		FuncUpdateItem: nil,
		FuncInsertItem: dal.InsertItem,
	}.Init()
	emailInstance = email.WorkerInstance{}.Init()
}

func BuildStreamlines() {
	itemGet := Factory.NewStreamline("/item/get", "get", "item")
	itemGet.Add("query items", itemInstance.GetItem)

	itemUpload := Factory.NewStreamline("/item/upload", "upload", "item")
	itemUpload.Add("val", authInstance.ValidateJwt)
	itemUpload.Add("rua", itemInstance.InsertItem)

	signup := Factory.NewStreamline("/auth/user/signup", "signup", "user")
	signup.Add("insert to db", authInstance.NewUser)
	signup.Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeActivate))
	signup.Add("send email", emailInstance.BuildActivationEmail)
	signup.Add("send email", emailInstance.SendEmail)


	login := Factory.NewStreamline("/auth/user/login", "login", "user")
	login.Add("get user form db", authInstance.ValidatePassword)
	login.Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeNormal))

	recoverPw := Factory.NewStreamline("/auth/user/recover", "recover", "user")
	recoverPw.Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeRecover))
	recoverPw.Add("send email", emailInstance.BuildRecoverEmail)
	recoverPw.Add("send email", emailInstance.SendEmail)

	activeUser := Factory.NewStreamline("/auth/user/activate", "activate", "user")
	activeUser.Add("val", authInstance.ValidateJwt)
	activeUser.Add("get user form db", authInstance.ActivateUser)
	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", base.BaseRequestFiller)
	}
}