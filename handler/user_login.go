package handler

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserLoginDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
}

type UserLoginInput struct {
	LoginName string
	HashedPassword string
}

type UserLoginResponse struct {
	Message string
	Token string
	Uid uint
}

func UserLogin (c *gin.Context) {
	input := UserLoginInput{}
	err := c.Bind(&input)
	if err != nil {
		return
	}
	domain := UserLoginDomain{
		CredUser: model.CredUser{
			LoginName: input.LoginName,
			HashedPassword: input.HashedPassword,
		},
	}

	s := factory.Factory.Get("/auth/user/login")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	c.JSON(code, UserLoginResponse{
		Token: domain.Token,
		Message: domain.BaseMessage,
		Uid: domain.ID,
	})
}
