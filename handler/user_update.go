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

type UserUpdateInput struct {
	Token string
	LoginName string
	HashedPassword string
	Status int
}

func UserUpdate(c *gin.Context) {
	input := UserUpdateInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := UserUpdateDomain{
		JwtAuth: model.JwtAuth{
			Token: input.Token,
		},
		CredUser: model.CredUser{
			HashedPassword: input.HashedPassword,
			LoginName:      input.LoginName,
			Status:         input.Status,
		},
	}

	s := factory.Factory.Get("/auth/user/update")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	c.JSON(code, nil)
}
