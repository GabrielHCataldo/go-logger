package logger

import (
	"testing"
	"time"
)

type test struct {
	Name         string    `json:"name,omitempty"`
	BirthDate    time.Time `json:"birthDate,omitempty" logger:"mask_end"`
	Document     string    `json:"document,omitempty" logger:"mask_start"`
	Password     string    `json:"password,omitempty" logger:"hide"`
	Emails       []string  `json:"emails,omitempty"`
	Balance      float32   `json:"balance,omitempty"`
	TotalBalance float64   `json:"totalBalance,omitempty" logger:"mask_start"`
}

func TestDebug(t *testing.T) {
	SetOptions(&options{
		Mode:                ModeDefault,
		DateFormat:          DateFormatFull24h,
		UTC:                 false,
		HideAllArgs:         false,
		HideArgDatetime:     false,
		HideArgCaller:       false,
		DisablePrefixColors: false,
	})
	structTest := test{
		Name:         "Foo Bar",
		BirthDate:    time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
		Document:     "02104996642",
		Password:     "123456",
		Emails:       []string{"gabriel@test.com", "gabrielcataldo.adm@gmail.com", "biel@test.org"},
		Balance:      30.89,
		TotalBalance: 200.17,
	}
	Debug("test test test", "2", "3", "json:", structTest)
}
