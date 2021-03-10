package item_upload

import (
	"github.com/dealmaker/model/obj"
	"github.com/dealmaker/shared/base"
)

type ItemUploadDomain struct {
	base.Base
	obj.Item

	Price int32
	Discount bool
}
