package factory

import (
	"github.com/dealmaker/dal"
	"github.com/dealmaker/procedure/email"
	"github.com/dealmaker/procedure/item"
	"github.com/dealmaker/shared/access_control"
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
var acInstance *access_control.WorkerInstance

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

	acInstance = access_control.WorkerInstance{
		ConfPath:   "./conf/rbac/model.conf",
		PolicyPath: "./conf/rbac/policy.csv",
	}.Init()
}

func BuildStreamlines() {
	Factory.NewStreamline("/item/get", "get", "item").
		Add("query items", itemInstance.GetItem)

	Factory.NewStreamline("/item/upload", "upload", "item").
		Add("val", authInstance.ValidateJwt).
		Add("rbac", acInstance.CheckAccess).
		Add("rua", itemInstance.InsertItem)

	Factory.NewStreamline("/auth/user/signup", "signup", "user").
		Add("insert to db", authInstance.NewUser).
		Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeActivate)).
		Add("send email", emailInstance.BuildActivationEmail).
		Add("send email", emailInstance.SendEmail)


	Factory.NewStreamline("/auth/user/login", "login", "user").
		Add("get user form db", authInstance.ValidatePassword).
		Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeNormal))

	Factory.NewStreamline("/auth/user/recover", "recover", "user").
		Add("sign_token", authInstance.SignTokenToScope(model.JwtScopeRecover)).
		Add("send email", emailInstance.BuildRecoverEmail).
		Add("send email", emailInstance.SendEmail)

	Factory.NewStreamline("/auth/user/activate", "activate", "user").
		Add("val", authInstance.ValidateJwt).
		Add("get user form db", authInstance.ActivateUser)

	Factory.NewStreamline("/auth/user/update", "update", "user").
		Add("validate jwt", authInstance.ValidateJwt).
		Add("update user", authInstance.UpdateUser)

	AddBaseRequestFillerToAll()
}

func AddBaseRequestFillerToAll() {
	for _,v := range Factory.GetAllStreamlines() {
		v.InsertFront("BaseRequestFiller", base.BaseRequestFiller)
	}
}