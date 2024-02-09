package logger

import (
	"encoding/json"
	"fmt"
	"github.com/GabrielHCataldo/go-helper/helper"
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
	// Custom prefix text
	CustomPrefixText string
	// Custom after prefix text
	CustomAfterPrefixText string
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
	// If true, it will disable all message colors (default: false)
	DisableMessageColors bool
	// If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
	//
	// IMPORTANT: If true, the format parameter will not work
	EnableJsonMsgFieldForSlice bool
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
//	- EnableJsonMsgFieldForSlice: If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
//	  IMPORTANT: If true, the format parameter will not work
//
// # Return:
//
//	None
func SetOptions(options *Options) {
	if helper.IsNotNil(options) {
		opts = options
	} else {
		opts = &Options{}
	}
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
	if opts.DontPrintEmptyMessage && helper.IsEmpty(msg) {
		return
	}
	if helper.Equals(opts.Mode, ModeJson) {
		printJsonMsg(logger, lvl, skipCaller+1, opts, format, msg...)
	} else {
		printDefaultMsg(logger, lvl, format, msg...)
	}
}

func getLogger(lvl level, skipCaller int, opts Options) *log.Logger {
	if helper.Equals(opts.Mode, ModeJson) {
		return log.New(os.Stdout, "", 0)
	}
	return log.New(os.Stdout, getLoggerNormalPrefix(lvl, skipCaller+1, opts), 0)
}

func prepareMsg(tag string, opts Options, msgContents ...any) []any {
	var processedMsg []any
	for _, msgContent := range msgContents {
		if helper.IsNil(msgContent) {
			continue
		} else if opts.DontPrintEmptyMessage && helper.IsEmpty(msgContent) {
			continue
		}
		processedValue := processMsgValue(msgContent, tag, opts)
		processedMsg = append(processedMsg, processedValue)
	}
	return processedMsg
}

func printDefaultMsg(logger *log.Logger, lvl level, format string, msg ...any) {
	nMsg := prepareMessageColor(lvl, msg...)
	if helper.IsNotEmpty(format) {
		logger.Printf(format, nMsg...)
	} else if opts.RemoveSpace {
		logger.Print(nMsg...)
	} else {
		logger.Println(nMsg...)
	}
}

func printJsonMsg(logger *log.Logger, lvl level, skipCaller int, opts Options, format string, v ...any) {
	var processedMsg any
	processedMsg = getLoggerJson(lvl, skipCaller, opts, format, v...)
	bytes, _ := json.Marshal(processedMsg)
	logger.Print(string(bytes))
}

func processMsgValue(value any, tag string, opts Options) any {
	var processedValue any
	isSub := helper.Equals(opts.Mode, ModeJson) && opts.EnableJsonMsgFieldForSlice
	if helper.IsStruct(value) {
		processedValue = prepareStructMsg(value, isSub, tag)
	} else if helper.IsMap(value) {
		processedValue = prepareMapMsg(value, isSub, tag)
	} else if helper.IsSlice(value) {
		processedValue = prepareSliceMsg(value, isSub, tag)
	} else {
		processedValue = prepareValue(value, tag)
	}
	return processedValue
}

func prepareStructMsg(value any, sub bool, tag string) any {
	v, t := reflectValueOf(value)
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldStruct := t.Field(i)
		var fieldRealValue any
		var fieldName string
		var fieldTag string
		fieldName = getJsonNameByTag(fieldStruct.Tag.Get("json"))
		fieldTag = tag
		if !fieldStruct.IsExported() || !fieldValue.IsValid() || !fieldValue.CanInterface() {
			continue
		}
		if helper.IsNotNil(fieldValue.Interface()) {
			fieldRealValue = fieldValue.Interface()
			if (helper.IsPointer(fieldRealValue) || helper.IsInterface(fieldRealValue)) &&
				helper.IsNotNil(fieldValue.Elem().Interface()) {
				fieldRealValue = fieldValue.Elem().Interface()
			}
		}
		if helper.IsEmpty(fieldName) {
			fieldName = fieldStruct.Name
		}
		if helper.IsEmpty(fieldTag) {
			fieldTag = fieldStruct.Tag.Get("logger")
		}
		if strings.Contains(fieldStruct.Tag.Get("json"), "-") || (helper.IsEmpty(fieldRealValue) &&
			strings.Contains(fieldStruct.Tag.Get("json"), "omitempty")) {
			continue
		}
		fieldValueProcessed := prepareValue(fieldRealValue, fieldTag)
		if helper.IsNotEmpty(fieldName) {
			result.Set(fieldName, fieldValueProcessed)
		}
	}
	if sub {
		return result
	}
	return helper.SimpleConvertToString(result)
}

func prepareMapMsg(value any, sub bool, tag string) any {
	v, _ := reflectValueOf(value)
	result := orderedmap.New()
	result.SetEscapeHTML(false)
	for _, key := range v.MapKeys() {
		mKey := key.Convert(v.Type().Key())
		mValue := v.MapIndex(mKey)
		mKeyString := helper.SimpleConvertToString(mKey.Interface())
		var mRealValue any
		if mValue.CanInterface() && helper.IsNotNil(mValue.Interface()) {
			mRealValue = mValue.Interface()
			if (helper.IsPointer(mRealValue) || helper.IsInterface(mRealValue)) &&
				helper.IsNotNil(mValue.Elem().Interface()) {
				mRealValue = mValue.Elem().Interface()
			}
		}
		mValueProcessed := prepareValue(mRealValue, tag)
		if helper.IsNotEmpty(mKeyString) {
			result.Set(mKeyString, mValueProcessed)
		}
	}
	if sub {
		return result
	}
	return helper.SimpleConvertToString(result)
}

