package logger

// Debug is a function that records informational messages based on the global Options variable, using the level DEBUG.
//
// # Usage Example:
//
// logger.Debug("test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Debug(v ...any) {
	printLog(levelDebug, 3, *opts, "", "", v...)
}

// Debugf is a function that records informational messages based on the Options global variable, using the DEBUG level.
// Unlike Debug, this function passes the format you want as a parameter, so the value passed will impact the format
// of the printed message.
//
// # Usage Example:
//
// logger.Debugf("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func Debugf(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, "", v...)
}

// DebugOpts is a function that records informational messages, using the DEBUG level. Unlike Debug, this function passes
// Options as a parameter, and it will only be persisted in that message, therefore, this Options passed as a parameter
// will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.DebugOpts(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func DebugOpts(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagHide, v...)
}

// DebugSkipCaller is a function that logs informational messages based on the Options global variable, using the DEBUG
// level. Unlike Debug, this function passes skipCaller as a parameter, therefore the value passed directly impacts the
// file name and line when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCaller(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func DebugSkipCaller(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", "", v...)
}

// DebugH is a function that records informational messages based on the Options global variable, using the DEBUG level.
// Unlike Debug, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugH("test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func DebugH(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagHide, v...)
}

// DebugMS is a function that records informational messages based on the Options global variable, using the DEBUG level.
// Unlike Debug, this function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugMS("test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func DebugMS(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagMaskStart, v...)
}

// DebugME is a function that records informational messages based on the Options global variable, using the DEBUG level.
// Unlike Debug, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugME("test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func DebugME(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagMaskEnd, v...)
}

// DebugfH is a function that records informational messages based on the Options global variable, using the DEBUG level,
// printing the parameters v based on the format entered as a parameter. Unlike Debugf, this function will hide all
// values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugfH("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func DebugfH(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagHide, v...)
}

// DebugfMS is a function that records informational messages based on the Options global variable, using the DEBUG level,
// printing the parameters v based on the format entered as a parameter. Unlike Debugf, this function will mask the
// first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugfMS("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func DebugfMS(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagMaskStart, v...)
}

// DebugfME is a function that records informational messages based on the Options global variable, using the DEBUG level,
// printing the parameters v based on the format entered as a parameter. Unlike Debugf, this function will mask the last
// half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugfME("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func DebugfME(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagMaskEnd, v...)
}

// DebugOptsH is a function that records informative messages based on the Options variable passed as a parameter,
// using the DEBUG level. Unlike DebugOpts, this function will hide all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.DebugOptsH(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: ****, ****, ***, *****
func DebugOptsH(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagHide, v...)
}

// DebugOptsMS is a function that records informative messages based on the Options variable passed as a parameter,
// using the DEBUG level. Unlike DebugOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.DebugOptsMS(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func DebugOptsMS(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagMaskStart, v...)
}

// DebugOptsME is a function that records informative messages based on the Options variable passed as a parameter,
// using the DEBUG level. Unlike DebugOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.DebugOptsME(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func DebugOptsME(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagMaskEnd, v...)
}

// DebugOptsf is a function that records informational messages based on the Options variable passed as a parameter,
// using the DEBUG level. Unlike DebugOpts, this function passes the format you want as a parameter, so the value passed
// will impact the format of the printed message.
//
// # Usage Example:
//
// logger.DebugOptsf("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func DebugOptsf(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

// DebugOptsfH is a function that records informational messages based on the Options variable passed as a parameter,
// using the DEBUG level, print the parameters v also based on the format entered as a parameter. Unlike DebugOptsf, this
// function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugOptsfH("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func DebugOptsfH(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

// DebugOptsfMS is a function that records informational messages based on the Options variable passed as a parameter,
// using the DEBUG level, print the parameters v also based on the format entered as a parameter. Unlike DebugOptsf, this
// function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugOptsfMS("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func DebugOptsfMS(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskStart, v...)
}

