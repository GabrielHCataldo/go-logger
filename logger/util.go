package logger

import (
	"path"
	"runtime"
	"strconv"
	"strings"
)

func getSystemCaller(skipCaller int) (fileName string, line string, funcName string) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(skipCaller, pc)
	f := runtime.FuncForPC(pc[0])
	file, lineInt := f.FileLine(pc[0])
	fileBase := path.Base(file)
	nameFunc := path.Base(f.Name())
	splitFunc := strings.Split(nameFunc, ".")
	if len(splitFunc) >= 3 {
		nameFunc = strings.Replace(nameFunc, splitFunc[0]+".", "", 1)
	}
	return fileBase, strconv.Itoa(lineInt), nameFunc
}
