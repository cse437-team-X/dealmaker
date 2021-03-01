package slice

type BaseInterface interface {
	GetBaseTime() int64
	SetBaseTime(v int64)
	GetBaseLogId() string
	SetBaseLogId(v string)
	SetBaseCode(v int)
}
