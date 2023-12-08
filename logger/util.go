package logger

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
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

func getJsonNameByTag(tag string) string {
	result := tag
	splitTag := strings.Split(tag, ",")
	if len(splitTag) > 0 {
		result = splitTag[0]
	}
	if len(result) == 0 || result == "-" {
		return ""
	}
	return result
}

func convertToString(v any) string {
	if v == nil {
		return ""
	}
	tr := reflect.TypeOf(v)
	vr := reflect.ValueOf(v)
	if !vr.IsValid() || vr.IsZero() {
		return ""
	}
	if tr.Kind() == reflect.Pointer {
		vr = vr.Elem()
	}
	if !vr.CanInterface() {
		return ""
	}
	switch vr.Interface().(type) {
	case int:
		return strconv.Itoa(v.(int))
	case int32:
		return strconv.Itoa(int(v.(int32)))
	case int64:
		return strconv.Itoa(int(v.(int64)))
	case bool:
		return strconv.FormatBool(v.(bool))
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case time.Time:
		return v.(time.Time).Format(DateFormatFull24h.Format())
	case string:
		return v.(string)
	default:
		return fmt.Sprintf(`"%s"`, v)
	}
}
