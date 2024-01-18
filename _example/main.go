package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"os"
	"os/signal"
	"time"
)

type test struct {
	Name          string    `json:"name,omitempty"`
	BirthDate     time.Time `json:"birthDate,omitempty"`
	Document      string    `json:"document,omitempty" logger:"mask_start"`
	Emails        []string  `json:"emails,omitempty"`
	Balances      []float32 `json:"balances,omitempty" logger:"mask_end"`
	TotalBalances []float64 `json:"totalBalances,omitempty"`
	Booleans      []bool    `json:"booleans,omitempty"`
	Bank          bankTest  `json:"bank,omitempty"`
}

type bankTest struct {
	AccountDigits string  `json:"accountDigits,omitempty"`
	Account       string  `json:"account,omitempty"`
	Balance       float32 `json:"balance,omitempty" logger:"hide"`
	TotalBalance  float64 `json:"totalBalance,omitempty"`
}

func main() {
	basic()
	globalJsonMode()
	partialJsonMode()
	maskHideStruct()
	maskHideFunc()
	formatFunc()
	skipCallerFun()
	customerOptions()
	dontPrintEmptyMessage()
	async()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		logger.Info("Stop application!")
	}
}

func basic() {
	basicMsg := getBasicMsg()
	logger.Info(basicMsg...)
	logger.Debug(basicMsg...)
	logger.Warning(basicMsg...)
	logger.Error(basicMsg...)
}

func globalJsonMode() {
	// set on global options
	logger.SetOptions(&logger.Options{
		// Print mode (default: ModeDefault)
		Mode: logger.ModeJson,
	})
	basicMsg := getBasicMsg()
	logger.Info(basicMsg...)
	logger.Debug(basicMsg...)
	logger.Warning(basicMsg...)
	logger.Error(basicMsg...)
}

func partialJsonMode() {
	opts := logger.Options{Mode: logger.ModeJson}
	basicMsg := getBasicMsg()
	logger.InfoOpts(opts, basicMsg...)
	// field msg json to slice
	opts.EnableJsonMsgFieldForSlice = true
	logger.DebugOpts(opts, basicMsg...)
	logger.Warning(basicMsg...)
	logger.Error(basicMsg...)
}

func maskHideStruct() {
	testStruct := getTestStruct()
	msg := []any{"test mask/hide struct:", testStruct}
	logger.Info(msg...)
	logger.Debug(msg...)
	logger.Warning(msg...)
	logger.Error(msg...)
}

func maskHideFunc() {
	msg := getBasicMsg()
	logger.Info(msg...)
	logger.InfoH(msg...)
	logger.InfoMS(msg...)
	logger.InfoME(msg...)
}

func formatFunc() {
	customOpts := logger.Options{
		// Print mode (default: ModeDefault)
		Mode: logger.ModeJson,
		// Argument date format (default: DateFormatFull24h)
		DateFormat: logger.DateFormatFull24h,
		// Custom prefix text
		CustomPrefixText: "",
		// Custom after prefix text (only if Mode is ModeDefault)
		CustomAfterPrefixText: "",
		// Enable asynchronous printing mode (default: false)
		EnableAsynchronousMode: false,
		// Enable argument date to be UTC (default: false)
		UTC: false,
		// Enable to not print empty message (default: false)
		DontPrintEmptyMessage: false,
		// Enable to remove spaces between parameters (default: false)
		RemoveSpace: false,
		// If true will hide all datetime and prefix arguments (default: false)
		HideAllArgs: false,
		// If true it will hide the datetime arguments (default: false)
		HideArgDatetime: false,
		// If true, it will hide the caller arguments (default: false)
		HideArgCaller: false,
		// If true, it will disable all argument and prefix colors (default: false)
		DisablePrefixColors: false,
		// If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
		//
		// IMPORTANT: If true, the format parameter will not work
		EnableJsonMsgFieldForSlice: false,
	}
	format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
	msg := getBasicMsg()
	logger.Infof(format, msg...)
	logger.DebugOptsf(format, customOpts, msg...)
	logger.WarningfMS(format, msg...)
	logger.ErrorOptsfMS(format, customOpts, msg...)
}

