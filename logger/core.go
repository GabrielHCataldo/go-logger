package logger

import (
	"encoding/json"
	"fmt"
	"github.com/GabrielHCataldo/go-logger/internal/util"
	"github.com/iancoleman/orderedmap"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
	"time"
)

type logJson struct {
	Level    string `json:"level,omitempty"`
	Datetime string `json:"datetime,omitempty"`
	File     string `json:"file,omitempty"`
	Func     string `json:"func,omitempty"`
	Line     string `json:"line,omitempty"`
	Msg      any    `json:"msg"`
}

var opts = &Options{}

type Options struct {
	// Print mode (default: ModeDefault)
	Mode Mode
	// Argument date format (default: DateFormatFull24h)
	DateFormat DateFormat
	// Enable asynchronous printing mode (default: false)
	EnableAsynchronousMode bool
	// Enable argument date to be UTC (default: false)
	UTC bool
	// Enable to not print empty message (default: false)
	DontPrintEmptyMessage bool
	// Enable to remove spaces between parameters (default: false)
	RemoveSpace bool
	// If true will hide all datetime and prefix arguments (default: false)
	HideAllArgs bool
	// If true it will hide the datetime arguments (default: false)
	HideArgDatetime bool
	// If true, it will hide the caller arguments (default: false)
	HideArgCaller bool
	// If true, it will disable all argument and prefix colors (default: false)
	DisablePrefixColors bool
	// If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
	//
	// IMPORTANT: If true, the format parameter will not work
	EnableMsgFieldForSlice bool
}

// SetOptions sets the global options to the specified options pointer.
//
// This function allows you to configure the behavior of the logging package
// by setting various options such as print mode, date format, asynchronous mode, etc.
//
// # Example usage:
//
//	logger.SetOptions(logger.Options{})
//
// # Parameters:
//
//	options - A pointer to an Options struct that contains the desired configuration options.
//
// # Note:
//
//	The Options struct contains the following fields:
//	- Mode: Print mode (default: ModeDefault)
//	- DateFormat: Argument date format (default: DateFormatFull24h)
//	- EnableAsynchronousMode: Enable asynchronous printing mode (default: false)
//	- UTC: Enable argument date to be UTC (default: false)
//	- DontPrintEmptyMessage: Enable to not print empty message (default: false)
//	- RemoveSpace: Enable to remove spaces between parameters (default: false)
//	- HideAllArgs: If true will hide all datetime and prefix arguments (default: false)
//	- HideArgDatetime: If true it will hide the datetime arguments (default: false)
//	- HideArgCaller: If true, it will hide the caller arguments (default: false)
//	- DisablePrefixColors: If true, it will disable all argument and prefix colors (default: false)
//	- EnableMsgFieldForSlice: If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
//	  IMPORTANT: If true, the format parameter will not work
//
// # Return:
//
//	None
func SetOptions(options *Options) {
	opts = options
}

// ResetOptionsToDefault resets the global options to their default values.
// This function initializes the opts variable with a new instance of the Options struct, effectively resetting all
// options to their default values.
//
// # Example usage:
//
//	logger.ResetOptionsToDefault()
//
// # Return:
//
//	None
func ResetOptionsToDefault() {
	opts = &Options{}
}

func printLog(lvl level, skipCaller int, opts Options, format string, tag string, v ...any) {
	if opts.EnableAsynchronousMode {
		opts.HideArgCaller = true
		go executePrintLog(lvl, 1, opts, format, tag, v...)
	} else {
		executePrintLog(lvl, skipCaller+3, opts, format, tag, v...)
	}
}

func executePrintLog(lvl level, skipCaller int, opts Options, format string, tag string, v ...any) {
	logger := getLogger(lvl, skipCaller+1, opts)
	msg := prepareMsg(tag, opts, v...)
	if opts.DontPrintEmptyMessage && (msg == nil && len(msg) == 0) {
		return
	}
	if opts.Mode == ModeJson {
		printJsonMsg(logger, lvl, skipCaller, opts, format, msg...)
	} else {
		printDefaultMsg(logger, format, msg...)
	}
}

