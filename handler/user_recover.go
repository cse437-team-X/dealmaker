package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
)

type UserRecoverDomain struct {
	base.Base
	model2.EmailContent
	model.CredUser
	model.JwtAuth
}

func UserRecover (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/recover")
	domain := UserLoginDomain{}
	loginName := c.Query("username")
	//err := c.Bind(&domain)
	domain.LoginName = loginName
	domain.Scope = "recover"
	//if err != nil {
	//	return
	//}
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)

	conv.Debugw("domain",domain)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}
}