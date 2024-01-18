package logger

import (
	"errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"time"
)

type customMapKey string

const customMapKeyString customMapKey = "customMapKeyV"

type test struct {
	Name                 string      `json:"name,omitempty"`
	BirthDate            time.Time   `json:"birthDate,omitempty" logger:"mask_end"`
	Document             string      `json:"document,omitempty" logger:"mask_start"`
	Password             string      `json:"password,omitempty" logger:"hide"`
	Married              bool        `json:"married"`
	Nil                  *bankTest   `json:"nil"`
	NilSlice             []any       `json:"nilSlice"`
	NilMap               map[any]any `json:"nilMap"`
	EmptyString          string      `json:"emptyString,omitempty"`
	EmptyBool            bool        `json:"emptyBool"`
	EmptyInt             int         `json:"emptyInt"`
	EmptyFloat           float64     `json:"emptyFloat"`
	Emails               []string    `json:"emails,omitempty"`
	PrivateEmails        []string    `json:"privateEmails,omitempty" logger:"mask_start"`
	Balances             []float32   `json:"balances,omitempty"`
	PrivateBalances      []float32   `json:"privateBalances,omitempty" logger:"mask_end"`
	TotalBalances        []float64   `json:"totalBalances,omitempty"`
	PrivateTotalBalances []float64   `json:"privateTotalBalances,omitempty" logger:"hide"`
	Booleans             []bool      `json:"booleans,omitempty"`
	PrivateBooleans      []bool      `json:"privateBooleans,omitempty" logger:"mask_end"`
	Bank                 bankTest    `json:"bank,omitempty"`
	PrivateBank          bankTest    `json:"privateBank,omitempty" logger:"mask_start"`
	Work                 workTest    `json:"work,omitempty"`
	PrivateWork          workTest    `json:"privateWork,omitempty" logger:"hide"`
	Home                 homeTest    `json:"home,omitempty"`
	PointerHome          *homeTest   `json:"pointerHome,omitempty"`
	MapInterface         map[any]any
	privateField         any
	HideFieldJson        any `json:"-"`
}

type bankTest struct {
	AccountDigits string  `json:"accountDigits,omitempty"`
	Account       string  `json:"account,omitempty"`
	Balance       float32 `json:"balance,omitempty"`
	TotalBalance  float64 `json:"totalBalance,omitempty"`
}

type workTest struct {
	Name          string  `json:"name,omitempty"`
	Address       string  `json:"address,omitempty"`
	AddressNumber string  `json:"addressNumber,omitempty"`
	Years         int64   `json:"years,omitempty"`
	Month         int32   `json:"month,omitempty"`
	Salary        float64 `json:"salary,omitempty"`
}

type homeTest struct {
	Address       string `json:"address,omitempty"`
	AddressNumber string `json:"addressNumber,omitempty"`
	Neighborhood  string `json:"neighborhood,omitempty"`
}

type tableTest struct {
	name       string
	format     string
	skipCaller int
	args       []any
}

func initSliceTest() []any {
	var a *any
	return []any{a, "test", 23, 23.2, true, false, initStructTest(), initMapTest(), nil, initPointerStructTest()}
}

func initMapTest() map[string]any {
	hTest := bankTest{
		AccountDigits: "123",
		Account:       "123981023",
		Balance:       30.89,
		TotalBalance:  200.17,
	}
	return map[string]any{
		"name":        "Foo Bar",
		"birthDate":   time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
		"document":    "02104996642",
		"password":    "12345672342342",
		"nil":         nil,
		"emptyString": "",
		"emptyInt":    0,
		"emptyFloat":  0.0,
		"married":     true,
		"emails":      []string{"gabriel@test.com", "gabrielcataldo.adm@gmail.com", "biel@test.org"},
		"bank":        hTest,
		"pointerBank": &hTest,
		"work": map[string]any{
			"name":          "Foo Bar Company",
			"address":       "Foo Bar Street",
			"addressNumber": "1234",
			"years":         2006,
			"month":         12,
			"salary":        1238.99,
			"anotherWork": workTest{
				Name:          "Foo Bar Company",
				Address:       "Foo Bar Street",
				AddressNumber: "1234",
				Years:         2006,
				Month:         12,
				Salary:        1238.99,
			},
		},
	}
}

