package model

import "github.com/dealmaker/shared/base"

type RequestUploadDomain struct {
	base.Base
	ExpectPrice int32
	RangeValue int32
	InterestTags []string
}
