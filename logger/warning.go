package logger

// Warning is a function that records informational messages based on the global Options variable, using the level WARNING.
//
// # Usage Example:
//
// logger.Warning("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Warning(v ...any) {
	printLog(levelWarning, 2, *opts, "", "", v...)
}

// Warningf is a function that records informational messages based on the Options global variable, using the WARNING level.
// Unlike Warning, this function passes the format you want as a parameter, so the value passed will impact the format
// of the printed message.
//
// # Usage Example:
//
// logger.Warningf("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func Warningf(format string, v ...any) {
	printLog(levelWarning, 2, *opts, format, "", v...)
}

// WarningOpts is a function that records informational messages, using the WARNING level. Unlike Warning, this function passes
// Options as a parameter, and it will only be persisted in that message, therefore, this Options passed as a parameter
// will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.WarningOpts(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func WarningOpts(opts Options, v ...any) {
	printLog(levelWarning, 2, opts, "", "", v...)
}

// WarningSkipCaller is a function that logs informational messages based on the Options global variable, using the WARNING
// level. Unlike Warning, this function passes skipCaller as a parameter, therefore the value passed directly impacts the
// file name and line when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCaller(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func WarningSkipCaller(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, "", "", v...)
}

// WarningH is a function that records informational messages based on the Options global variable, using the WARNING level.
// Unlike Warning, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningH("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func WarningH(v ...any) {
	printLog(levelWarning, 2, *opts, "", loggerTagHide, v...)
}

// WarningMS is a function that records informational messages based on the Options global variable, using the WARNING level.
// Unlike Warning, this function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningMS("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func WarningMS(v ...any) {
	printLog(levelWarning, 2, *opts, "", loggerTagMaskStart, v...)
}

// WarningME is a function that records informational messages based on the Options global variable, using the WARNING level.
// Unlike Warning, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningME("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarningME(v ...any) {
	printLog(levelWarning, 2, *opts, "", loggerTagMaskEnd, v...)
}

// WarningfH is a function that records informational messages based on the Options global variable, using the WARNING level,
// printing the parameters v based on the format entered as a parameter. Unlike Warningf, this function will hide all
// values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningfH("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func WarningfH(format string, v ...any) {
	printLog(levelWarning, 2, *opts, format, loggerTagHide, v...)
}

// WarningfMS is a function that records informational messages based on the Options global variable, using the WARNING level,
// printing the parameters v based on the format entered as a parameter. Unlike Warningf, this function will mask the
// first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningfMS("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func WarningfMS(format string, v ...any) {
	printLog(levelWarning, 2, *opts, format, loggerTagMaskStart, v...)
}

// WarningfME is a function that records informational messages based on the Options global variable, using the WARNING level,
// printing the parameters v based on the format entered as a parameter. Unlike Warningf, this function will mask the last
// half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningfME("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func WarningfME(format string, v ...any) {
	printLog(levelWarning, 2, *opts, format, loggerTagMaskEnd, v...)
}

// WarningOptsH is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARNING level. Unlike WarningOpts, this function will hide all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarningOptsH(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: ****, ****, ***, *****
func WarningOptsH(opts Options, v ...any) {
	printLog(levelWarning, 2, opts, "", loggerTagHide, v...)
}

// WarningOptsMS is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARNING level. Unlike WarningOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarningOptsMS(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func WarningOptsMS(opts Options, v ...any) {
	printLog(levelWarning, 2, opts, "", loggerTagMaskStart, v...)
}

// WarningOptsME is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARNING level. Unlike WarningOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarningOptsME(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarningOptsME(opts Options, v ...any) {
	printLog(levelWarning, 2, opts, "", loggerTagMaskEnd, v...)
}

// WarningOptsf is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARNING level. Unlike WarningOpts, this function passes the format you want as a parameter, so the value passed
// will impact the format of the printed message.
//
// # Usage Example:
//
// logger.WarningOptsf("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func WarningOptsf(format string, opts Options, v ...any) {
	printLog(levelWarning, 2, opts, format, "", v...)
}

// WarningOptsfH is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARNING level, print the parameters v also based on the format entered as a parameter. Unlike WarningOptsf, this
// function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningOptsfH("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func WarningOptsfH(format string, opts Options, v ...any) {
	printLog(levelWarning, 2, opts, format, loggerTagHide, v...)
}

