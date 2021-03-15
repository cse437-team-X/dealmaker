package base

import (
	"fmt"
	"gitee.com/fat_marmota/streamline"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BaseInterface interface {
	GetTime() int64
	SetTime(int64)
	GetLogId() string
	SetLogId(string)
}

const MAX_RAND_VAL = 8192

func genLogId(c *streamline.ConveyorBelt, t time.Time, ip string) string {
	timeStr := t.Format("20060102150405")
	ipStrip := strings.ReplaceAll(ip, ".", "")
	if ipStrip == "::1" {
		ipStrip = "127001"
	}
	c.Debugw("ip", ip, "ips", ipStrip)
	ipNum, err := strconv.ParseInt(ipStrip,10,64)
	if err != nil {
		c.Errorw("err", err.Error())
	}
	rnd := rand.Intn(MAX_RAND_VAL)
	ipHex := strconv.FormatInt(ipNum, 16)
	rndHex := strconv.FormatInt(int64(rnd), 16)

	logid := fmt.Sprintf("%s%s%s", timeStr, ipHex, rndHex)
	c.Debugw("logid", logid)
	return logid
}

func BaseRequestFiller(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BaseInterface)
	reqTime := time.Now()
	data.SetTime(reqTime.Unix())
	data.SetLogId(genLogId(c, reqTime, "fakeip"))//c.Ctx.(*gin.Context).ClientIP()))

	return http.StatusOK
}