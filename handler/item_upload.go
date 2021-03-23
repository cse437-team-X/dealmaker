package handler

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type UploadItemDomain struct {
	base.Base
	model2.JwtAuth
	model.Item
}

func ItemUpload(c *gin.Context) {
	s := factory.Factory.Get("/item/upload")

	domain := UploadItemDomain{}
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
	c.JSON(code, nil)
}