package logger

// Error is a function that records informational messages based on the global Options variable, using the level ERROR.
//
// # Usage Example:
//
// logger.Error("test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Error(v ...any) {
	printLog(levelError, 3, *opts, "", "", v...)
}

// Errorf is a function that records informational messages based on the Options global variable, using the ERROR level.
// Unlike Error, this function passes the format you want as a parameter, so the value passed will impact the format
// of the printed message.
//
// # Usage Example:
//
// logger.Errorf("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func Errorf(format string, v ...any) {
	printLog(levelError, 3, *opts, format, "", v...)
}

// ErrorOpts is a function that records informational messages, using the ERROR level. Unlike Error, this function passes
// Options as a parameter, and it will only be persisted in that message, therefore, this Options passed as a parameter
// will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.ErrorOpts(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func ErrorOpts(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagHide, v...)
}

// ErrorSkipCaller is a function that logs informational messages based on the Options global variable, using the ERROR
// level. Unlike Error, this function passes skipCaller as a parameter, therefore the value passed directly impacts the
// file name and line when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCaller(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func ErrorSkipCaller(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", "", v...)
}

// ErrorH is a function that records informational messages based on the Options global variable, using the ERROR level.
// Unlike Error, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorH("test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func ErrorH(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagHide, v...)
}

// ErrorMS is a function that records informational messages based on the Options global variable, using the ERROR level.
// Unlike Error, this function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorMS("test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func ErrorMS(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagMaskStart, v...)
}

// ErrorME is a function that records informational messages based on the Options global variable, using the ERROR level.
// Unlike Error, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorME("test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func ErrorME(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagMaskEnd, v...)
}

// ErrorfH is a function that records informational messages based on the Options global variable, using the ERROR level,
// printing the parameters v based on the format entered as a parameter. Unlike Errorf, this function will hide all
// values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorfH("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func ErrorfH(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagHide, v...)
}

// ErrorfMS is a function that records informational messages based on the Options global variable, using the ERROR level,
// printing the parameters v based on the format entered as a parameter. Unlike Errorf, this function will mask the
// first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorfMS("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func ErrorfMS(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagMaskStart, v...)
}

// ErrorfME is a function that records informational messages based on the Options global variable, using the ERROR level,
// printing the parameters v based on the format entered as a parameter. Unlike Errorf, this function will mask the last
// half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorfME("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func ErrorfME(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagMaskEnd, v...)
}

// ErrorOptsH is a function that records informative messages based on the Options variable passed as a parameter,
// using the ERROR level. Unlike ErrorOpts, this function will hide all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.ErrorOptsH(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: ****, ****, ***, *****
func ErrorOptsH(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagHide, v...)
}

// ErrorOptsMS is a function that records informative messages based on the Options variable passed as a parameter,
// using the ERROR level. Unlike ErrorOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.ErrorOptsMS(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func ErrorOptsMS(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagMaskStart, v...)
}

// ErrorOptsME is a function that records informative messages based on the Options variable passed as a parameter,
// using the ERROR level. Unlike ErrorOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.ErrorOptsME(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func ErrorOptsME(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagMaskEnd, v...)
}

// ErrorOptsf is a function that records informational messages based on the Options variable passed as a parameter,
// using the ERROR level. Unlike ErrorOpts, this function passes the format you want as a parameter, so the value passed
// will impact the format of the printed message.
//
// # Usage Example:
//
// logger.ErrorOptsf("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func ErrorOptsf(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

// ErrorOptsfH is a function that records informational messages based on the Options variable passed as a parameter,
// using the ERROR level, print the parameters v also based on the format entered as a parameter. Unlike ErrorOptsf, this
// function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorOptsfH("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func ErrorOptsfH(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

// ErrorOptsfMS is a function that records informational messages based on the Options variable passed as a parameter,
// using the ERROR level, print the parameters v also based on the format entered as a parameter. Unlike ErrorOptsf, this
// function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorOptsfMS("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func ErrorOptsfMS(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskStart, v...)
}

