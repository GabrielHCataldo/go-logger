package logger

func Warning(v ...any) {
	printLog(levelWarning, 3, *opts, "", "", v...)
}

func Warningf(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, "", v...)
}

func WarningOpts(opts Options, v ...any) {
	printLog(levelWarning, 3, opts, "", loggerTagHide, v...)
}

func WarningSkipCaller(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, "", "", v...)
}

func WarningH(v ...any) {
	printLog(levelWarning, 3, *opts, "", loggerTagHide, v...)
}

func WarningMS(v ...any) {
	printLog(levelWarning, 3, *opts, "", loggerTagMaskStart, v...)
}

func WarningME(v ...any) {
	printLog(levelWarning, 3, *opts, "", loggerTagMaskEnd, v...)
}

func WarningfH(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, loggerTagHide, v...)
}

func WarningfMS(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, loggerTagMaskStart, v...)
}

func WarningfME(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, loggerTagMaskEnd, v...)
}

func WarningOptsH(opts Options, v ...any) {
	printLog(levelWarning, 3, opts, "", loggerTagHide, v...)
}

func WarningOptsMS(opts Options, v ...any) {
	printLog(levelWarning, 3, opts, "", loggerTagMaskStart, v...)
}

func WarningOptsME(opts Options, v ...any) {
	printLog(levelWarning, 3, opts, "", loggerTagMaskEnd, v...)
}

func WarningOptsf(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagHide, v...)
}

func WarningOptsfH(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagHide, v...)
}

func WarningOptsfMS(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagMaskStart, v...)
}

func WarningOptsfME(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagMaskEnd, v...)
}

func WarningSkipCallerf(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagHide, v...)
}

func WarningSkipCallerH(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, "", loggerTagHide, v...)
}

func WarningSkipCallerMS(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, "", loggerTagMaskStart, v...)
}

func WarningSkipCallerME(skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, "", loggerTagMaskEnd, v...)
}

func WarningSkipCallerfH(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagHide, v...)
}

func WarningSkipCallerfMS(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func WarningSkipCallerfME(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func WarningSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, "", "", v...)
}

func WarningSkipCallerOptsf(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, "", v...)
}

func WarningSkipCallerOptsH(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, "", loggerTagHide, v...)
}

func WarningSkipCallerOptsMS(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, "", loggerTagMaskStart, v...)
}

func WarningSkipCallerOptsME(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, "", loggerTagMaskEnd, v...)
}

func WarningSkipCallerOptsfH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagHide, v...)
}

func WarningSkipCallerOptsfMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func WarningSkipCallerOptsfME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
