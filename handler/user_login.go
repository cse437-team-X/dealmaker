package handler

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/resp_def"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
)

type UserLoginDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
}

func UserLogin (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/login")
	//loginName := c.Query("username")

	domain := UserLoginDomain{}
	err := c.Bind(&domain)
	//domain.LoginName = loginName
	if err != nil {
		return
	}
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)

	conv.Debugw("domain",domain)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}
	c.JSON(code, resp_def.UserLoginResponse{
		Token: domain.Token,
	})
}
