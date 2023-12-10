package logger

func Debug(v ...any) {
	printLog(levelDebug, 3, *opts, "", "", v...)
}

func Debugf(format string, v ...any) {
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

func DebugfH(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagHide, v...)
}

func DebugfMS(format string, v ...any) {
	printLog(levelDebug, 3, *opts, format, loggerTagMaskStart, v...)
}

func DebugfME(format string, v ...any) {
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

func DebugOptsf(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

func DebugOptsfH(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagHide, v...)
}

func DebugOptsfMS(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskStart, v...)
}

func DebugOptsfME(format string, opts Options, v ...any) {
	printLog(levelDebug, 3, opts, format, loggerTagMaskEnd, v...)
}

func DebugSkipCallerf(format string, skipCaller int, v ...any) {
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

func DebugSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagHide, v...)
}

func DebugSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func DebugSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelDebug, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func DebugSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, "", "", v...)
}

func DebugSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
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

func DebugSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagHide, v...)
}

func DebugSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func DebugSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelDebug, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