func initStructTest() test {
	bTest := bankTest{
		AccountDigits: "123",
		Account:       "123981023",
		Balance:       30.89,
		TotalBalance:  200.17,
	}
	wTest := workTest{
		Name:          "Foo Bar Company",
		Address:       "Foo Bar Street",
		AddressNumber: "1234",
		Years:         2006,
		Month:         12,
		Salary:        1238.99,
	}
	hTest := homeTest{
		Address:       "Foo Bar Street",
		AddressNumber: "23",
		Neighborhood:  "Bar",
	}
	return test{
		Name:                 "Foo Bar",
		BirthDate:            time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
		Document:             "02104996642",
		Password:             "123456",
		Married:              true,
		Nil:                  nil,
		NilSlice:             nil,
		NilMap:               nil,
		Emails:               []string{"gabriel@test.com", "gabrielcataldo.231@gmail.com", "biel@test.org"},
		PrivateEmails:        []string{"gabriel@test.com", "gabrielcataldo.231@gmail.com", "biel@test.org"},
		Balances:             []float32{10.88, 11, 13.99, 12391.23, 23321},
		PrivateBalances:      []float32{10.88, 11, 13.99, 12391.23, 23321},
		TotalBalances:        []float64{20.88, 22, 23.99, 1023.23, 2400.1232},
		PrivateTotalBalances: []float64{20.88, 22, 23.99, 1023.23, 2400.1232},
		Booleans:             []bool{true, false, true, false},
		PrivateBooleans:      []bool{true, false, true, false},
		Bank:                 bTest,
		PrivateBank:          bTest,
		Work:                 wTest,
		PrivateWork:          wTest,
		Home:                 hTest,
		PointerHome:          &hTest,
		MapInterface: map[any]any{
			1:                  "test",
			"test":             1,
			true:               1,
			false:              "test",
			customMapKeyString: "custom value",
		},
		privateField:  "private not show",
		HideFieldJson: "not show",
	}
}

func initPointerStructTest() *test {
	sTest := initStructTest()
	return &sTest
}

func initPointerSliceTest() *[]any {
	s := initSliceTest()
	return &s
}

func initSliceBytes() []byte {
	return []byte("test slice bytes string")
}

func initPointerString() *string {
	s := "test pointer string"
	return &s
}

func initTables() []tableTest {
	return []tableTest{
		{
			"no arguments", "%v, %v, %v, %v, %v", 1, []any{nil, "", 0, map[string]any{}, []any{}},
		},
		{
			"normal arguments", "%s, %s, %s, %s, %s, %s, %s, %s last is %s", 1, []any{"test", true, 12.3, 200,
				initPointerString(), time.Now(), initPointerStructTest(), initPointerSliceTest(), initSliceBytes()},
		},
		{
			"map argument", "%v %v", 1, []any{"map:", initMapTest()},
		},
		{
			"struct argument", "%v %v", 1, []any{"struct:", initStructTest()},
		},
		{
			"slice argument", "%v %v", 1, []any{"slice:", initSliceTest()},
		},
		{
			"error argument", "%v %v", 1, []any{"error:", errors.New("new error test")},
		},
	}
}

func initOptionsTest() {
	ResetOptionsToDefault()
	SetOptions(getOptionsTest())
}

func initOptionsTestNil() {
	SetOptions(nil)
}

func getOptionsTest() *Options {
	return &Options{
		Mode:                       RandomMode(),
		DateFormat:                 RandomDateFormat(),
		CustomPrefixText:           RandomCustomPrefix(),
		CustomAfterPrefixText:      "after prefix",
		EnableAsynchronousMode:     helper.RandomBool(),
		UTC:                        helper.RandomBool(),
		DontPrintEmptyMessage:      helper.RandomBool(),
		RemoveSpace:                helper.RandomBool(),
		HideAllArgs:                helper.RandomBool(),
		HideArgDatetime:            helper.RandomBool(),
		HideArgCaller:              helper.RandomBool(),
		DisablePrefixColors:        helper.RandomBool(),
		EnableJsonMsgFieldForSlice: helper.RandomBool(),
	}
}
