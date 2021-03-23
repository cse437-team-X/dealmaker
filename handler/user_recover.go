package handler

import (
	"github.com/dealmaker/factory"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
)

func UserRecover (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/recover")
	domain := UserLoginDomain{}
	loginName := c.Query("username")
	//err := c.Bind(&domain)
	domain.LoginName = loginName
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