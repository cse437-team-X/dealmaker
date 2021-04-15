package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserActivateDomain struct {
	base.Base
	model.JwtAuth
	model2.EmailContent
}

func ActivateUser(c *gin.Context) {
	token := c.Query("token")

	domain := UserActivateDomain{}
	domain.Token = token


	s := factory.Factory.Get("/auth/user/activate")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}
	c.JSON(code, nil)
}
