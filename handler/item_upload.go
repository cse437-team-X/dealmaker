package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/item_upload"
	"github.com/dealmaker/resp_def"
	"github.com/gin-gonic/gin"
)

func ItemUpload(c *gin.Context) {
	s := factory.Factory.Get("/item/upload")
	domain := item_upload.ItemUploadDomain{}
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
	c.JSON(code, resp_def.UserLoginResponse{
		Token: "Success",
	})
}