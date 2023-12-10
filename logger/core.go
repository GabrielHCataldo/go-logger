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

var opts = &Options{}

type Options struct {
	Mode                   Mode
	DateFormat             DateFormat
	EnableAsynchronousMode bool
	UTC                    bool
	DontPrintEmptyMessage  bool
	RemoveSpace            bool
	HideAllArgs            bool
	HideArgDatetime        bool
	HideArgCaller          bool
	DisablePrefixColors    bool
}

type logJson struct {
	Level    string `json:"level,omitempty"`
	Datetime string `json:"datetime,omitempty"`
	File     string `json:"file,omitempty"`
	Func     string `json:"func,omitempty"`
	Line     string `json:"line,omitempty"`
	Msg      string `json:"msg"`
}

func SetOptions(options *Options) {
	opts = options
}

func ResetOptionsToDefault() {
	opts = &Options{}
}

func printLog(lvl level, skipCaller int, opts Options, format string, tag string, v ...any) {
	if opts.EnableAsynchronousMode {
		go executePrintLog(lvl, skipCaller, opts, format, tag, v...)
	} else {
		executePrintLog(lvl, skipCaller, opts, format, tag, v...)
	}
}

func executePrintLog(lvl level, skipCaller int, opts Options, format string, tag string, v ...any) {
	logger := getLogger(lvl, skipCaller+1, opts)
	msg := prepareMsg(lvl, skipCaller, tag, opts, v...)
	if opts.DontPrintEmptyMessage && (msg == nil && len(msg) == 0) {
		return
	}
	if len(format) != 0 {
		logger.Printf(format, msg...)
	} else if opts.RemoveSpace {
		logger.Print(msg...)
	} else {
		logger.Println(msg...)
	}
}

func getLogger(lvl level, skipCaller int, opts Options) *log.Logger {
	if opts.Mode == ModeJson {
		return log.New(os.Stdout, "", 0)
	}
	return log.New(os.Stdout, getLoggerNormalPrefix(lvl, skipCaller, opts), 0)
}

func prepareMsg(lvl level, skipCaller int, tag string, opts Options, msgContents ...any) []any {
	var processedMsg []any
	for _, msgContent := range msgContents {
		if msgContent == nil {
			continue
		}
		valueType := reflect.TypeOf(msgContent)
		value := reflect.ValueOf(msgContent)
		if valueType.Kind() == reflect.Pointer || valueType.Kind() == reflect.Interface {
			valueType = valueType.Elem()
			value = value.Elem()
		}
		if opts.DontPrintEmptyMessage && util.IsZeroReflect(value) {
			continue
		}
		processedValue := processMsgValue(valueType, value, tag)
		processedMsg = append(processedMsg, processedValue)
	}
	if opts.Mode == ModeJson {
		return prepareModeJsonMsg(lvl, skipCaller+1, opts, processedMsg...)
	}
	return processedMsg
}

func processMsgValue(valueType reflect.Type, value reflect.Value, tag string) any {
	var processedValue any
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

func getLogJson(lvl level, skipCaller int, opts Options, v ...any) logJson {
	lg := logJson{
		Level: lvl.String(),
		Msg:   strings.Replace(fmt.Sprintln(v...), "\n", "", -1),
	}
	if opts.HideAllArgs {
		return lg
	}
	if !opts.HideArgDatetime {
		lg.Datetime = getArgDatetime(opts)
	}
	if !opts.HideArgCaller {
		fileName, line, funcName := util.GetCallerInfo(skipCaller + 1)
		lg.File = fileName
		lg.Func = funcName
		lg.Line = line
	}
	return lg
}

func prepareModeJsonMsg(lvl level, skipCaller int, opts Options, v ...any) []any {
	if opts.DontPrintEmptyMessage && len(v) == 0 {
		return []any{}
	}
	var processedMsg any
	processedMsg = getLogJson(lvl, skipCaller, opts, v...)
	bytes, err := json.Marshal(processedMsg)
	if err == nil {
		processedMsg = strings.ReplaceAll(string(bytes), "\\", "")
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
	bytes, _ := json.Marshal(result)
	return string(bytes)
}

func prepareMapMsg(v reflect.Value, sub bool, tag string) any {
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for _, key := range v.MapKeys() {
		mKey := key.Convert(v.Type().Key())
		mValue := v.MapIndex(mKey)
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
	bytes, _ := json.Marshal(result)
	return string(bytes)
}

func prepareSliceMsg(v reflect.Value, sub bool, tag string) any {
	var result []any
	for i := 0; i < v.Len(); i++ {
		fieldValue := v.Index(i)
		if util.IsNilValueReflect(fieldValue) {
			result = append(result, nil)
			continue
		} else if fieldValue.Kind() == reflect.Pointer || fieldValue.Kind() == reflect.Interface {
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
	bytes, _ := json.Marshal(result)
	return string(bytes)
}

func buildDatetimeString(opts Options) string {
	if opts.HideAllArgs || opts.HideArgDatetime {
		return ""
	}
	return " " + getArgDatetime(opts) + "]"
}

func getLoggerNormalPrefix(lvl level, skipCaller int, opts Options) string {
	var b strings.Builder
	if !opts.HideAllArgs && !opts.HideArgDatetime {
		b.WriteString("[")
	}
	datetimeString := buildDatetimeString(opts)
	b.WriteString(getArgLogLevel(lvl, opts))
	b.WriteString(datetimeString)
	if !opts.HideAllArgs && !opts.HideArgCaller {
		b.WriteString(" " + getArgCaller(skipCaller+1) + ":")
	} else if datetimeString == "" {
		b.WriteString(":")
	}
	b.WriteString(" ")
	return b.String()
}

func getArgLogLevel(lvl level, opts Options) string {
	var color string
	if opts.DisablePrefixColors {
		color = level("").Color()
	} else {
		color = lvl.Color()
	}
	return strings.Join([]string{color, lvl.String(), "\u001B[0m"}, "")
}

func getArgDatetime(opts Options) string {
	return getCurrentTime(opts.UTC).Format(opts.DateFormat.Format())
}

func getArgCaller(skipCaller int) string {
	fileName, line, _ := util.GetCallerInfo(skipCaller)
	return fileName + ":" + line
}

func getFieldValue(fieldType reflect.Type, fieldValue reflect.Value, tag string) any {
	return getValueByReflect(fieldType, fieldValue, tag)
}

func getValueByReflect(t reflect.Type, v reflect.Value, tag string) any {
	if _, ok := v.Interface().(time.Time); ok {
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
		if strings.Contains(tag, loggerTagHide) ||
			strings.Contains(tag, loggerTagMaskStart) ||
			strings.Contains(tag, loggerTagMaskEnd) {
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
	if strings.Contains(tag, loggerTagHide) {
		s = maskString(s, '*')
	} else if strings.Contains(tag, loggerTagMaskStart) {
		s = maskStartOrEndOfString(s, '*', true)
	} else if strings.Contains(tag, loggerTagMaskEnd) {
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
