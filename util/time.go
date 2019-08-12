package util

import "time"

// 得到当前时间的时间戳
func GetCurrTs() int64{
	return time.Now().Unix()
}