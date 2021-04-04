package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
)

type UserActivateDomain struct {
	base.Base
	model.JwtAuth
	model2.EmailContent
}


func ActivateUser(c *gin.Context) {
	s := factory.Factory.Get("/auth/user/activate")
	token := c.Query("token")

	domain := UserActivateDomain{}
	domain.Token = token
	//err := c.Bind(&domain)
	//domain.LoginName = loginName

	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)

	conv.Debugw("domain",domain)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}
	c.JSON(code, nil)
}
