package handler

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/codegen/idl"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin (c *gin.Context) {
	s := factory.Factory.Get("/auth/user/login")
	domain := model.UserLoginDomain{}
	c.Bind(&domain)
	log.Debugf("%v", domain)
	conv := streamline.NewConveyorBelt(s, c, &domain)
	conv.Run()

	c.JSON(http.StatusOK, idl.UserLoginResponse{
		Code:    domain.BaseCode,
		Message: "No message, actually",
	})
}