package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/codegen/idl"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/base_model"
	"github.com/gin-gonic/gin"
)

func UserLogin (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/login")
	domain := base_model.UserLoginDomain{}
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