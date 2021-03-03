package model

type ItemUploadDomain struct {
	// Request
	BaseDomain

	UploaderId string
	Price int32
	Discount bool
}
