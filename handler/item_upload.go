package handler

import (
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
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
	s := factory.Factory.Get("/item/upload")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}
	resp := UploadItemResponse{
		Message: domain.BaseMessage,
		ItemID:  domain.ObjId,
	}

	c.JSON(code, resp)
}