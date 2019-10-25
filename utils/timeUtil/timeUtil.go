package timeUtil

import (
	"time"
)

const CN_LOCATION ="Asia/Chongqing"

func ParseWithCnLocation(layout string, timeStr string) (time.Time, error) {
	if l, err := time.LoadLocation(CN_LOCATION); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, err := time.ParseInLocation(layout, timeStr, l)
		return lt, err
	}
}
