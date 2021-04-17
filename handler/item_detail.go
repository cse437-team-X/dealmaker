package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ItemDetailDomain struct {
	base.Base
	model2.GetItemDomain
}

type ItemDetailResponse struct {
	Message string
	Items []model2.Item
}

func ItemDetailHandler(c *gin.Context) {
	input := c.Query("obj_id")

	domain := ItemDetailDomain{
		GetItemDomain: model2.GetItemDomain{
			QueryFilter: model2.QueryFilter{ObjId: input},
		},
	}

	s := factory.Factory.Get("/item/detail")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	resp := ItemDetailResponse{
		Message: domain.BaseMessage,
		Items:   domain.Result,
	}
	c.JSON(code, resp)
}