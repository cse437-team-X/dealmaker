package handler

import (
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
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

	code := ExecuteStreamline(c, "/auth/user/recover", domain)
	c.JSON(code ,nil)
}