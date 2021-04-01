package handler

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserSignupDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
}

func UserSignup(c *gin.Context) {
	s := factory.Factory.Get("/auth/user/signup")
	domain := UserSignupDomain{}

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
