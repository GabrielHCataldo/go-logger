package logger

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
	"go-logger/internal/util"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

var opts *options

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
	opts = options
}

func initOptions() {
	if opts == nil {
		opts = &options{}
	}
}

// Info is a function that records informational messages based on the global options variable, using the level INFO.
//
// Usage Example:
//
// logger.Info("test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Info(v ...any) {
	printLog(levelInfo, 3, v...)
}

// InfoH is a function that records informational messages based on the global options variable, using the INFO level,
// different from Info, this function will hide all values passed as parameters (v ...any).
//
// Usage Example:
//
// logger.InfoH("test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func InfoH(v ...any) {
	printLogH(levelInfo, 3, v...)
}

// InfoMS is a function that logs informational messages based on the global options variable, using the INFO level,
// unlike Info, this function will hide the initial half of the values passed as parameters (v ...any).
//
// Usage Example:
//
// logger.InfoMS("test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func InfoMS(v ...any) {
	printLogMS(levelInfo, 3, v...)
}

// InfoME is a function that logs informational messages based on the global options variable, using the INFO level,
// unlike Info, this function will hide the end half of the values passed as parameters (v ...any).
//
// Usage Example:
//
// logger.InfoME("test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func InfoME(v ...any) {
	printLogME(levelInfo, 3, v...)
}

func InfoSkipCaller(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, v...)
}

func InfoOptsH(optsArg options, v ...any) {
	opts = &optsArg
	printLogH(levelInfo, 3, v...)
}

func InfoOptsMS(optsArg options, v ...any) {
	opts = &optsArg
	printLogMS(levelInfo, 3, v...)
}

func InfoOptsME(optsArg options, v ...any) {
	opts = &optsArg
	printLogME(levelInfo, 3, v...)
}

func InfoSkipCallerH(skipCaller int, v ...any) {
	printLogH(levelInfo, skipCaller, v...)
}

func InfoSkipCallerMS(skipCaller int, v ...any) {
	printLogMS(levelInfo, skipCaller, v...)
}

func InfoSkipCallerME(skipCaller int, v ...any) {
	printLogME(levelInfo, skipCaller, v...)
}

func Warning(v ...any) {
	printLog(levelWarning, 3, v...)
}

func WarningH(v ...any) {
	printLogH(levelWarning, 3, v...)
}

func WarningMS(v ...any) {
	printLogMS(levelWarning, 3, v...)
}

func WarningME(v ...any) {
	printLogME(levelWarning, 3, v...)
}

func WarningSkipCaller(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, v...)
}

func WarningOptsH(optsArg options, v ...any) {
	opts = &optsArg
	printLogH(levelWarning, 3, v...)
}

func WarningOptsMS(optsArg options, v ...any) {
	opts = &optsArg
	printLogMS(levelWarning, 3, v...)
}

func WarningOptsME(optsArg options, v ...any) {
	opts = &optsArg
	printLogME(levelWarning, 3, v...)
}

func WarningSkipCallerH(skipCaller int, v ...any) {
	printLogH(levelWarning, skipCaller, v...)
}

func WarningSkipCallerMS(skipCaller int, v ...any) {
	printLogMS(levelWarning, skipCaller, v...)
}

func WarningSkipCallerME(skipCaller int, v ...any) {
	printLogME(levelWarning, skipCaller, v...)
}

func Debug(v ...any) {
	printLog(levelDebug, 3, v...)
}

func DebugH(v ...any) {
	printLogH(levelDebug, 3, v...)
}

func DebugMS(v ...any) {
	printLogMS(levelDebug, 3, v...)
}

func DebugME(v ...any) {
	printLogME(levelDebug, 3, v...)
}

func DebugSkipCaller(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, v...)
}

func DebugOptsH(optsArg options, v ...any) {
	opts = &optsArg
	printLogH(levelDebug, 3, v...)
}

func DebugOptsMS(optsArg options, v ...any) {
	opts = &optsArg
	printLogMS(levelDebug, 3, v...)
}

func DebugOptsME(optsArg options, v ...any) {
	opts = &optsArg
	printLogME(levelDebug, 3, v...)
}

func DebugSkipCallerH(skipCaller int, v ...any) {
	printLogH(levelDebug, skipCaller, v...)
}

func DebugSkipCallerMS(skipCaller int, v ...any) {
	printLogMS(levelDebug, skipCaller, v...)
}

func DebugSkipCallerME(skipCaller int, v ...any) {
	printLogME(levelDebug, skipCaller, v...)
}

func Error(v ...any) {
	printLog(levelError, 3, v...)
}

func ErrorF(format string, v ...any) {
	printLogF(levelError, 3, format, v...)
}

func ErrorH(v ...any) {
	printLogH(levelError, 3, v...)
}

func ErrorMS(v ...any) {
	printLogMS(levelError, 3, v...)
}

func ErrorME(v ...any) {
	printLogME(levelError, 3, v...)
}

func ErrorSkipCaller(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, v...)
}

func ErrorOptsH(optsArg options, v ...any) {
	opts = &optsArg
	printLogH(levelError, 3, v...)
}

