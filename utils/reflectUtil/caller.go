package reflectUtil

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func FormatCallerLineKey(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	baseFile := filepath.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	if idx := strings.LastIndex(funcName, "."); idx != -1 {
		funcName = funcName[idx+1:]
	}
	sb := strings.Builder{}
	sb.Grow(96)
	sb.WriteString(baseFile)
	sb.WriteString("#")
	sb.WriteString(funcName)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(line))
	return sb.String()
}