// ErrorOptsfME is a function that records informational messages based on the Options variable passed as a parameter,
// using the ERROR level, print the parameters v also based on the format entered as a parameter. Unlike ErrorOptsf, this
// function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorOptsfME("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func ErrorOptsfME(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskEnd, v...)
}

// ErrorSkipCallerf is a function that logs informational messages based on the Options global variable, using the
// ERROR level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike ErrorSkipCaller, this function passes the format you want as a parameter, so the value passed will have an
// impact on the format of the printed message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerf("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func ErrorSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagHide, v...)
}

// ErrorSkipCallerH is a function that logs informational messages based on the Options global variable, using the
// ERROR level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike ErrorSkipCaller, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerH(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: ****, ****, ***, *****
func ErrorSkipCallerH(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagHide, v...)
}

// ErrorSkipCallerMS is a function that logs informational messages based on the Options global variable, using the
// ERROR level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike ErrorSkipCaller, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerMS(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func ErrorSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

// ErrorSkipCallerME is a function that logs informational messages based on the Options global variable, using the
// ERROR level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike ErrorSkipCaller, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerME(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func ErrorSkipCallerME(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

// ErrorSkipCallerfH is a function that writes informational messages based on the Options global variable, using the
// ERROR level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike ErrorSkipCallerf, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerfH("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func ErrorSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagHide, v...)
}

// ErrorSkipCallerfMS is a function that writes informational messages based on the Options global variable, using the
// ERROR level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike ErrorSkipCallerf, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerfMS("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func ErrorSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

// ErrorSkipCallerfME is a function that writes informational messages based on the Options global variable, using the
// ERROR level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike ErrorSkipCallerf, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerfME("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func ErrorSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

// ErrorSkipCallerOpts is a function that records informational messages, using the ERROR level, the file name and line
// information in the message will be based on the value of the skipCaller parameter. Unlike ErrorSkipCaller, this
// function passes Options as a parameter, and it will only be persisted in that message, therefore,
// this Options passed as a parameter will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOpts(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func ErrorSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", "", v...)
}

// ErrorSkipCallerOptsf is a function that records informational messages based on the Options variable passed as
// a parameter, using the ERROR level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike ErrorSkipCallerOpts, this function passes the format you want as a parameter, so the
// value passed will impact the format of the printed message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsf("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func ErrorSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, "", v...)
}

// ErrorSkipCallerOptsH is a function that records informational messages based on the Options variable passed as
// a parameter, using the ERROR level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike ErrorSkipCallerOpts, this function will hide all values passed as parameters v when
// printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsH(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:23: **** **** *** *****
func ErrorSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagHide, v...)
}

// ErrorSkipCallerOptsMS is a function that records informational messages based on the Options variable passed as
// a parameter, using the ERROR level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike ErrorSkipCallerOpts, this function will mask the first half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsMS(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func ErrorSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagMaskStart, v...)
}

// ErrorSkipCallerOptsME is a function that records informational messages based on the Options variable passed as
// a parameter, using the ERROR level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike ErrorSkipCallerOpts, this function will mask the last half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsME(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: te** tr** 1** 10***
func ErrorSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

// ErrorSkipCallerOptsfH is a function that records informational messages based on the Options variable passed
// as a parameter, using the ERROR level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike ErrorSkipCallerOptsf, this function will mask the last half of all values passed as v settings when
// printing a message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsfH("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func ErrorSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagHide, v...)
}

// ErrorSkipCallerOptsfMS is a function that records informational messages based on the Options variable passed
// as a parameter, using the ERROR level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike ErrorSkipCallerOptsf, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsfMS("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func ErrorSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskStart, v...)
}

// ErrorSkipCallerOptsfME is a function that records informational messages based on the Options variable passed
// as a parameter, using the ERROR level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike ErrorSkipCallerOptsf, this function will mask the last half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.ErrorSkipCallerOptsfME("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [ERROR 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func ErrorSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
