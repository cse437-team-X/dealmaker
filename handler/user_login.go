package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/resp_def"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

func UserLogin (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/login")
	domain := struct {
		base.Base
		auth_db.UserCredModel
		model.JwtAuth
	}{}
	err := c.Bind(&domain)

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
		Token: domain.GetToken(),
	})
}
