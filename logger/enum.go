package logger

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"io"
	"math/rand"
	"os"
)

type BaseEnum interface {
	isEnumValid() bool
}

type DateFormat string
type Mode string
type level string

const (
	loggerTagHide      = "hide"
	loggerTagMaskStart = "mask_start"
	loggerTagMaskEnd   = "mask_end"
)
const (
	DateFormatFull24h                    DateFormat = "2006/01/02 15:04:05"
	DateFormatFull12h                    DateFormat = "2006/01/02 3:04:05PM"
	DateFormatNormal                     DateFormat = "2006/01/02"
	DateFormatTime24h                    DateFormat = "15:04:05"
	DateFormatTime12h                    DateFormat = "3:04:05PM"
	DateFormatTextDate                   DateFormat = "January 2, 2006"
	DateFormatTextDatetime24h            DateFormat = "January 2, 2006 15:04:05"
	DateFormatTextDatetime12h            DateFormat = "January 2, 2006 3:04:05PM"
	DateFormatTextDateWithWeekday        DateFormat = "Monday, January 2, 2006"
	DateFormatTextDatetime24hWithWeekday DateFormat = "Monday, January 2, 2006 15:04:05"
	DateFormatTextDatetime12hWithWeekday DateFormat = "Monday, January 2, 2006 3:04:05PM"
	DateFormatAbbrTextDate               DateFormat = "Jan 2 Mon"
	DateFormatAbbrTextDatetime24h        DateFormat = "Jan 2 Mon 15:04:05"
	DateFormatAbbrTextDatetime12h        DateFormat = "Jan 2 Mon 3:04:05PM"
)
const (
	ModeDefault Mode = "DEFAULT"
	ModeJson    Mode = "JSON"
)
const (
	StyleReset      = "\x1b[0m"
	StyleBold       = "\x1b[1m"
	StyleDim        = "\x1b[2m"
	StyleUnderscore = "\x1b[4m"
	StyleBlink      = "\x1b[5m"
	StyleReverse    = "\x1b[7m"
	StyleHidden     = "\x1b[8m"

	ForegroundBlack   = "\x1b[30m"
	ForegroundRed     = "\x1b[31m"
	ForegroundGreen   = "\x1b[32m"
	ForegroundYellow  = "\x1b[33m"
	ForegroundBlue    = "\x1b[34m"
	ForegroundMagenta = "\x1b[35m"
	ForegroundCyan    = "\x1b[36m"
	ForegroundWhite   = "\x1b[37m"

	BackgroundBlack   = "\x1b[40m"
	BackgroundRed     = "\x1b[41m"
	BackgroundGreen   = "\x1b[42m"
	BackgroundYellow  = "\x1b[43m"
	BackgroundBlue    = "\x1b[44m"
	BackgroundMagenta = "\x1b[45m"
	BackgroundCyan    = "\x1b[46m"
	BackgroundWhite   = "\x1b[47m"
)
const (
	InfoLevel  level = "INFO"
	DebugLevel level = "DEBUG"
	WarnLevel  level = "WARN"
	ErrorLevel level = "ERROR"
)

func (d DateFormat) isEnumValid() bool {
	switch d {
	case DateFormatFull24h, DateFormatFull12h, DateFormatNormal, DateFormatTime24h, DateFormatTime12h,
		DateFormatTextDate, DateFormatTextDatetime24h, DateFormatTextDatetime12h, DateFormatTextDateWithWeekday,
		DateFormatTextDatetime24hWithWeekday, DateFormatTextDatetime12hWithWeekday, DateFormatAbbrTextDate,
		DateFormatAbbrTextDatetime24h, DateFormatAbbrTextDatetime12h:
		return true
	}
	return false
}

func (d DateFormat) Format() string {
	if d.isEnumValid() {
		return string(d)
	}
	return string(DateFormatFull24h)
}

func (l level) acceptedLevels() []level {
	switch l {
	case InfoLevel:
		return []level{InfoLevel, WarnLevel, ErrorLevel}
	case WarnLevel:
		return []level{WarnLevel, ErrorLevel}
	case ErrorLevel:
		return []level{ErrorLevel}
	default:
		return []level{InfoLevel, DebugLevel, WarnLevel, ErrorLevel}
	}
}

func (l level) colorLevel() string {
	switch l {
	case InfoLevel:
		return ForegroundBlue
	case DebugLevel:
		return ForegroundCyan
	case WarnLevel:
		return ForegroundYellow
	case ErrorLevel:
		return ForegroundRed
	default:
		return StyleReset
	}
}

func (l level) colorMessage() string {
	switch l {
	default:
		return StyleReset
	}
}

func (l level) string() string {
	return string(l)
}

func (l level) dontAccepted(lvl level) bool {
	return helper.NotContains(l.acceptedLevels(), lvl)
}

func RandomOut() io.Writer {
	if helper.RandomBool() {
		return os.Stdout
	}
	return nil
}

func RandomLevel() level {
	ls := []level{InfoLevel, DebugLevel, WarnLevel, ErrorLevel}
	return ls[rand.Intn(4)]
}

func RandomMode() Mode {
	m := []Mode{ModeDefault, ModeJson}
	return m[rand.Intn(2)]
}

func RandomDateFormat() DateFormat {
	ds := []DateFormat{"",
		DateFormatFull24h, DateFormatFull12h, DateFormatNormal, DateFormatTime24h, DateFormatTime12h,
		DateFormatTextDate, DateFormatTextDatetime24h, DateFormatTextDatetime12h, DateFormatTextDateWithWeekday,
		DateFormatTextDatetime24hWithWeekday, DateFormatTextDatetime12hWithWeekday, DateFormatAbbrTextDate,
		DateFormatAbbrTextDatetime24h, DateFormatAbbrTextDatetime12h,
	}
	return ds[rand.Intn(15)]
}

func RandomCustomPrefix() string {
	var s string
	if helper.RandomBool() {
		s = "custom prefix"
	}
	return s
}
