package stringUtil

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	B  int64 = 1
	KB       = 1 << 10
	MB       = 1 << 20
	GB       = 1 << 30
	TB       = 1 << 40
	PB       = 1 << 50
)

type SizeParser struct {
	// 使用十进制还是二进制
	UseBinaryUnits bool
}

func NewSizeParser(useBinary bool) *SizeParser {
	return &SizeParser{UseBinaryUnits: useBinary}
}
func (p *SizeParser) Parse(sizeStr string) (int64, error) {
	sizeStr = strings.TrimSpace(sizeStr)
	if sizeStr == "" {
		return 0, fmt.Errorf("空字符串")
	}

	// 分离数字和单位
	var numStr, unitStr string
	for i, r := range sizeStr {
		if unicode.IsDigit(r) || r == '.' {
			numStr += string(r)
		} else {
			unitStr = strings.ToUpper(strings.TrimSpace(sizeStr[i:]))
			break
		}
	}

	if numStr == "" {
		return 0, fmt.Errorf("未找到数字部分")
	}

	// 解析数字
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, fmt.Errorf("解析数字失败: %v", err)
	}

	// 处理单位
	if unitStr == "" {
		unitStr = "B"
	}

	var multiplier int64
	if p.UseBinaryUnits {
		// 二进制单位
		multiplier = p.getBinaryMultiplier(unitStr)
	} else {
		// 十进制单位
		multiplier = p.getDecimalMultiplier(unitStr)
	}

	if multiplier == 0 {
		return 0, fmt.Errorf("未知的单位: %s", unitStr)
	}

	return int64(num * float64(multiplier)), nil
}
func (p *SizeParser) getBinaryMultiplier(unit string) int64 {
	switch unit {
	case "B":
		return B
	case "K", "KB", "KI", "KIB":
		return KB
	case "M", "MB", "MI", "MIB":
		return MB
	case "G", "GB", "GI", "GIB":
		return GB
	case "T", "TB", "TI", "TIB":
		return TB
	case "P", "PB", "PI", "PIB":
		return PB
	default:
		return 0
	}
}
func (p *SizeParser) getDecimalMultiplier(unit string) int64 {
	switch unit {
	case "B":
		return B
	case "K", "KB":
		return 1000
	case "M", "MB":
		return 1000000
	case "G", "GB":
		return 1000000000
	case "T", "TB":
		return 1000000000000
	case "P", "PB":
		return 1000000000000000
	default:
		return 0
	}
}
