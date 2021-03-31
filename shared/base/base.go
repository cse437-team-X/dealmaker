package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type BaseInterface interface {
	GetBase() *Base
}

const MAX_RAND_VAL = 8192

func genLogId(c *streamline.ConveyorBelt, t time.Time) string {
	timeStr := t.Format("20060102150405")

	rnd := rand.Intn(MAX_RAND_VAL)
	rndHex := strconv.FormatInt(int64(rnd), 16)

	logid := fmt.Sprintf("%s%s", timeStr, rndHex)
	c.Debugw("logid", logid)
	return logid
}

func BaseRequestFiller(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BaseInterface).GetBase()
	reqTime := time.Now()
	data.BaseTime = reqTime.Unix()
	data.BaseLogId = genLogId(c, reqTime)
	c.Ctx.(*gin.Context).Header("log-id", data.BaseLogId)
	return http.StatusOK
}