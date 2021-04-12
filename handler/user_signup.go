package handler

import (
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type UserSignupDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
	model2.EmailContent
}

type UserSignupInput struct {
	LoginName string
	HashedPassword string
}

func UserSignup(c *gin.Context) {
	input := UserSignupInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := UserSignupDomain{
		CredUser: model.CredUser{
			HashedPassword: input.HashedPassword,
			LoginName:      input.LoginName,
		},
	}
	code := ExecuteStreamline(c, "/auth/user/signup", domain)

	c.JSON(code, nil)
}
