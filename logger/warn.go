package logger

// Warn is a function that records informational messages based on the global Options variable, using the level WARN.
//
// # Usage Example:
//
// logger.Warn("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Warn(v ...any) {
	printLog(WarnLevel, 2, *opts, "", "", v...)
}

// Warnf is a function that records informational messages based on the Options global variable, using the WARN level.
// Unlike Warn, this function passes the format you want as a parameter, so the value passed will impact the format
// of the printed message.
//
// # Usage Example:
//
// logger.Warnf("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func Warnf(format string, v ...any) {
	printLog(WarnLevel, 2, *opts, format, "", v...)
}

// WarnOpts is a function that records informational messages, using the WARN level. Unlike Warn, this function passes
// Options as a parameter, and it will only be persisted in that message, therefore, this Options passed as a parameter
// will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.WarnOpts(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func WarnOpts(opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, "", "", v...)
}

// WarnSkipCaller is a function that logs informational messages based on the Options global variable, using the WARN
// level. Unlike Warn, this function passes skipCaller as a parameter, therefore the value passed directly impacts the
// file name and line when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCaller(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func WarnSkipCaller(skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, "", "", v...)
}

// WarnH is a function that records informational messages based on the Options global variable, using the WARN level.
// Unlike Warn, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnH("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func WarnH(v ...any) {
	printLog(WarnLevel, 2, *opts, "", loggerTagHide, v...)
}

// WarnMS is a function that records informational messages based on the Options global variable, using the WARN level.
// Unlike Warn, this function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnMS("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func WarnMS(v ...any) {
	printLog(WarnLevel, 2, *opts, "", loggerTagMaskStart, v...)
}

// WarnME is a function that records informational messages based on the Options global variable, using the WARN level.
// Unlike Warn, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnME("test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarnME(v ...any) {
	printLog(WarnLevel, 2, *opts, "", loggerTagMaskEnd, v...)
}

// WarnfH is a function that records informational messages based on the Options global variable, using the WARN level,
// printing the parameters v based on the format entered as a parameter. Unlike Warnf, this function will hide all
// values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnfH("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func WarnfH(format string, v ...any) {
	printLog(WarnLevel, 2, *opts, format, loggerTagHide, v...)
}

// WarnfMS is a function that records informational messages based on the Options global variable, using the WARN level,
// printing the parameters v based on the format entered as a parameter. Unlike Warnf, this function will mask the
// first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnfMS("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func WarnfMS(format string, v ...any) {
	printLog(WarnLevel, 2, *opts, format, loggerTagMaskStart, v...)
}

// WarnfME is a function that records informational messages based on the Options global variable, using the WARN level,
// printing the parameters v based on the format entered as a parameter. Unlike Warnf, this function will mask the last
// half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnfME("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func WarnfME(format string, v ...any) {
	printLog(WarnLevel, 2, *opts, format, loggerTagMaskEnd, v...)
}

// WarnOptsH is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARN level. Unlike WarnOpts, this function will hide all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarnOptsH(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: ****, ****, ***, *****
func WarnOptsH(opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, "", loggerTagHide, v...)
}

// WarnOptsMS is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARN level. Unlike WarnOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarnOptsMS(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func WarnOptsMS(opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, "", loggerTagMaskStart, v...)
}

// WarnOptsME is a function that records informative messages based on the Options variable passed as a parameter,
// using the WARN level. Unlike WarnOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarnOptsME(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarnOptsME(opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, "", loggerTagMaskEnd, v...)
}

// WarnOptsf is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARN level. Unlike WarnOpts, this function passes the format you want as a parameter, so the value passed
// will impact the format of the printed message.
//
// # Usage Example:
//
// logger.WarnOptsf("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func WarnOptsf(format string, opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, format, "", v...)
}

// WarnOptsfH is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARN level, print the parameters v also based on the format entered as a parameter. Unlike WarnOptsf, this
// function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnOptsfH("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func WarnOptsfH(format string, opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, format, loggerTagHide, v...)
}

// WarnOptsfMS is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARN level, print the parameters v also based on the format entered as a parameter. Unlike WarnOptsf, this
// function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnOptsfMS("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func WarnOptsfMS(format string, opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, format, loggerTagMaskStart, v...)
}

