package handler

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/resp_def"
	"github.com/gin-gonic/gin"
)

func UserRecover (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/recover")
	domain := UserLoginDomain{}
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
		Token: domain.Token,
	})
}