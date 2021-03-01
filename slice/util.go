package slice

import (
	"gitee.com/fat_marmota/streamline"
	"time"
)

func BaseRequestFiller(c *streamline.ConveyorBelt) error {
	data := c.DataPanel.(BaseInterface)
	data.SetBaseTime(time.Now().UnixNano())
	data.SetBaseLogId(_logIdGen())
	return nil
}

var _logid = 0
func _logIdGen() string {
	_logid ++
	return string(rune(_logid))
}
