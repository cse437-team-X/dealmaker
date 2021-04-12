package handler

import (
	model2 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type ItemGetDomain struct {
	base.Base
	model.JwtAuth
	model2.GetItemDomain
}

type ItemGetInput struct {
	Token string
	model2.QueryFilter
}

type ItemGetResponse struct {
	Message string
	Items []model2.Item
}

func ItemGetHandler(c *gin.Context) {
	input := ItemGetInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := ItemGetDomain{
		JwtAuth:       model.JwtAuth{
			Token: input.Token,
		},
		GetItemDomain: model2.GetItemDomain{
			QueryFilter: input.QueryFilter,
		},
	}

	code := ExecuteStreamline(c, "/item/get", domain)

	resp := ItemGetResponse{
		Message: domain.BaseMessage,
		Items:  domain.Result,
	}
	c.JSON(code, resp)
}