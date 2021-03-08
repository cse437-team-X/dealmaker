package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/codegen/idl"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/item_upload"
	"github.com/gin-gonic/gin"
)

func ItemUpload(c *gin.Context) {
	s := factory.Factory.Get("/item/upload")
	domain := item_upload.ItemUploadDomain{}
	err := c.Bind(&domain)
	if err != nil {
		return
	}
	conv := streamline.NewConveyorBelt(s, c, &domain)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}
	c.JSON(code, idl.UserLoginResponse{
		Message: "Success",
	})
}