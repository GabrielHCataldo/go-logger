package logger

func Error(v ...any) {
	printLog(levelError, 3, *opts, "", "", v...)
}

func Errorf(format string, v ...any) {
	printLog(levelError, 3, *opts, format, "", v...)
}

func ErrorOpts(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagHide, v...)
}

func ErrorSkipCaller(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", "", v...)
}

func ErrorH(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagHide, v...)
}

func ErrorMS(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagMaskStart, v...)
}

func ErrorME(v ...any) {
	printLog(levelError, 3, *opts, "", loggerTagMaskEnd, v...)
}

func ErrorfH(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagHide, v...)
}

func ErrorfMS(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagMaskStart, v...)
}

func ErrorfME(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagMaskEnd, v...)
}

func ErrorOptsH(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagHide, v...)
}

func ErrorOptsMS(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagMaskStart, v...)
}

func ErrorOptsME(opts Options, v ...any) {
	printLog(levelError, 3, opts, "", loggerTagMaskEnd, v...)
}

func ErrorOptsf(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

func ErrorOptsfH(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

func ErrorOptsfMS(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskStart, v...)
}

func ErrorOptsfME(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskEnd, v...)
}

func ErrorSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagHide, v...)
}

func ErrorSkipCallerH(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagHide, v...)
}

func ErrorSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

func ErrorSkipCallerME(skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

func ErrorSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagHide, v...)
}

func ErrorSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func ErrorSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func ErrorSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", "", v...)
}

func ErrorSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, "", v...)
}

func ErrorSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagHide, v...)
}

func ErrorSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagMaskStart, v...)
}

func ErrorSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

func ErrorSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagHide, v...)
}

func ErrorSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func ErrorSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
