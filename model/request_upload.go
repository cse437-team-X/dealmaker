package model

import "github.com/dealmaker/shared/base"

type RequestUploadDomain struct {
	base.BaseDomain
	ExpectPrice int32
	RangeValue int32
	InterestTags []string
}
