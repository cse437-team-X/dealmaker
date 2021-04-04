package handler

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/factory"
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

func ItemGetHandler(c *gin.Context) {
	s := factory.Factory.Get("/item/get")

	domain := ItemGetDomain{}
	err := c.Bind(&domain)
	if err != nil {
		return
	}
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code, err := conv.Run()
	if err != nil {
		c.AbortWithStatus(code)
		return
	}

	res := make(map[string]interface{})
	res["items"] = domain.Result
	c.JSON(code, res)
}