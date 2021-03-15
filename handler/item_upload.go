package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

func ItemUpload(c *gin.Context) {
	s := factory.Factory.Get("/item/upload")
	domain := struct {
		base.Base
		model.ItemModel
	}{}
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