func ErrorOptsMS(optsArg options, v ...any) {
	opts = &optsArg
	printLogMS(levelError, 3, v...)
}

func ErrorOptsME(optsArg options, v ...any) {
	opts = &optsArg
	printLogME(levelError, 3, v...)
}

func ErrorSkipCallerH(skipCaller int, v ...any) {
	printLogH(levelError, skipCaller, v...)
}

func ErrorSkipCallerMS(skipCaller int, v ...any) {
	printLogMS(levelError, skipCaller, v...)
}

func ErrorSkipCallerME(skipCaller int, v ...any) {
	printLogME(levelError, skipCaller, v...)
}

func printLog(lvl level, skipCaller int, v ...any) {
	getLogger(lvl, skipCaller+1).Println(prepareMsg(lvl, skipCaller, "", v...)...)
}

func printLogH(lvl level, skipCaller int, v ...any) {
	getLogger(lvl, skipCaller+1).Println(prepareMsg(lvl, skipCaller, "hide", v...)...)
}

func printLogMS(lvl level, skipCaller int, v ...any) {
	getLogger(lvl, skipCaller+1).Println(prepareMsg(lvl, skipCaller, "mask_start", v...)...)
}

func printLogME(lvl level, skipCaller int, v ...any) {
	getLogger(lvl, skipCaller+1).Println(prepareMsg(lvl, skipCaller, "mask_end", v...)...)
}

func printLogF(lvl level, skipCaller int, format string, v ...any) {
	getLogger(lvl, skipCaller+1).Printf(format, prepareMsg(lvl, skipCaller, "", v...)...)
}

func getLogger(lvl level, skipCaller int) *log.Logger {
	initOptions()
	if opts.Mode == ModeJson {
		return log.New(os.Stdout, "", 0)
	}
	return log.New(os.Stdout, getLoggerNormalPrefix(lvl, skipCaller), 0)
}

func buildDatetimeString() string {
	if opts.HideAllArgs || opts.HideArgDatetime {
		return ""
	}

	return " " + getArgDatetime() + "]"
}

func getLoggerNormalPrefix(lvl level, skipCaller int) string {
	var b strings.Builder
	if !opts.HideAllArgs && !opts.HideArgDatetime {
		b.WriteString("[")
	}
	datetimeString := buildDatetimeString()
	b.WriteString(getArgLogLevel(lvl))
	b.WriteString(datetimeString)
	if !opts.HideAllArgs && !opts.HideArgCaller {
		b.WriteString(" " + getArgCaller(skipCaller+1) + ":")
	} else if datetimeString == "" {
		b.WriteString(":")
	}
	b.WriteString(" ")
	return b.String()
}

func prepareMsg(lvl level, skipCaller int, tag string, msgContents ...any) []any {
	var processedMsg []any
	for _, msgContent := range msgContents {
		valueType := reflect.TypeOf(msgContent)
		value := reflect.ValueOf(msgContent)
		processedValue := processMsgValue(valueType, value, tag)
		processedMsg = append(processedMsg, processedValue)
	}
	if opts.Mode == ModeJson {
		return prepareModeJsonMsg(lvl, skipCaller+1, processedMsg...)
	}
	return processedMsg
}

func processMsgValue(valueType reflect.Type, value reflect.Value, tag string) any {
	var processedValue any
	if valueType.Kind() == reflect.Pointer {
		valueType = valueType.Elem()
		processedValue = value.Elem()
	}
	switch valueType.Kind() {
	case reflect.Struct:
		processedValue = prepareStructMsg(valueType, value, false, tag)
	case reflect.Map:
		processedValue = prepareMapMsg(value, false, tag)
	case reflect.Slice, reflect.Array:
		processedValue = prepareSliceMsg(value, false, tag)
	default:
		processedValue = value
	}
	return processedValue
}

func getLogJson(lvl level, skipCaller int, v ...any) logJson {
	lg := logJson{
		Level: lvl.String(),
		Msg:   strings.Replace(fmt.Sprintln(v...), "\n", "", -1),
	}
	if opts.HideAllArgs {
		return lg
	}
	if !opts.HideArgDatetime {
		lg.Datetime = getArgDatetime()
	}
	if !opts.HideArgCaller {
		fileName, line, funcName := util.GetCallerInfo(skipCaller + 1)
		lg.File = fileName
		lg.Func = funcName
		lg.Line = line
	}
	return lg
}

func prepareModeJsonMsg(lvl level, skipCaller int, v ...any) []any {
	var processedMsg any
	lgJson := getLogJson(lvl, skipCaller, v...)
	bytes, err := json.Marshal(lgJson)
	if err == nil {
		processedMsg = strings.ReplaceAll(string(bytes), "\\", "")
	} else {
		processedMsg = lgJson
	}
	return []any{processedMsg}
}

