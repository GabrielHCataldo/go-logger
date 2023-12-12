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
	if len(result) == 0 || result == "omitempty" {
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

func IsNilValueReflect(v reflect.Value) bool {
	if (v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Pointer ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func ||
		v.Kind() == reflect.UnsafePointer ||
		v.Kind() == reflect.Map) && v.IsNil() {
		return true
	}
	return false
}

func IsZeroReflect(v reflect.Value) bool {
	return v.IsZero() || IsNilValueReflect(v) ||
		(v.Kind() == reflect.Map && len(v.MapKeys()) == 0) ||
		(v.Kind() == reflect.Struct && v.NumField() == 0) ||
		(v.Kind() == reflect.Slice && v.Len() == 0) ||
		(v.Kind() == reflect.Array && v.Len() == 0)
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
	case error:
		return t.Error()
	default:
		return fmt.Sprintf(`%s`, t)
	}
}
