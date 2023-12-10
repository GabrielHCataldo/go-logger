package logger

func Error(v ...any) {
	printLog(levelError, 3, *opts, "", "", v...)
}

func ErrorF(format string, v ...any) {
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

func ErrorFH(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagHide, v...)
}

func ErrorFMS(format string, v ...any) {
	printLog(levelError, 3, *opts, format, loggerTagMaskStart, v...)
}

func ErrorFME(format string, v ...any) {
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

func ErrorOptsF(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

func ErrorOptsFH(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagHide, v...)
}

func ErrorOptsFMS(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskStart, v...)
}

func ErrorOptsFME(format string, opts Options, v ...any) {
	printLog(levelError, 3, opts, format, loggerTagMaskEnd, v...)
}

func ErrorSkipCallerF(format string, skipCaller int, v ...any) {
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

func ErrorSkipCallerFH(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagHide, v...)
}

func ErrorSkipCallerFMS(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func ErrorSkipCallerFME(format string, skipCaller int, v ...any) {
	printLog(levelError, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func ErrorSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, "", "", v...)
}

func ErrorSkipCallerOptsF(format string, skipCaller int, opts Options, v ...any) {
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

func ErrorSkipCallerOptsFH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagHide, v...)
}

func ErrorSkipCallerOptsFMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func ErrorSkipCallerOptsFME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelError, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
