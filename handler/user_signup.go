package handler

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserSignupDomain struct {
	base.Base
	auth_db.UserCredModel
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