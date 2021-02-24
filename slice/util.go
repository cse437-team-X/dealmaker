package slice

import (
	"gitee.com/fat_marmota/streamline"
	"time"
)

type FillBaseRequest interface {
	SetBaseTime(v int64)
}

type FillBaseResponse interface {
	SetBaseLogId(v string)
}

func BaseRequestFiller(self *streamline.Streamline, in interface{}, out interface{}) error {
	indata := in.(FillBaseRequest)
	indata.SetBaseTime(time.Now().UnixNano())
	outdata := out.(FillBaseResponse)
	outdata.SetBaseLogId(_logIdGen())
	return nil
}

var _logid = 0
func _logIdGen() string {
	_logid ++
	return string(rune(_logid))
}
