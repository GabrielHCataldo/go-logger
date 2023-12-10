package logger

// Info is a function that records informational messages based on the global Options variable, using the level INFO.
//
// # Usage Example:
//
// logger.Info("test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func Info(v ...any) {
	printLog(levelInfo, 3, *opts, "", "", v...)
}

// Infof is a function that records informational messages based on the Options global variable, using the INFO level.
// Unlike Info, this function passes the format you want as a parameter, so the value passed will impact the format
// of the printed message.
//
// # Usage Example:
//
// logger.Infof("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func Infof(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, "", v...)
}

// InfoOpts is a function that records informational messages, using the INFO level. Unlike Info, this function passes
// Options as a parameter, and it will only be persisted in that message, therefore, this Options passed as a parameter
// will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.InfoOpts(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: test true 112 10.99
func InfoOpts(opts Options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagHide, v...)
}

// InfoSkipCaller is a function that logs informational messages based on the Options global variable, using the INFO
// level. Unlike Info, this function passes skipCaller as a parameter, therefore the value passed directly impacts the
// file name and line when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCaller(2, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:139: te** tr** 1** 10***
func InfoSkipCaller(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", "", v...)
}

// InfoH is a function that records informational messages based on the Options global variable, using the INFO level.
// Unlike Info, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoH("test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **** **** *** *****
func InfoH(v ...any) {
	printLog(levelInfo, 3, *opts, "", loggerTagHide, v...)
}

// InfoMS is a function that records informational messages based on the Options global variable, using the INFO level.
// Unlike Info, this function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoMS("test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func InfoMS(v ...any) {
	printLog(levelInfo, 3, *opts, "", loggerTagMaskStart, v...)
}

// InfoME is a function that records informational messages based on the Options global variable, using the INFO level.
// Unlike Info, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoME("test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func InfoME(v ...any) {
	printLog(levelInfo, 3, *opts, "", loggerTagMaskEnd, v...)
}

// InfofH is a function that records informational messages based on the Options global variable, using the INFO level,
// printing the parameters v based on the format entered as a parameter. Unlike Infof, this function will hide all
// values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfofH("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func InfofH(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagHide, v...)
}

// InfofMS is a function that records informational messages based on the Options global variable, using the INFO level,
// printing the parameters v based on the format entered as a parameter. Unlike Infof, this function will mask the
// first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfofMS("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func InfofMS(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagMaskStart, v...)
}

// InfofME is a function that records informational messages based on the Options global variable, using the INFO level,
// printing the parameters v based on the format entered as a parameter. Unlike Infof, this function will mask the last
// half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfofME("%s, %s, %s, last is %s", "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func InfofME(format string, v ...any) {
	printLog(levelInfo, 3, *opts, format, loggerTagMaskEnd, v...)
}

// InfoOptsH is a function that records informative messages based on the Options variable passed as a parameter,
// using the INFO level. Unlike InfoOpts, this function will hide all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.InfoOptsH(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: ****, ****, ***, *****
func InfoOptsH(opts Options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagHide, v...)
}

// InfoOptsMS is a function that records informative messages based on the Options variable passed as a parameter,
// using the INFO level. Unlike InfoOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.InfoOptsMS(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **st **ue **2 ***99
func InfoOptsMS(opts Options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagMaskStart, v...)
}

// InfoOptsME is a function that records informative messages based on the Options variable passed as a parameter,
// using the INFO level. Unlike InfoOpts, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.InfoOptsME(Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func InfoOptsME(opts Options, v ...any) {
	printLog(levelInfo, 3, opts, "", loggerTagMaskEnd, v...)
}

// InfoOptsf is a function that records informational messages based on the Options variable passed as a parameter,
// using the INFO level. Unlike InfoOpts, this function passes the format you want as a parameter, so the value passed
// will impact the format of the printed message.
//
// # Usage Example:
//
// logger.InfoOptsf("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: test, true, 112, last is 10.99
func InfoOptsf(format string, opts Options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagHide, v...)
}

// InfoOptsfH is a function that records informational messages based on the Options variable passed as a parameter,
// using the INFO level, print the parameters v also based on the format entered as a parameter. Unlike InfoOptsf, this
// function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoOptsfH("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: ****, ****, ***, last is *****
func InfoOptsfH(format string, opts Options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagHide, v...)
}

// InfoOptsfMS is a function that records informational messages based on the Options variable passed as a parameter,
// using the INFO level, print the parameters v also based on the format entered as a parameter. Unlike InfoOptsf, this
// function will mask the first half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoOptsfMS("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: **st, **ue, **2, last is ***99
func InfoOptsfMS(format string, opts Options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagMaskStart, v...)
}

// InfoOptsfME is a function that records informational messages based on the Options variable passed as a parameter,
// using the INFO level, print the parameters v also based on the format entered as a parameter. Unlike InfoOptsf, this
// function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoOptsfME("%s, %s, %s, last is %s", Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te**, tr**, 1**, last is 10***
func InfoOptsfME(format string, opts Options, v ...any) {
	printLog(levelInfo, 3, opts, format, loggerTagMaskEnd, v...)
}

