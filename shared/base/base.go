package base

import (
	"fmt"
	"gitee.com/fat_marmota/streamline"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BaseInterface interface {
	GetBaseTime() int64
	SetBaseTime(int64)
	GetBaseLogId() string
	SetBaseLogId(string)
	SetSessionId(string)
	GetSessionId() string
}

const MAX_RAND_VAL = 8192

func genLogId(c *streamline.ConveyorBelt, t time.Time, ip string) string {
	timeStr := time.Now().Format("20060102150405")
	ipStrip := strings.ReplaceAll(ip, ".", "")
	ipNum, err := strconv.ParseInt(ipStrip,10,64)
	if err != nil {
		c.Logger.Errorw("Generate logid", "err", err.Error())
	}
	rnd := rand.Intn(MAX_RAND_VAL)
	ipHex := strconv.FormatInt(ipNum, 16)
	rndHex := strconv.FormatInt(int64(rnd), 16)

	logid := fmt.Sprintf("%s%s%s", timeStr, ipHex, rndHex)
	return logid
}

func BaseRequestFiller(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BaseInterface)
	reqTime := time.Now()
	data.SetBaseTime(reqTime.Unix())
	data.SetBaseLogId(genLogId(c, reqTime, c.Ctx.(*gin.Context).ClientIP()))

	//session := sessions.Default(c.Ctx.(*gin.Context))
	//id:=session.Get(SessionId)
	//SessionIdGen()
	//if id == nil {
	//	session.Set(SessionId, SessionIdGen())
	//	session.Save()
	//}
	//c.Logger.Debugw("Session test", SessionId, id)

	return http.StatusOK
}