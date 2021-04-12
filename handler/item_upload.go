package handler

import (
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type UploadItemDomain struct {
	base.Base
	model2.JwtAuth
	model.Item
}

type UploadItemInput struct {
	Token string
	model.Item
}

type UploadItemResponse struct {
	Message string
	ItemID string
}

func ItemUpload(c *gin.Context) {
	input := UploadItemInput{}
	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := UploadItemDomain{
		JwtAuth: model2.JwtAuth{Token: input.Token},
		Item:    input.Item,
	}
	code := ExecuteStreamline(c, "/item/upload", domain)

	resp := UploadItemResponse{
		Message: domain.BaseMessage,
		ItemID:  domain.ObjId,
	}

	c.JSON(code, resp)
}