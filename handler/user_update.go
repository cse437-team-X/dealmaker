package handler

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserUpdateDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
}

func UserUpdate(c *gin.Context) {
	s := factory.Factory.Get("/auth/user/update")
	domain := UserUpdateDomain{}

	err := c.Bind(&domain)

	domain.Token = c.Query("token")

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