func skipCallerFun() {
	subFunc()
}

func subFunc() {
	customOpts := logger.Options{
		Mode: logger.ModeJson,
		// Argument date format (default: DateFormatFull24h)
		DateFormat: logger.DateFormatFull12h,
		// Custom prefix text
		CustomPrefixText: "",
		// Custom after prefix text (only if Mode is ModeDefault)
		CustomAfterPrefixText: "",
		// Enable asynchronous printing mode (default: false)
		EnableAsynchronousMode: false,
		// Enable argument date to be UTC (default: false)
		UTC: false,
		// Enable to not print empty message (default: false)
		DontPrintEmptyMessage: false,
		// Enable to remove spaces between parameters (default: false)
		RemoveSpace: false,
		// If true will hide all datetime and prefix arguments (default: false)
		HideAllArgs: false,
		// If true it will hide the datetime arguments (default: false)
		HideArgDatetime: false,
		// If true, it will hide the caller arguments (default: false)
		HideArgCaller: false,
		// If true, it will disable all argument and prefix colors (default: false)
		DisablePrefixColors: false,
		// If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
		//
		// IMPORTANT: If true, the format parameter will not work
		EnableJsonMsgFieldForSlice: false,
	}
	format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
	msg := getBasicMsg()
	logger.InfoSkipCaller(1, msg...)
	logger.DebugSkipCallerOpts(1, customOpts, msg...)
	logger.WarningSkipCallerOptsf(format, 2, customOpts, msg...)
	logger.ErrorSkipCaller(2, msg...)
}

func customerOptions() {
	logger.SetOptions(getCustomOptionsExample())
	msg := getBasicMsg()
	format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
	logger.Info(msg...)
	logger.ResetOptionsToDefault()
	logger.Debug(msg...)
	logger.WarningSkipCallerOptsf(format, 1, *getCustomOptionsExample(), msg...)
	logger.ErrorOptsf(format, *getCustomOptionsExample(), msg...)
}

func dontPrintEmptyMessage() {
	opt := logger.Options{DontPrintEmptyMessage: true}
	msg := getEmptyBasicMsg()
	logger.SetOptions(&opt)
	logger.Info(msg...)
	logger.Debug(msg...)
	logger.Warning(msg...)
	logger.Error(msg...)
}

func async() {
	asyncOpt := logger.Options{EnableAsynchronousMode: true}
	msg := getBasicMsg()
	logger.SetOptions(&asyncOpt)
	logger.Info(msg...)
	logger.ResetOptionsToDefault()
	logger.InfoOpts(asyncOpt)
}

func getBasicMsg() []any {
	s := []any{
		"text string",
		1,
		12.213,
		true,
		nil,
		time.Now(),
	}
	m := map[string]any{
		"int":      1,
		"float":    12.213,
		"string":   "text string",
		"bool":     true,
		"nilValue": nil,
		"time":     time.Now(),
	}
	return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}

func getCustomOptionsExample() *logger.Options {
	return &logger.Options{
		Mode:                  logger.RandomMode(),
		DateFormat:            logger.RandomDateFormat(),
		UTC:                   helper.RandomBool(),
		DontPrintEmptyMessage: helper.RandomBool(),
		RemoveSpace:           helper.RandomBool(),
		HideAllArgs:           helper.RandomBool(),
		HideArgDatetime:       helper.RandomBool(),
		HideArgCaller:         helper.RandomBool(),
		DisablePrefixColors:   helper.RandomBool(),
	}
}

func getEmptyBasicMsg() []any {
	return []any{"basic example with empty any values", nil, "", 0, map[string]any{}, []any{}}
}

func getTestStruct() test {
	bank := bankTest{
		AccountDigits: "123",
		Account:       "123981023",
		Balance:       30.89,
		TotalBalance:  200.17,
	}
	return test{
		Name:      "Foo Bar",
		BirthDate: time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
		Document:  "02104996642",
		Emails:    []string{"gabriel@test.com", "gabrielcataldo.231@gmail.com", "biel@test.org"},
		Balances:  []float32{10.88, 11, 13.99, 12391.23, 23321},
		Bank:      bank,
	}
}
