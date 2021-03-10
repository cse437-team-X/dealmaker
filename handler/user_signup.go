package handler

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/model"
	"github.com/dealmaker/codegen/idl"
	"github.com/dealmaker/factory"
	"github.com/gin-gonic/gin"
)

func UserSignup (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/signup")
	domain := model.UserInfoDomain{}
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