// DebugOptsfME is a function that records informational messages based on the Options variable passed as a parameter,
// using the DEBUG level, print the parameters v also based on the format entered as a parameter. Unlike DebugOptsf, this
// function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugOptsfME("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func DebugOptsfME(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskEnd, v...)
}

// DebugSkipCallerf is a function that logs informational messages based on the Options global variable, using the
// DEBUG level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike DebugSkipCaller, this function passes the format you want as a parameter, so the value passed will have an
// impact on the format of the printed message.
//
// # Usage Example:
//
// logger.DebugSkipCallerf("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func DebugSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagHide, v...)
}

// DebugSkipCallerH is a function that logs informational messages based on the Options global variable, using the
// DEBUG level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike DebugSkipCaller, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerH(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: ****, ****, ***, *****
func DebugSkipCallerH(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagHide, v...)
}

// DebugSkipCallerMS is a function that logs informational messages based on the Options global variable, using the
// DEBUG level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike DebugSkipCaller, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerMS(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func DebugSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

// DebugSkipCallerME is a function that logs informational messages based on the Options global variable, using the
// DEBUG level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike DebugSkipCaller, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerME(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func DebugSkipCallerME(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

// DebugSkipCallerfH is a function that writes informational messages based on the Options global variable, using the
// DEBUG level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike DebugSkipCallerf, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerfH("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func DebugSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagHide, v...)
}

// DebugSkipCallerfMS is a function that writes informational messages based on the Options global variable, using the
// DEBUG level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike DebugSkipCallerf, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerfMS("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func DebugSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

// DebugSkipCallerfME is a function that writes informational messages based on the Options global variable, using the
// DEBUG level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike DebugSkipCallerf, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerfME("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func DebugSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

// DebugSkipCallerOpts is a function that records informational messages, using the DEBUG level, the file name and line
// information in the message will be based on the value of the skipCaller parameter. Unlike DebugSkipCaller, this
// function passes Options as a parameter, and it will only be persisted in that message, therefore,
// this Options passed as a parameter will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.DebugSkipCallerOpts(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func DebugSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", "", v...)
}

// DebugSkipCallerOptsf is a function that records informational messages based on the Options variable passed as
// a parameter, using the DEBUG level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike DebugSkipCallerOpts, this function passes the format you want as a parameter, so the
// value passed will impact the format of the printed message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsf("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func DebugSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, "", v...)
}

// DebugSkipCallerOptsH is a function that records informational messages based on the Options variable passed as
// a parameter, using the DEBUG level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike DebugSkipCallerOpts, this function will hide all values passed as parameters v when
// printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsH(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:23: **** **** *** *****
func DebugSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagHide, v...)
}

// DebugSkipCallerOptsMS is a function that records informational messages based on the Options variable passed as
// a parameter, using the DEBUG level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike DebugSkipCallerOpts, this function will mask the first half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsMS(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func DebugSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagMaskStart, v...)
}

// DebugSkipCallerOptsME is a function that records informational messages based on the Options variable passed as
// a parameter, using the DEBUG level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike DebugSkipCallerOpts, this function will mask the last half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsME(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: te** tr** 1** 10***
func DebugSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

// DebugSkipCallerOptsfH is a function that records informational messages based on the Options variable passed
// as a parameter, using the DEBUG level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike DebugSkipCallerOptsf, this function will mask the last half of all values passed as v settings when
// printing a message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsfH("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func DebugSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagHide, v...)
}

// DebugSkipCallerOptsfMS is a function that records informational messages based on the Options variable passed
// as a parameter, using the DEBUG level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike DebugSkipCallerOptsf, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsfMS("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func DebugSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskStart, v...)
}

// DebugSkipCallerOptsfME is a function that records informational messages based on the Options variable passed
// as a parameter, using the DEBUG level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike DebugSkipCallerOptsf, this function will mask the last half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.DebugSkipCallerOptsfME("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [DEBUG 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func DebugSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
