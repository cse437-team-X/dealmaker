package slice

import (
	"crypto"
	"fmt"
	"gitee.com/fat_marmota/infra/log"
	"time"
)

var _logid = 0

func _logIdGen() string {
	_logid ++
	return string(rune(_logid))
}

func TimeMilli() int64 {
	return time.Now().UnixNano()/1000
}

func SessionIdGen() string {
	// dealmaker to SALT
	str := "dea1m4k3r" + "::" + fmt.Sprint(TimeMilli())
	hasher := crypto.SHA1.New()
	hasher.Write([]byte(str))
	res := fmt.Sprintf("%x", hasher.Sum(nil))
	log.Debugf("Gen session %v", res)
	return res
}