package logger

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

var logInfo = log.New(os.Stdout, "[ASAAS \u001b[34mINFO: \u001B[0m", log.LstdFlags)
var logWarning = log.New(os.Stdout, "[ASAAS \u001b[33mWARNING: \u001B[0m", log.LstdFlags)
var logError = log.New(os.Stdout, "[ASAAS \u001b[31mERROR: \u001b[0m", log.LstdFlags)
var logDebug = log.New(os.Stdout, "[DEBUG: \u001B[0m", 0)

var opt *options

type options struct {
	Mode                Mode
	DateFormat          DateFormat
	UTC                 bool
	HideAllArgs         bool
	HideArgDatetime     bool
	HideArgCaller       bool
	DisablePrefixColors bool
}

type logJson struct {
	Level    string `json:"level,omitempty"`
	Datetime string `json:"datetime,omitempty"`
	File     string `json:"file,omitempty"`
	Func     string `json:"func,omitempty"`
	Line     string `json:"line,omitempty"`
	Msg      string `json:"msg,omitempty"`
}

func SetOptions(options *options) {
	opt = options
}

func Info(v ...any) {
	//print(levelInfo, 2, v...)
}

func InfoSkipCaller(skipCaller int, v ...any) {
	logInfo.Print(fmt.Sprintln(v...))
}

func Warning(v ...any) {
	logWarning.Print(fmt.Sprintln(v...))
}

func WarningSkipCaller(skipCaller int, v ...any) {
	logWarning.Print(fmt.Sprintln(v...))
}

func Debug(v ...any) {
	printLog(levelDebug, 3, v...)
	//logDebug.Print(fmt.Sprintln(v...))
}

func DebugSkipCaller(skipCaller int, v ...any) {
	logDebug.Print(fmt.Sprintln(v...))
}

func Error(v ...any) {
	logError.Print(fmt.Sprintln(v...))
}

func Errorf(format string, v ...any) {
	logError.Print(fmt.Sprintf(format, v...))
}

func ErrorSkipCaller(skipCaller int, v ...any) {
	logError.Print(fmt.Sprintln(v...))
}

func printLog(l level, skipCaller int, v ...any) {
	logg(l, skipCaller+1).Println(prepareMsg(l, skipCaller, v...)...)
}

func logg(l level, skipCaller int) *log.Logger {
	switch opt.Mode {
	case ModeJson:
		return log.New(os.Stdout, "", 0)
	}
	return loggNormal(l, skipCaller+1)
}

func loggNormal(l level, skipCaller int) *log.Logger {
	var prefix string
	if !opt.HideAllArgs && !opt.HideArgDatetime {
		prefix += "["
	}
	prefix += getArgLogLevel(l)
	if !opt.HideAllArgs {
		if !opt.HideArgDatetime {
			prefix += " " + getArgDatetime() + "]"
		}
		if !opt.HideArgCaller {
			prefix += " " + getArgCaller(skipCaller+1) + ":"
		}
	} else {
		prefix += ":"
	}
	prefix += " "
	return log.New(os.Stdout, prefix, 0)
}

func prepareMsg(l level, skipCaller int, vs ...any) []any {
	var msg []any
	for _, v := range vs {
		var vr any
		t := reflect.TypeOf(v)
		vf := reflect.ValueOf(v)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
			vr = vf.Elem()
		}
		switch t.Kind() {
		case reflect.Struct, reflect.Map:
			vr = prepareStructMsg(t, vf)
			break
		case reflect.Slice, reflect.Array:
			vr = prepareSliceMsg(vf, "")
			break
		default:
			vr = v
		}
		msg = append(msg, vr)
	}
	switch opt.Mode {
	case ModeJson:
		return prepareMsgJson(l, skipCaller+1, msg...)
	}
	return msg
}

