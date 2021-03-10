package base

type BaseDomain struct {
	BaseTime int64
	BaseLogId string
	BaseSessionId string
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

func (b *BaseDomain) GetSessionId() string {
	return b.BaseSessionId
}
func (b *BaseDomain) SetSessionId(s string) {
	b.BaseSessionId = s
}
