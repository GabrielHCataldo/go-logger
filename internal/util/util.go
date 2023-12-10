package util

import (
	"fmt"
	"math/rand"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetCallerInfo(skipCaller int) (fileName string, line string, funcName string) {
	pc := make([]uintptr, 1)
	runtime.Callers(skipCaller, pc)
	funcInfo := runtime.FuncForPC(pc[0])
	file, lineInt := funcInfo.FileLine(pc[0])
	fileBase := path.Base(file)
	name := formatFuncName(funcInfo.Name())
	return fileBase, strconv.Itoa(lineInt), name
}

func GetJsonNameByTag(tag string) string {
	result := ""
	splitTag := strings.Split(tag, ",")
	if len(splitTag) > 0 {
		result = splitTag[0]
	}
	if len(result) == 0 || result == "-" {
		return ""
	}
	return result
}

func ContainsJsonOmitemptyByTag(tag string) bool {
	return strings.Contains(tag, "omitempty")
}

func ConvertToString(a any) string {
	if a == nil {
		return ""
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if !v.CanInterface() {
		return ""
	}
	return convertValueBasedOnType(v.Interface())
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func IsNilValueReflect(fieldValue reflect.Value) bool {
	if (fieldValue.Kind() == reflect.Interface ||
		fieldValue.Kind() == reflect.Pointer ||
		fieldValue.Kind() == reflect.Slice ||
		fieldValue.Kind() == reflect.Chan ||
		fieldValue.Kind() == reflect.Func ||
		fieldValue.Kind() == reflect.UnsafePointer ||
		fieldValue.Kind() == reflect.Map) && fieldValue.IsNil() {
		return true
	}
	return false
}

func formatFuncName(name string) string {
	name = path.Base(name)
	splitName := strings.Split(name, ".")
	if len(splitName) >= 3 {
		name = strings.Replace(name, splitName[0]+".", "", 1)
	}
	return name
}

func convertValueBasedOnType(i any) (output string) {
	switch t := i.(type) {
	case int:
		return strconv.Itoa(t)
	case int32:
		return strconv.Itoa(int(t))
	case int64:
		return strconv.Itoa(int(t))
	case bool:
		return strconv.FormatBool(t)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case time.Time:
		return t.Format(time.RFC3339)
	case string:
		return t
	default:
		return fmt.Sprintf(`"%s"`, t)
	}
}
