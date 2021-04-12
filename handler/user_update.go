package handler

import (
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
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
	code := ExecuteStreamline(c, "/auth/user/update", domain)
	c.JSON(code, nil)
}
