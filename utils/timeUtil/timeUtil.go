package timeUtil

import (
	"fmt"
	"strconv"
	"time"
)

const CN_LOCATION = "Asia/Chongqing"

func ParseWithCnLocation(layout string, timeStr string) (time.Time, error) {
	if l, err := time.LoadLocation(CN_LOCATION); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, err := time.ParseInLocation(layout, timeStr, l)
		return lt, err
	}
}

func ParseTsToUnixNano(ts string) (int64, error) {
	tslen := len(ts)
	var factor int64 = 0
	//秒：1591343262 长度10
	if tslen == 9 || tslen == 10 || tslen == 11 {
		factor = int64(time.Second)
	} else if tslen == 12 || tslen == 13 || tslen == 14 { //ms
		factor = int64(time.Millisecond)
	} else if tslen == 15 || tslen == 16 || tslen == 17 { //us
		factor = int64(time.Microsecond)
	} else if tslen == 18 || tslen == 19 || tslen == 20 { //ns
		factor = int64(time.Nanosecond)
	} else {
		return 0, fmt.Errorf("wrong timestamp value")
	}
	v, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return 0, err
	}
	return v * factor, nil
}

// ms
func FormatMsTimestampString(ts time.Time) string {
	return strconv.FormatInt(ts.UnixNano()/1000000, 10)
}

// ms
func FormatMsTimestamp(ts time.Time) int64 {
	return ts.UnixNano() / 1000000
}