func prepareStructMsg(t reflect.Type, v reflect.Value, sub bool, tag string) any {
	if x, ok := v.Interface().(time.Time); ok {
		return x.Format(time.RFC3339)
	}
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		fieldTag := tag
		jsonName := util.GetJsonNameByTag(fieldType.Tag.Get("json"))
		if util.IsNilValueReflect(fieldValue) {
			if !util.ContainsJsonOmitemptyByTag(fieldType.Tag.Get("json")) {
				result.Set(jsonName, nil)
			}
			continue
		}
		if len(fieldTag) == 0 {
			fieldTag = fieldType.Tag.Get("logger")
		}
		fieldValueProcessed := getFieldValue(fieldType.Type, fieldValue, fieldTag)
		if len(jsonName) != 0 && fieldValueProcessed != nil {
			result.Set(jsonName, fieldValueProcessed)
		}
	}
	if sub {
		return result.Values()
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return v.Interface()
	}
	return string(bytes)
}

func prepareMapMsg(v reflect.Value, sub bool, tag string) any {
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for _, key := range v.MapKeys() {
		if util.IsNilValueReflect(key) {
			continue
		}
		mKey := key.Convert(v.Type().Key())
		mValue := v.MapIndex(mKey)
		if mKey.IsZero() || util.IsNilValueReflect(mKey) {
			continue
		}
		mKeyString := util.ConvertToString(mKey.Interface())
		if len(mKeyString) != 0 && util.IsNilValueReflect(mValue) {
			result.Set(mKeyString, nil)
			continue
		}
		mValueProcessed := getFieldValue(mValue.Elem().Type(), mValue.Elem(), tag)
		if len(mKeyString) != 0 && mValueProcessed != nil {
			result.Set(mKeyString, mValueProcessed)
		}
	}
	if sub {
		return result.Values()
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return v.Interface()
	}
	return string(bytes)
}

func prepareSliceMsg(v reflect.Value, sub bool, tag string) any {
	var result []any
	for i := 0; i < v.Len(); i++ {
		fieldValue := v.Index(i)
		if util.IsNilValueReflect(fieldValue) {
			result = append(result, nil)
			continue
		} else if fieldValue.Kind() == reflect.Pointer {
			fieldValue = fieldValue.Elem()
		}
		fieldValueProcessed := getFieldValue(fieldValue.Type(), fieldValue, tag)
		if fieldValueProcessed != nil {
			result = append(result, fieldValueProcessed)
		}
	}
	if sub {
		return result
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return v.Interface()
	}
	return string(bytes)
}

func getArgLogLevel(lvl level) string {
	var color string
	if !opts.DisablePrefixColors {
		color = lvl.Color()
	}
	return strings.Join([]string{color, lvl.String(), "\u001B[0m"}, "")
}

func getArgDatetime() string {
	return getCurrentTime(opts.UTC).Format(opts.DateFormat.Format())
}

func getArgCaller(skipCaller int) string {
	fileName, line, _ := util.GetCallerInfo(skipCaller)
	return fileName + ":" + line
}

func getFieldValue(fieldType reflect.Type, fieldValue reflect.Value, tag string) any {
	valueByReflect := getValueByReflect(fieldType, fieldValue, tag)
	if valueByReflect != nil {
		return valueByReflect
	}
	return nil
}

func getValueByReflect(t reflect.Type, v reflect.Value, tag string) any {
	if !v.CanInterface() {
		return nil
	} else if _, ok := v.Interface().(time.Time); ok {
		return convertStringReflectValue(v, tag)
	}
	switch t.Kind() {
	case reflect.Struct:
		return prepareStructMsg(t, v, true, tag)
	case reflect.Map:
		return prepareMapMsg(v, true, tag)
	case reflect.Array, reflect.Slice:
		return prepareSliceMsg(v, true, tag)
	default:
		if strings.Contains(tag, "hide") || strings.Contains(tag, "mask_start") || strings.Contains(tag, "mask_end") {
			return convertStringReflectValue(v, tag)
		}
		return v.Interface()
	}
}

func convertStringReflectValue(v reflect.Value, tag string) string {
	s := util.ConvertToString(v.Interface())
	if len(s) == 0 {
		return s
	}
	if strings.Contains(tag, "hide") {
		s = maskString(s, '*')
	} else if strings.Contains(tag, "mask_start") {
		s = maskStartOrEndOfString(s, '*', true)
	} else if strings.Contains(tag, "mask_end") {
		s = maskStartOrEndOfString(s, '*', false)
	}
	return s
}

func maskString(s string, mask rune) string {
	maskedString := make([]rune, len(s))
	for i := range maskedString {
		maskedString[i] = mask
	}
	return string(maskedString)
}

func maskStartOrEndOfString(s string, mask rune, maskStart bool) string {
	if len(s) == 1 {
		return string(mask)
	}
	maskedString := make([]rune, len(s))
	copy(maskedString, []rune(s))
	maskIndex := len(s) / 2
	if maskStart {
		for i := 0; i < maskIndex; i++ {
			maskedString[i] = mask
		}
	} else {
		for i := maskIndex; i < len(s); i++ {
			maskedString[i] = mask
		}
	}
	return string(maskedString)
}

func getCurrentTime(useUTC bool) time.Time {
	t := time.Now()
	if useUTC {
		return t.UTC()
	}
	return t.Local()
}
