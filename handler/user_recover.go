package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserRecoverDomain struct {
	base.Base
	model2.EmailContent
	model.CredUser
	model.JwtAuth
}

func UserRecover (c *gin.Context) {
	domain := UserLoginDomain{}
	loginName := c.Query("username")
	domain.LoginName = loginName
	domain.Scope = "recover"

	s := factory.Factory.Get("/auth/user/recover")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	c.JSON(code ,nil)
}