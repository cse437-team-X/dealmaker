package slice

import "time"

var _logid = 0
func _logIdGen() string {
	_logid ++
	return string(rune(_logid))
}

func TimeMilli() int64 {
	return time.Now().UnixNano()/1000
}