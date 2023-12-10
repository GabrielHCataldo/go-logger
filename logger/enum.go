package logger

import "math/rand"

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
	DateFormatFull24h                    DateFormat = "2006/01/02 15:04:09"
	DateFormatFull12h                    DateFormat = "2006/01/02 3:04:05PM"
	DateFormatNormal                     DateFormat = "2006/01/02"
	DateFormatTime24h                    DateFormat = "15:04:05"
	DateFormatTime12h                    DateFormat = "3:04:05 PM"
	DateFormatTextDate                   DateFormat = "January 2, 2022"
	DateFormatTextDatetime24h            DateFormat = "January 2, 2022 17:07:09"
	DateFormatTextDatetime12h            DateFormat = "January 2, 2022 3:04:05 PM"
	DateFormatTextDateWithWeekday        DateFormat = "Monday, January 2, 2006"
	DateFormatTextDatetime24hWithWeekday DateFormat = "Monday, January 2, 2006 17:07:09"
	DateFormatTextDatetime12hWithWeekday DateFormat = "Monday, January 2, 2006 3:04:05 PM"
	DateFormatAbbrTextDate               DateFormat = "Jan 2 Mon"
	DateFormatAbbrTextDatetime24h        DateFormat = "Jan 2 Mon 15:04:09"
	DateFormatAbbrTextDatetime12h        DateFormat = "Jan 2 Mon 3:04:05 PM"
)
const (
	ModeDefault Mode = "DEFAULT"
	ModeJson    Mode = "JSON"
)
const (
	levelInfo    level = "INFO"
	levelDebug   level = "DEBUG"
	levelWarning level = "WARNING"
	levelError   level = "ERROR"
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

func (l level) Color() string {
	switch l {
	case levelInfo:
		return "\u001b[34m"
	case levelDebug:
		return "\u001b[36m"
	case levelWarning:
		return "\u001b[33m"
	case levelError:
		return "\u001b[31m"
	default:
		return "\u001B[0m"
	}
}

func (l level) String() string {
	return string(l)
}

func randomMode() Mode {
	m := []Mode{ModeDefault, ModeJson}
	return m[rand.Intn(2)]
}

func randomDateFormat() DateFormat {
	ds := []DateFormat{"",
		DateFormatFull24h, DateFormatFull12h, DateFormatNormal, DateFormatTime24h, DateFormatTime12h,
		DateFormatTextDate, DateFormatTextDatetime24h, DateFormatTextDatetime12h, DateFormatTextDateWithWeekday,
		DateFormatTextDatetime24hWithWeekday, DateFormatTextDatetime12hWithWeekday, DateFormatAbbrTextDate,
		DateFormatAbbrTextDatetime24h, DateFormatAbbrTextDatetime12h,
	}
	return ds[rand.Intn(15)]
}
