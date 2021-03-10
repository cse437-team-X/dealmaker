package handler

import (
	"fmt"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/shared/base"
)

func GenLogMeta(c *streamline.ConveyorBelt) string {
	d := c.DataDomain.(base.BaseInterface)
	res := fmt.Sprintf("%s %s", c.S.Name, d.GetLogId())
	return res
}