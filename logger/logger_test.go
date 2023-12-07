package logger

import (
	"testing"
	"time"
)

type test struct {
	Name      string    `json:"name,omitempty"`
	BirthDate time.Time `json:"birthDate,omitempty"`
	Document  string    `json:"document,omitempty" logger:"mask"`
	Password  string    `json:"password,omitempty" logger:"hide"`
}

func TestDebug(t *testing.T) {
	SetOptions(&options{
		Mode:                ModeDefault,
		DateFormat:          DateFormatAbbrTextDatetime24h,
		UTC:                 false,
		HideAllArgs:         false,
		HideArgDatetime:     false,
		HideArgCaller:       false,
		DisablePrefixColors: false,
	})
	structTest := test{
		Name:      "Foo Bar",
		BirthDate: time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
		Document:  "02104996642",
		Password:  "123456",
	}
	Debug("test test test", "2", "3", "json:", structTest)
}
