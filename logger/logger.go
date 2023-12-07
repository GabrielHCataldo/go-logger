package logger

import (
	"encoding/json"
	"fmt"
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
	//todo -> aqui precisamos setar o hide e mask tags colocadas no objeto para dar o replace no valor
	// criar tbm InfoMask e assim por diante para variaveis
	for _, v := range vs {
		var vr any
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
		switch t.Kind() {
		case reflect.Struct, reflect.Slice, reflect.Map, reflect.Array:
			b, err := json.Marshal(vr)
			if err == nil {
				vr = string(b)
			} else {
				vr = v
			}
			break
		default:
			vr = v
		}
		msg = append(msg, vr)
	}
	switch opt.Mode {
	case ModeJson:
		return prepareMsgJson(l, skipCaller, msg)
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
		vr = string(b)
	} else {
		vr = jLog
	}
	return []any{vr}
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
