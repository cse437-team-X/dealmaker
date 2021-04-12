package handler

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
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
	code := ExecuteStreamline(c, "/auth/user/login", domain)
	c.JSON(code, UserLoginResponse{
		Token: domain.Token,
		Message: domain.BaseMessage,
	})
}
