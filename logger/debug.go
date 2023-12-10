package logger

func Debug(v ...any) {
	printLog(levelDebug, 3, *opts, "", "", v...)
}

func DebugF(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, "", v...)
}

func DebugOpts(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagHide, v...)
}

func DebugSkipCaller(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", "", v...)
}

func DebugH(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagHide, v...)
}

func DebugMS(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagMaskStart, v...)
}

func DebugME(v ...any) {
	printLog(levelDebug, 3, *opts, "", loggerTagMaskEnd, v...)
}

func DebugFH(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagHide, v...)
}

func DebugFMS(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagMaskStart, v...)
}

func DebugFME(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagMaskEnd, v...)
}

func DebugOptsH(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagHide, v...)
}

func DebugOptsMS(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagMaskStart, v...)
}

func DebugOptsME(opts Options, v ...any) {
	printLog(levelDebug, 3, opts, "", loggerTagMaskEnd, v...)
}

func DebugOptsF(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

func DebugOptsFH(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

func DebugOptsFMS(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskStart, v...)
}

func DebugOptsFME(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskEnd, v...)
}

func DebugSkipCallerF(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagHide, v...)
}

func DebugSkipCallerH(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagHide, v...)
}

func DebugSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

func DebugSkipCallerME(skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

func DebugSkipCallerFH(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagHide, v...)
}

func DebugSkipCallerFMS(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func DebugSkipCallerFME(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func DebugSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", "", v...)
}

func DebugSkipCallerOptsF(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, "", v...)
}

func DebugSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagHide, v...)
}

func DebugSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagMaskStart, v...)
}

func DebugSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

func DebugSkipCallerOptsFH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagHide, v...)
}

func DebugSkipCallerOptsFMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func DebugSkipCallerOptsFME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
