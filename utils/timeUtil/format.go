package timeUtil

import (
	"bytes"
	"fmt"
	"time"
)

func FormatDurationDesc(dur time.Duration) string {
	bf := new(bytes.Buffer)
	if dur.Seconds() < 0 {
		bf.WriteString(fmt.Sprintf("%d秒", int(dur.Seconds())))
		return bf.String()
	}

	days := int(dur.Hours()) / 24
	if days >= 1 {
		bf.WriteString(fmt.Sprintf("%d天", days))
	}
	hours := int(dur.Hours()) % 24
	if hours >= 1 || bf.Len() > 0 {
		bf.WriteString(fmt.Sprintf("%d小时", hours))
	}
	minus := int(dur.Minutes()) % 60
	if minus >= 1 || bf.Len() > 0 {
		bf.WriteString(fmt.Sprintf("%d分", minus))
	}
	secs := int(dur.Seconds()) % 60
	if secs >= 1 || bf.Len() > 0 {
		bf.WriteString(fmt.Sprintf("%d秒", secs))
	}
	return bf.String()
}
