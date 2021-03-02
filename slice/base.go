package slice

import (
	"gitee.com/fat_marmota/streamline"
	"net/http"
)

type BaseInterface interface {
	GetBaseTime() int64
	SetBaseTime(v int64)
	GetBaseLogId() string
	SetBaseLogId(v string)
}

func BaseRequestFiller(c *streamline.ConveyorBelt) int {
	data := c.DataPanel.(BaseInterface)
	data.SetBaseTime(TimeMilli())
	data.SetBaseLogId(_logIdGen())
	c.Logger.Debugf("Test")
	return http.StatusOK
}