package handler

import (
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type UserActivateDomain struct {
	base.Base
	model.JwtAuth
	model2.EmailContent
}

func ActivateUser(c *gin.Context) {
	token := c.Query("token")

	domain := UserActivateDomain{}
	domain.Token = token

	code := ExecuteStreamline(c, "/auth/user/activate", domain)
	c.JSON(code, nil)
}