func getLogger(lvl level, skipCaller int, opts Options) *log.Logger {
	if opts.Mode == ModeJson {
		return log.New(os.Stdout, "", 0)
	}
	return log.New(os.Stdout, getLoggerNormalPrefix(lvl, skipCaller+1, opts), 0)
}

func prepareMsg(tag string, opts Options, msgContents ...any) []any {
	var processedMsg []any
	for _, msgContent := range msgContents {
		if msgContent == nil {
			continue
		}
		valueType := reflect.TypeOf(msgContent)
		value := reflect.ValueOf(msgContent)
		if valueType.Kind() == reflect.Pointer || valueType.Kind() == reflect.Interface {
			if x, ok := value.Interface().(error); ok {
				s := x.Error()
				valueType = reflect.TypeOf(s)
				value = reflect.ValueOf(s)
			} else {
				valueType = valueType.Elem()
				value = value.Elem()
			}
		}
		if opts.DontPrintEmptyMessage && util.IsZeroReflect(value) {
			continue
		}
		processedValue := processMsgValue(valueType, value, tag, opts)
		processedMsg = append(processedMsg, processedValue)
	}
	return processedMsg
}

func printDefaultMsg(logger *log.Logger, format string, msg ...any) {
	if len(format) != 0 {
		logger.Printf(format, msg...)
	} else if opts.RemoveSpace {
		logger.Print(msg...)
	} else {
		logger.Println(msg...)
	}
}

func printJsonMsg(logger *log.Logger, lvl level, skipCaller int, opts Options, format string, v ...any) {
	var processedMsg any
	processedMsg = getLoggerJson(lvl, skipCaller+1, opts, format, v...)
	bytes, _ := json.Marshal(processedMsg)
	logger.Print(string(bytes))
}

func processMsgValue(valueType reflect.Type, value reflect.Value, tag string, opts Options) any {
	var processedValue any
	isSub := opts.Mode == ModeJson && opts.EnableMsgFieldForSlice
	switch valueType.Kind() {
	case reflect.Struct:
		processedValue = prepareStructMsg(valueType, value, isSub, tag)
	case reflect.Map:
		processedValue = prepareMapMsg(value, isSub, tag)
	case reflect.Slice, reflect.Array:
		processedValue = prepareSliceMsg(value, isSub, tag)
	default:
		processedValue = getValueByReflect(valueType, value, tag)
	}
	return processedValue
}

func prepareStructMsg(t reflect.Type, v reflect.Value, sub bool, tag string) any {
	if _, ok := v.Interface().(time.Time); ok {
		return convertStringReflectValue(v, tag)
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
		fieldValueProcessed := getValueByReflect(fieldType.Type, fieldValue, fieldTag)
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
		mValueProcessed := getValueByReflect(mValue.Elem().Type(), mValue.Elem(), tag)
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
		indexValue := v.Index(i)
		if util.IsNilValueReflect(indexValue) {
			result = append(result, nil)
			continue
		} else if indexValue.Kind() == reflect.Pointer || indexValue.Kind() == reflect.Interface {
			indexValue = indexValue.Elem()
		}
		indexValueProcessed := getValueByReflect(indexValue.Type(), indexValue, tag)
		if indexValueProcessed != nil {
			result = append(result, indexValueProcessed)
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
		b.WriteString(" " + getArgCaller(skipCaller) + ":")
	} else if datetimeString == "" {
		b.WriteString(":")
	}
	b.WriteString(" ")
	return b.String()
}

func getLoggerJson(lvl level, skipCaller int, opts Options, format string, v ...any) logJson {
	var msg any
	if opts.EnableMsgFieldForSlice {
		msg = v
	} else {
		if len(format) != 0 {
			msg = fmt.Sprintf(format, v...)
		} else {
			msg = strings.Replace(fmt.Sprintln(v...), "\n", "", -1)
		}
	}
	lg := logJson{
		Level: lvl.String(),
		Msg:   msg,
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
	skipCaller = int(math.Max(float64(skipCaller), 0))
	fileName, line, _ := util.GetCallerInfo(skipCaller)
	return fileName + ":" + line
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
