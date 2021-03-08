package slice

import (
	"gitee.com/fat_marmota/streamline"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseInterface interface {
	GetBaseTime() int64
	SetBaseTime(int64)
	GetBaseLogId() string
	SetBaseLogId(string)
	SetSessionId(string)
	GetSessionId() string
}

func BaseRequestFiller(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BaseInterface)
	data.SetBaseTime(TimeMilli())
	data.SetBaseLogId(_logIdGen())

	session := sessions.Default(c.Ctx.(*gin.Context))
	id:=session.Get(SessionId)
	//SessionIdGen()
	if id == nil {
		session.Set(SessionId, SessionIdGen())
		session.Save()
	}
	c.Logger.Debugw("Session test", SessionId, id)

	return http.StatusOK
}