// WarningOptsfMS is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARNING level, print the parameters v also based on the format entered as a parameter. Unlike WarningOptsf, this
// function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningOptsfMS("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func WarningOptsfMS(format string, opts Options, v ...any) {
	printLog(levelWarning, 2, opts, format, loggerTagMaskStart, v...)
}

// WarningOptsfME is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARNING level, print the parameters v also based on the format entered as a parameter. Unlike WarningOptsf, this
// function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningOptsfME("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func WarningOptsfME(format string, opts Options, v ...any) {
	printLog(levelWarning, 2, opts, format, loggerTagMaskEnd, v...)
}

// WarningSkipCallerf is a function that logs informational messages based on the Options global variable, using the
// WARNING level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarningSkipCaller, this function passes the format you want as a parameter, so the value passed will have an
// impact on the format of the printed message.
//
// # Usage Example:
//
// logger.WarningSkipCallerf("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarningSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, format, "", v...)
}

// WarningSkipCallerH is a function that logs informational messages based on the Options global variable, using the
// WARNING level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarningSkipCaller, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerH(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: ****, ****, ***, *****
func WarningSkipCallerH(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, "", loggerTagHide, v...)
}

// WarningSkipCallerMS is a function that logs informational messages based on the Options global variable, using the
// WARNING level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarningSkipCaller, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerMS(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func WarningSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, "", loggerTagMaskStart, v...)
}

// WarningSkipCallerME is a function that logs informational messages based on the Options global variable, using the
// WARNING level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarningSkipCaller, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerME(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarningSkipCallerME(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, "", loggerTagMaskEnd, v...)
}

// WarningSkipCallerfH is a function that writes informational messages based on the Options global variable, using the
// WARNING level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarningSkipCallerf, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerfH("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func WarningSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, format, loggerTagHide, v...)
}

// WarningSkipCallerfMS is a function that writes informational messages based on the Options global variable, using the
// WARNING level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarningSkipCallerf, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerfMS("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func WarningSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, format, loggerTagMaskStart, v...)
}

// WarningSkipCallerfME is a function that writes informational messages based on the Options global variable, using the
// WARNING level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarningSkipCallerf, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerfME("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func WarningSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller+1, *opts, format, loggerTagMaskEnd, v...)
}

// WarningSkipCallerOpts is a function that records informational messages, using the WARNING level, the file name and line
// information in the message will be based on the value of the skipCaller parameter. Unlike WarningSkipCaller, this
// function passes Options as a parameter, and it will only be persisted in that message, therefore,
// this Options passed as a parameter will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.WarningSkipCallerOpts(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarningSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, "", "", v...)
}

// WarningSkipCallerOptsf is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARNING level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarningSkipCallerOpts, this function passes the format you want as a parameter, so the
// value passed will impact the format of the printed message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsf("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarningSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, format, "", v...)
}

// WarningSkipCallerOptsH is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARNING level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarningSkipCallerOpts, this function will hide all values passed as parameters v when
// printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsH(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:23: **** **** *** *****
func WarningSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, "", loggerTagHide, v...)
}

// WarningSkipCallerOptsMS is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARNING level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarningSkipCallerOpts, this function will mask the first half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsMS(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func WarningSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, "", loggerTagMaskStart, v...)
}

// WarningSkipCallerOptsME is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARNING level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarningSkipCallerOpts, this function will mask the last half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsME(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: te** tr** 1** 10***
func WarningSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, "", loggerTagMaskEnd, v...)
}

// WarningSkipCallerOptsfH is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARNING level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarningSkipCallerOptsf, this function will mask the last half of all values passed as v settings when
// printing a message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsfH("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func WarningSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, format, loggerTagHide, v...)
}

// WarningSkipCallerOptsfMS is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARNING level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarningSkipCallerOptsf, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsfMS("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func WarningSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, format, loggerTagMaskStart, v...)
}

// WarningSkipCallerOptsfME is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARNING level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarningSkipCallerOptsf, this function will mask the last half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarningSkipCallerOptsfME("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARNING 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func WarningSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller+1, opts, format, loggerTagMaskEnd, v...)
}
