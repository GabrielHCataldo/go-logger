package logger

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
	printLog(levelInfo, 3, *opts, "", "", v...)
}

func InfoF(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, "", v...)
}

func InfoOpts(opts options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagHide, v...)
}

// InfoSkipCaller is a function that logs informational messages based on the global options variable, using the INFO level,
// unlike Info, this function passes skipCaller as a parameter where the value passed directly
// impacts the file and line information of the message file (v ...any).
//
// Usage Example:
//
// logger.InfoSkipCaller(2, "test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func InfoSkipCaller(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", "", v...)
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
	printLog(levelInfo, 3, *opts, "", loggerTagHide, v...)
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
	printLog(levelInfo, 3, *opts, "", loggerTagMaskStart, v...)
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
	printLog(levelInfo, 3, *opts, "", loggerTagMaskEnd, v...)
}

func InfoFH(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagHide, v...)
}

func InfoFMS(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagMaskStart, v...)
}

func InfoFME(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagMaskEnd, v...)
}

// InfoOptsH is a function that logs informational messages based on the global options variable, using the INFO level,
// unlike Info, this function passes skipCaller as a parameter where the value passed directly
// impacts the file and line information of the message file (v ...any).
//
// Usage Example:
//
// logger.InfoSkipCaller(2, "test", true, 112, 10.99)
//
// Result (options default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func InfoOptsH(opts options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagHide, v...)
}

func InfoOptsMS(opts options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagMaskStart, v...)
}

func InfoOptsME(opts options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagMaskEnd, v...)
}

func InfoOptsF(format string, opts options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagHide, v...)
}

func InfoOptsFH(format string, opts options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagHide, v...)
}

func InfoOptsFMS(format string, opts options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagMaskStart, v...)
}

func InfoOptsFME(format string, opts options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagMaskEnd, v...)
}

func InfoSkipCallerF(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagHide, v...)
}

func InfoSkipCallerH(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagHide, v...)
}

func InfoSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

func InfoSkipCallerME(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

func InfoSkipCallerFH(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagHide, v...)
}

func InfoSkipCallerFMS(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func InfoSkipCallerFME(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func InfoSkipCallerOpts(skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", "", v...)
}

func InfoSkipCallerOptsF(format string, skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, "", v...)
}

func InfoSkipCallerOptsH(skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagHide, v...)
}

func InfoSkipCallerOptsMS(skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagMaskStart, v...)
}

func InfoSkipCallerOptsME(skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

func InfoSkipCallerOptsFH(format string, skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagHide, v...)
}

func InfoSkipCallerOptsFMS(format string, skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func InfoSkipCallerOptsFME(format string, skipCaller int, opts options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
