package util

import (
	"time"
)

func GetCurrentTimeMilis() int {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	return int(millis)
}
