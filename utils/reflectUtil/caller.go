package reflectUtil

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type CallerLayerInfo struct {
	FileName string
	Function string
	Line     int
	Ok       bool
}

func GetCallerLayerInfo(skip int) CallerLayerInfo {
	result := CallerLayerInfo{}
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		result.Ok = ok
		return result
	}
	baseFile := filepath.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	if idx := strings.LastIndex(funcName, "."); idx != -1 {
		funcName = funcName[idx+1:]
	}
	result.FileName = baseFile
	result.Function = funcName
	result.Line = line
	result.Ok = ok
	return result
}

func FormatCallerLineKey(skip int) string {
	callerInfo := GetCallerLayerInfo(skip + 1)

	sb := strings.Builder{}
	sb.Grow(96)
	sb.WriteString(callerInfo.FileName)
	sb.WriteString("#")
	sb.WriteString(callerInfo.Function)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(callerInfo.Line))
	return sb.String()
}