// WarnOptsfME is a function that records informational messages based on the Options variable passed as a parameter,
// using the WARN level, print the parameters v also based on the format entered as a parameter. Unlike WarnOptsf, this
// function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnOptsfME("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func WarnOptsfME(format string, opts Options, v ...any) {
	printLog(WarnLevel, 2, opts, format, loggerTagMaskEnd, v...)
}

// WarnSkipCallerf is a function that logs informational messages based on the Options global variable, using the
// WARN level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarnSkipCaller, this function passes the format you want as a parameter, so the value passed will have an
// impact on the format of the printed message.
//
// # Usage Example:
//
// logger.WarnSkipCallerf("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarnSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, format, "", v...)
}

// WarnSkipCallerH is a function that logs informational messages based on the Options global variable, using the
// WARN level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarnSkipCaller, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerH(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: ****, ****, ***, *****
func WarnSkipCallerH(skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, "", loggerTagHide, v...)
}

// WarnSkipCallerMS is a function that logs informational messages based on the Options global variable, using the
// WARN level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarnSkipCaller, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerMS(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func WarnSkipCallerMS(skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, "", loggerTagMaskStart, v...)
}

// WarnSkipCallerME is a function that logs informational messages based on the Options global variable, using the
// WARN level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike WarnSkipCaller, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerME(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func WarnSkipCallerME(skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, "", loggerTagMaskEnd, v...)
}

// WarnSkipCallerfH is a function that writes informational messages based on the Options global variable, using the
// WARN level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarnSkipCallerf, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerfH("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func WarnSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, format, loggerTagHide, v...)
}

// WarnSkipCallerfMS is a function that writes informational messages based on the Options global variable, using the
// WARN level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarnSkipCallerf, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerfMS("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func WarnSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, format, loggerTagMaskStart, v...)
}

// WarnSkipCallerfME is a function that writes informational messages based on the Options global variable, using the
// WARN level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike WarnSkipCallerf, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerfME("%s, %s, %s, last is %s", 2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func WarnSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(WarnLevel, skipCaller+1, *opts, format, loggerTagMaskEnd, v...)
}

// WarnSkipCallerOpts is a function that records informational messages, using the WARN level, the file name and line
// information in the message will be based on the value of the skipCaller parameter. Unlike WarnSkipCaller, this
// function passes Options as a parameter, and it will only be persisted in that message, therefore,
// this Options passed as a parameter will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.WarnSkipCallerOpts(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarnSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, "", "", v...)
}

// WarnSkipCallerOptsf is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARN level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarnSkipCallerOpts, this function passes the format you want as a parameter, so the
// value passed will impact the format of the printed message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsf("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func WarnSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, format, "", v...)
}

// WarnSkipCallerOptsH is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARN level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarnSkipCallerOpts, this function will hide all values passed as parameters v when
// printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsH(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:23: **** **** *** *****
func WarnSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, "", loggerTagHide, v...)
}

// WarnSkipCallerOptsMS is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARN level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarnSkipCallerOpts, this function will mask the first half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsMS(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func WarnSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, "", loggerTagMaskStart, v...)
}

// WarnSkipCallerOptsME is a function that records informational messages based on the Options variable passed as
// a parameter, using the WARN level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike WarnSkipCallerOpts, this function will mask the last half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsME(2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: te** tr** 1** 10***
func WarnSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, "", loggerTagMaskEnd, v...)
}

// WarnSkipCallerOptsfH is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARN level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarnSkipCallerOptsf, this function will mask the last half of all values passed as v settings when
// printing a message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsfH("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func WarnSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, format, loggerTagHide, v...)
}

// WarnSkipCallerOptsfMS is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARN level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarnSkipCallerOptsf, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsfMS("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func WarnSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, format, loggerTagMaskStart, v...)
}

// WarnSkipCallerOptsfME is a function that records informational messages based on the Options variable passed
// as a parameter, using the WARN level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike WarnSkipCallerOptsf, this function will mask the last half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.WarnSkipCallerOptsfME("%s, %s, %s, last is %s", 2, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [WARN 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func WarnSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(WarnLevel, skipCaller+1, opts, format, loggerTagMaskEnd, v...)
}
