package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ItemDeleteDomain struct {
	base.Base
	model2.ItemUpdate
	model.JwtAuth
}

type ItemDeleteInput struct {
	Token string
	ObjId string
}

type ItemDeleteResponse struct {
	Message string
}

func ItemDeleteHandler(c *gin.Context) {
	input := ItemDeleteInput{}

	c.Bind(&input)

	domain := ItemDeleteDomain{
		ItemUpdate: model2.ItemUpdate{ObjId: input.ObjId},
		JwtAuth: model.JwtAuth{
			Token: input.Token,
		},
	}

	s := factory.Factory.Get("/item/delete")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	resp := ItemGetResponse{
		Message: domain.BaseMessage,
	}
	c.JSON(code, resp)
}