package model

type BaseDomain struct {
	BaseTime int64
	BaseCode int
	BaseLogId string
}

func (b *BaseDomain) GetBaseTime() int64 {
	return b.BaseTime
}
func (b *BaseDomain) SetBaseTime(v int64) {
	b.BaseTime = v
}
func (b *BaseDomain) GetBaseLogId() string {
	return b.BaseLogId
}
func (b *BaseDomain) SetBaseLogId(v string) {
	b.BaseLogId = v
}
func (b *BaseDomain) SetBaseCode(v int) {
	b.BaseCode = v
}