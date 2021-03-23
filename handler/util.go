package handler

import (
	"fmt"
	"github.com/itzmeerkat/streamline"
	"github.com/dealmaker/shared/base"
)

func GenLogMeta(c *streamline.ConveyorBelt) string {
	d := c.DataDomain.(base.BaseInterface)
	res := fmt.Sprintf("%s %s", c.S.Name, d.GetBase().BaseLogId)
	return res
}