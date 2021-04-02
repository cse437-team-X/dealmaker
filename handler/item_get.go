package handler

import (
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/procedure/item"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
)

type ItemGetDomain struct {
	base.Base
	model.JwtAuth
	item.ItemGet
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