func prepareMsgJson(l level, skipCaller int, v ...any) []any {
	var vr any
	jLog := logJson{
		Level: l.String(),
		Msg:   strings.Replace(fmt.Sprintln(v...), "\n", "", -1),
	}
	if !opt.HideAllArgs {
		if !opt.HideArgDatetime {
			jLog.Datetime = getArgDatetime()
		}
		if !opt.HideArgCaller {
			fileName, line, funcName := getSystemCaller(skipCaller + 1)
			jLog.File = fileName
			jLog.Func = funcName
			jLog.Line = line
		}
	}
	b, err := json.Marshal(jLog)
	if err == nil {
		vr = strings.ReplaceAll(string(b), "\\", "")
	} else {
		vr = jLog
	}
	return []any{vr}
}

func prepareStructMsg(t reflect.Type, v reflect.Value) any {
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		if f.IsZero() || !f.IsValid() || (f.Kind() == reflect.Pointer && f.IsNil()) {
			continue
		}
		k := getJsonNameByTag(ft.Tag.Get("json"))
		fv := reflect.ValueOf(getValueByReflect(ft.Type, f, ft.Tag.Get("logger")))
		if fv.IsValid() && !fv.IsZero() && len(k) != 0 {
			result.Set(k, fv.String())
		}
	}
	b, err := json.Marshal(result)
	if err != nil {
		return v.Interface()
	}
	return string(b)
}

func prepareSliceMsg(v reflect.Value, loggerTag string) any {
	var result []string
	for i := 0; i < v.Len(); i++ {
		f := v.Index(i)
		if f.IsZero() || !f.IsValid() || (f.Kind() == reflect.Pointer && f.IsNil()) {
			continue
		} else if f.Kind() == reflect.Pointer {
			f = f.Elem()
		}
		fv := reflect.ValueOf(getValueByReflect(f.Type(), f, loggerTag))
		if fv.IsValid() && !fv.IsZero() {
			result = append(result, fv.String())
		}
	}
	if len(result) == 0 {
		b, err := json.Marshal(result)
		if err != nil {
			return v.Interface()
		}
		return string(b)
	}
	return strings.Join(result[:], ", ")
}

func getValueByReflect(t reflect.Type, v reflect.Value, loggerTag string) any {
	if !v.CanInterface() {
		return nil
	}
	switch v.Interface().(type) {
	case time.Time:
		return convertStringReflectValue(v, loggerTag)
	}
	switch v.Kind() {
	case reflect.Struct, reflect.Map:
		return prepareStructMsg(t, v)
	case reflect.Array, reflect.Slice:
		return prepareSliceMsg(v, loggerTag)
	}
	return convertStringReflectValue(v, loggerTag)
}

func convertStringReflectValue(v reflect.Value, loggerTag string) string {
	s := convertToString(v.Interface())
	if len(s) > 0 {
		if strings.Contains(loggerTag, "hide") {
			vm := ""
			for i := 0; i < len(s); i++ {
				vm += "*"
			}
			s = vm
		} else if strings.Contains(loggerTag, "mask_start") {
			if len(s) == 1 {
				s = "*"
			} else {
				vm := ""
				for i := 0; i < len(s)/2; i++ {
					vm += "*"
				}
				vm += s[len(s)/2:]
				s = vm
			}
		} else if strings.Contains(loggerTag, "mask_end") {
			if len(s) == 1 {
				s = "*"
			} else {
				vm := s[:len(s)/2]
				for i := 0; i < len(s)/2; i++ {
					vm += "*"
				}
				s = vm
			}
		}
	}
	return s
}

func getArgLogLevel(l level) string {
	var result string
	if !opt.DisablePrefixColors {
		result = l.Color()
	}
	result += l.String() + "\u001B[0m"
	return result
}

func getArgDatetime() string {
	t := time.Now().Local()
	if opt.UTC {
		t = time.Now().UTC()
	}
	return t.Format(opt.DateFormat.Format())
}

func getArgCaller(skipCaller int) string {
	fileName, line, _ := getSystemCaller(skipCaller)
	return fileName + ":" + line
}
