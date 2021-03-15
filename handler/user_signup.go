package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSignup(c *gin.Context) {
	s := factory.Factory.Get("/auth/user/signup")
	domain := struct {
		base.Base
		auth_db.UserCredModel
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
	c.JSON(http.StatusOK, nil)
}