func prepareSliceMsg(value any, sub bool, tag string) any {
	v, _ := reflectValueOf(value)
	var result []any
	for i := 0; i < v.Len(); i++ {
		var indexRealValue any
		indexValue := v.Index(i)
		if indexValue.CanInterface() && helper.IsNotNil(indexValue.Interface()) {
			indexRealValue = indexValue.Interface()
			if (helper.IsPointer(indexRealValue) || helper.IsInterface(indexRealValue)) &&
				helper.IsNotNil(indexValue.Elem().Interface()) {
				indexRealValue = indexValue.Elem().Interface()
			}
		}
		indexValueProcessed := prepareValue(indexRealValue, tag)
		result = append(result, indexValueProcessed)
	}
	if sub {
		return result
	}
	return helper.SimpleConvertToString(result)
}

func buildDatetimeString(opts Options) string {
	if opts.HideAllArgs || opts.HideArgDatetime {
		return ""
	}
	return " " + getArgDatetime(opts) + "]"
}

func getLoggerNormalPrefix(lvl level, skipCaller int, opts Options) string {
	var b strings.Builder
	if helper.IsNotEmpty(opts.CustomPrefixText) {
		b.WriteString(opts.CustomPrefixText)
		b.WriteString(" ")
	}
	if !opts.HideAllArgs && !opts.HideArgDatetime {
		b.WriteString("[")
	}
	datetimeString := buildDatetimeString(opts)
	b.WriteString(getArgLogLevel(lvl, opts))
	b.WriteString(datetimeString)
	if !opts.HideAllArgs && !opts.HideArgCaller {
		b.WriteString(" ")
		b.WriteString(StyleUnderscore)
		b.WriteString(getArgCaller(skipCaller))
		b.WriteString(StyleReset)
		b.WriteString(":")
	} else if helper.IsEmpty(datetimeString) {
		b.WriteString(":")
	}
	b.WriteString(" ")
	if helper.IsNotEmpty(opts.CustomAfterPrefixText) {
		b.WriteString(opts.CustomAfterPrefixText)
		b.WriteString(" ")
	}
	prepareMessageColorOnPrefix(lvl, &b)
	return b.String()
}

func getLoggerJson(lvl level, skipCaller int, opts Options, format string, v ...any) logJson {
	var msg any
	var prefixes []any
	if helper.IsNotEmpty(opts.CustomPrefixText) {
		prefixes = append(prefixes, opts.CustomPrefixText)
	}
	if opts.EnableJsonMsgFieldForSlice {
		msg = v
	} else if helper.IsNotEmpty(format) && helper.IsNotEmpty(prefixes) {
		prefix := fmt.Sprintln(prefixes)
		msg = fmt.Sprintln(prefix, fmt.Sprintf(format, v...))
	} else if helper.IsNotEmpty(format) {
		msg = fmt.Sprintf(format, v...)
	} else {
		msg = helper.Sprintln(v...)
	}
	lg := logJson{
		Level: lvl.string(),
		Msg:   msg,
	}
	if opts.HideAllArgs {
		return lg
	}
	if !opts.HideArgDatetime {
		lg.Datetime = getArgDatetime(opts)
	}
	if !opts.HideArgCaller {
		fileName, line, funcName := helper.GetCallerInfo(skipCaller)
		lg.File = fileName
		lg.Func = funcName
		lg.Line = line
	}
	return lg
}

func getArgLogLevel(lvl level, opts Options) string {
	color := StyleBold
	if opts.DisablePrefixColors {
		color += level("").colorLevel()
	} else {
		color += lvl.colorLevel()
	}
	return strings.Join([]string{color, lvl.string(), StyleReset}, "")
}

func getArgDatetime(opts Options) string {
	return getCurrentTime(opts.UTC).Format(opts.DateFormat.Format())
}

func getArgCaller(skipCaller int) string {
	skipCaller = int(math.Max(float64(skipCaller), 0))
	fileName, line, _ := helper.GetCallerInfo(skipCaller)
	return fileName + ":" + line
}

func prepareMessageColorOnPrefix(lvl level, b *strings.Builder) {
	if !opts.DisableMessageColors {
		b.WriteString(lvl.colorMessage())
	}
}

func prepareMessageColor(lvl level, msg ...any) []any {
	if !opts.DisableMessageColors {
		var nMsg []any
		for _, vMsg := range msg {
			sMsg := helper.SimpleConvertToString(vMsg)
			sMsg = strings.ReplaceAll(sMsg, "\n", fmt.Sprint("\n", lvl.colorMessage()))
			nMsg = append(nMsg, sMsg)
		}
		return nMsg
	}
	return msg
}

func prepareValue(value any, tag string) any {
	if helper.IsStruct(value) {
		return prepareStructMsg(value, true, tag)
	} else if helper.IsMap(value) {
		return prepareMapMsg(value, true, tag)
	} else if helper.IsSlice(value) {
		return prepareSliceMsg(value, true, tag)
	} else if helper.IsNotNil(value) {
		return convertValueToString(value, tag)
	} else {
		return nil
	}
}

func convertValueToString(value any, tag string) string {
	s := helper.SimpleConvertToString(value)
	if helper.IsEmpty(s) {
		return s
	} else if strings.Contains(tag, loggerTagHide) {
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

func reflectValueOf(value any) (reflect.Value, reflect.Type) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if helper.IsPointer(value) || helper.IsInterface(value) {
		v = v.Elem()
		t = t.Elem()
	}
	return v, t
}

func getJsonNameByTag(tag string) string {
	result := ""
	splitTag := strings.Split(tag, ",")
	if helper.IsNotEmpty(splitTag) {
		result = splitTag[0]
	}
	if helper.IsEmpty(result) || helper.Equals(result, "omitempty") {
		return ""
	}
	return result
}
