package item_upload

import (
	"github.com/dealmaker/model/obj"
	"github.com/dealmaker/shared/base"
)

type ItemUploadDomain struct {
	base.BaseDomain
	obj.Item
	//base_model.AuthedUserDomain

	Price int32
	Discount bool
}
