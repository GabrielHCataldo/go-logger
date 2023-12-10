package logger

func Warning(v ...any) {
	printLog(levelWarning, 3, *opts, "", "", v...)
}

func WarningF(format string, v ...any) {
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

func WarningFH(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, loggerTagHide, v...)
}

func WarningFMS(format string, v ...any) {
	printLog(levelWarning, 3, *opts, format, loggerTagMaskStart, v...)
}

func WarningFME(format string, v ...any) {
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

func WarningOptsF(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagHide, v...)
}

func WarningOptsFH(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagHide, v...)
}

func WarningOptsFMS(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagMaskStart, v...)
}

func WarningOptsFME(format string, opts Options, v ...any) {
	printLog(levelWarning, 3, opts, format, loggerTagMaskEnd, v...)
}

func WarningSkipCallerF(format string, skipCaller int, v ...any) {
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

func WarningSkipCallerFH(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagHide, v...)
}

func WarningSkipCallerFMS(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagMaskStart, v...)
}

func WarningSkipCallerFME(format string, skipCaller int, v ...any) {
	printLog(levelWarning, skipCaller, *opts, format, loggerTagMaskEnd, v...)
}

func WarningSkipCallerOpts(skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, "", "", v...)
}

func WarningSkipCallerOptsF(format string, skipCaller int, opts Options, v ...any) {
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

func WarningSkipCallerOptsFH(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagHide, v...)
}

func WarningSkipCallerOptsFMS(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagMaskStart, v...)
}

func WarningSkipCallerOptsFME(format string, skipCaller int, opts Options, v ...any) {
	printLog(levelWarning, skipCaller, opts, format, loggerTagMaskEnd, v...)
}
