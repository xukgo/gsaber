package stringUtil

import (
	"fmt"
	"time"
)

func FormatHumanizeFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func FormatHumanizeSpeed(size int64, duration time.Duration) string {
	speed := float64(size) / duration.Seconds()
	const unit = 1024
	if speed < unit {
		return fmt.Sprintf("%.0f B/s", speed)
	}
	div, exp := float64(unit), 0
	for n := speed / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB/s", speed/div, "KMGTPE"[exp])
}