// InfoSkipCallerf is a function that logs informational messages based on the Options global variable, using the
// INFO level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike InfoSkipCaller, this function passes the format you want as a parameter, so the value passed will have an
// impact on the format of the printed message.
//
// # Usage Example:
//
// logger.InfoSkipCallerf("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func InfoSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagHide, v...)
}

// InfoSkipCallerH is a function that logs informational messages based on the Options global variable, using the
// INFO level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike InfoSkipCaller, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerH(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: ****, ****, ***, *****
func InfoSkipCallerH(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagHide, v...)
}

// InfoSkipCallerMS is a function that logs informational messages based on the Options global variable, using the
// INFO level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike InfoSkipCaller, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerMS(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func InfoSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

// InfoSkipCallerME is a function that logs informational messages based on the Options global variable, using the
// INFO level, the file name and line information in the message will be based on the value of the skipCaller parameter.
// Unlike InfoSkipCaller, this function will mask the last half of all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerME(3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] example.go:239: te** tr** 1** 10***
func InfoSkipCallerME(skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

// InfoSkipCallerfH is a function that writes informational messages based on the Options global variable, using the
// INFO level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike InfoSkipCallerf, this function will hide all values passed as parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerfH("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func InfoSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagHide, v...)
}

// InfoSkipCallerfMS is a function that writes informational messages based on the Options global variable, using the
// INFO level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike InfoSkipCallerf, this function will mask the first half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerfMS("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func InfoSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

// InfoSkipCallerfME is a function that writes informational messages based on the Options global variable, using the
// INFO level. The file name and the line information in the message will be based on the value of the skipCaller
// parameter, this function also prints the v parameters based on the format passed in the format parameter.
// Unlike InfoSkipCallerf, this function will mask the last half of all values passed as parameters v when printing
// the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerfME("%s, %s, %s, last is %s", 3, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func InfoSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelInfo, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

// InfoSkipCallerOpts is a function that records informational messages, using the INFO level, the file name and line
// information in the message will be based on the value of the skipCaller parameter. Unlike InfoSkipCaller, this
// function passes Options as a parameter, and it will only be persisted in that message, therefore,
// this Options passed as a parameter will not impact other flows using global Options.
//
// # Usage Example:
//
// logger.InfoSkipCallerOpts(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func InfoSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", "", v...)
}

// InfoSkipCallerOptsf is a function that records informational messages based on the Options variable passed as
// a parameter, using the INFO level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike InfoSkipCallerOpts, this function passes the format you want as a parameter, so the
// value passed will impact the format of the printed message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsf("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: test, true, 112, last is 10.99
func InfoSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, "", v...)
}

// InfoSkipCallerOptsH is a function that records informational messages based on the Options variable passed as
// a parameter, using the INFO level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike InfoSkipCallerOpts, this function will hide all values passed as parameters v when
// printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsH(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:23: **** **** *** *****
func InfoSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagHide, v...)
}

// InfoSkipCallerOptsMS is a function that records informational messages based on the Options variable passed as
// a parameter, using the INFO level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike InfoSkipCallerOpts, this function will mask the first half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsMS(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: **st **ue **2 ***99
func InfoSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagMaskStart, v...)
}

// InfoSkipCallerOptsME is a function that records informational messages based on the Options variable passed as
// a parameter, using the INFO level, the file name and line information in the message will be based on the value of
// the skipCaller parameter. Unlike InfoSkipCallerOpts, this function will mask the last half of all values passed as
// parameters v when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsME(3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: te** tr** 1** 10***
func InfoSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

// InfoSkipCallerOptsfH is a function that records informational messages based on the Options variable passed
// as a parameter, using the INFO level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike InfoSkipCallerOptsf, this function will mask the last half of all values passed as v settings when
// printing a message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsfH("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:23: ****, ****, ***, last is *****
func InfoSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagHide, v...)
}

// InfoSkipCallerOptsfMS is a function that records informational messages based on the Options variable passed
// as a parameter, using the INFO level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike InfoSkipCallerOptsf, this function will mask the first half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsfMS("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: **st, **ue, **2, last is ***99
func InfoSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagMaskStart, v...)
}

// InfoSkipCallerOptsfME is a function that records informational messages based on the Options variable passed
// as a parameter, using the INFO level. The file name and line information in the message will be based on the value
// of the skipCaller parameter, this function also prints the v parameters based on the format passed in the format
// parameter. Unlike InfoSkipCallerOptsf, this function will mask the last half of all values passed as parameters v
// when printing the message.
//
// # Usage Example:
//
// logger.InfoSkipCallerOptsfME("%s, %s, %s, last is %s", 3, Options{}, "test", true, 112, 10.99)
//
// # Result (default):
//
// [INFO 2023/12/09 19:26:09] subexample.go:239: te**, tr**, 1**, last is 10***
func InfoSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelInfo, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
