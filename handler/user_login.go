package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	"github.com/dealmaker/resp_def"
	"github.com/dealmaker/shared/auth"
	"github.com/gin-gonic/gin"
)

func UserLogin (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/login")
	domain := struct {
		model.UserInfoDomain
		auth.JwtAuth
	}{}
	err := c.Bind(&domain)
	if err != nil {
		return
	}
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}
	c.JSON(code, resp_def.UserLoginResponse{
		Token: domain.GetToken(),
	})
}