package item_upload

import (
	"github.com/dealmaker/base_model"
	"github.com/dealmaker/base_model/obj"
)

type ItemUploadDomain struct {
	base_model.BaseDomain
	obj.Item
	//base_model.AuthedUserDomain

	Price int32
	Discount bool
}
