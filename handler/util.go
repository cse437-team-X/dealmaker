package handler

import (
	"fmt"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

func GenLogMeta(c *streamline.ConveyorBelt) string {
	d := c.DataDomain.(base.BaseInterface)
	res := fmt.Sprintf("%s %s", c.S.Name, d.GetBase().BaseLogId)
	return res
}

func ExecuteStreamline(ctx *gin.Context, streamlineName string, domain interface{}) int {
	s := factory.Factory.Get(streamlineName)
	conv := streamline.NewConveyorBelt(s, ctx, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		ctx.AbortWithStatusJSON(code, domain.(base.BaseInterface).GetBase())
	}
	return code
}