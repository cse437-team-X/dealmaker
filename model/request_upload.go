package model

type RequestUploadDomain struct {
	BaseDomain
	ExpectPrice int32
	RangeValue int32
	InterestTags []string
}
