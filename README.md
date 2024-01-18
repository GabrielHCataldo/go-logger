go-logger
=================
<img align="right" src="gopher-logger.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.1.3-vividgreen.svg)](https://github.com/GabrielHCataldo/go-logger/releases/tag/v1.1.3)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-logger)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-logger)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-logger/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-logger?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-logger/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-logger)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-logger?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-logger/logger)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-logger/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-logger/actions&#41;)

[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-logger/logger/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-logger/logger?badge&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-logger/logger&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-logger&#41;)

The go-logger project came to facilitate and increase your project's log customization capacity, being the most complete
from the market. Below are functions to use:

- Log Level (Info, Debug, Warning, Error)
- Different message modes (Default or Json)
- Information mask for all types
- Information hide for all types
- Mask customization (Mask start or Mask end)
- Print log asynchronously
- Customization globally or locally
- Date and time format customization
- Have the option to print only the message has content
- Simplicity in calls

Installation
------------

Use go get.

	go get github.com/GabrielHCataldo/go-logger

Then import the go-logger package into your own code.

```go
import "github.com/GabrielHCataldo/go-logger/logger"
```

Usability and documentation
------------
**IMPORTANT**: Always check the documentation in the structures and functions fields.
For more details on the examples, visit [All examples link](https://github/GabrielHCataldo/go-logger/blob/main/_example/main.go)

Using go-logger's features is very easy, let's start with a basic example:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    basicMsg := getBasicMsg()
    logger.Info(basicMsg...)
    logger.Debug(basicMsg...)
    logger.Warning(basicMsg...)
    logger.Error(basicMsg...)
}

func getBasicMsg() []any {
    return []any{"basic example with empty any values", nil, "", 0, map[string]any{}, []any{}}
}
```

Outputs:

    [INFO 2023/12/10 17:23:09] main.go:5: basic example with empty any values  0 {} null
    [DEBUG 2023/12/10 17:23:09] main.go:6: basic example with empty any values  0 {} null
    [WARNING 2023/12/10 17:23:09] main.go:7: basic example with empty any values  0 {} null
    [ERROR 2023/12/10 17:23:09] main.go:8: basic example with empty any values  0 {} null

We can modify this default structure to a JSON log structure, just edit its global options like
below:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    // set on global options
    logger.SetOptions(&logger.Options{
        // Print mode (default: ModeDefault)
        Mode: logger.ModeJson,
        // Argument date format (default: DateFormatFull24h)
        DateFormat: logger.DateFormatFull24h,
        // Custom prefix text
        CustomPrefixText: "",
        // Custom after prefix text (only if Mode is ModeDefault)
        CustomAfterPrefixText: "",
        // Enable asynchronous printing mode (default: false)
        EnableAsynchronousMode: false,
        // Enable argument date to be UTC (default: false)
        UTC: false,
        // Enable to not print empty message (default: false)
        DontPrintEmptyMessage: false,
        // Enable to remove spaces between parameters (default: false)
        RemoveSpace: false,
        // If true will hide all datetime and prefix arguments (default: false)
        HideAllArgs: false,
        // If true it will hide the datetime arguments (default: false)
        HideArgDatetime: false,
        // If true, it will hide the caller arguments (default: false)
        HideArgCaller: false,
        // If true, it will disable all argument and prefix colors (default: false)
        DisablePrefixColors: false,
        // If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
        //
        // IMPORTANT: If true, the format parameter will not work
        EnableJsonMsgFieldForSlice: false,
    })
    basicMsg := getBasicMsg()
    logger.Info(basicMsg...)
    logger.Debug(basicMsg...)
    logger.Warning(basicMsg...)
    logger.Error(basicMsg...)
}

func getBasicMsg() []any {
    return []any{"basic example with empty any values", nil, "", 0, map[string]any{}, []any{}}
}
```

Outputs:

    {"level":"INFO","datetime":"2023/12/10 17:18:09","file":"main.go","func":"main.main","line":"9","msg":"basic example with empty any values  0 {} null"}
    {"level":"DEBUG","datetime":"2023/12/10 17:18:09","file":"main.go","func":"main.main","line":"10","msg":"basic example with empty any values  0 {} null"}
    {"level":"WARNING","datetime":"2023/12/10 17:18:09","file":"main.go","func":"main.main","line":"11","msg":"basic example with empty any values  0 {} null"}
    {"level":"ERROR","datetime":"2023/12/10 17:18:09","file":"main.go","func":"main.main","line":"12","msg":"basic example with empty any values  0 {} null"}

You can also pass custom options as a parameter so as not to impact the defined global settings,
just use the functions with the **Opts** pattern, see:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    opts := logger.Options{
        // Print mode (default: ModeDefault)
        Mode: logger.ModeJson,
        // Argument date format (default: DateFormatFull24h)
        DateFormat: logger.DateFormatFull24h,
        // Custom prefix text
        CustomPrefixText: "",
        // Custom after prefix text (only if Mode is ModeDefault)
        CustomAfterPrefixText: "",
        // Enable asynchronous printing mode (default: false)
        EnableAsynchronousMode: false,
        // Enable argument date to be UTC (default: false)
        UTC: false,
        // Enable to not print empty message (default: false)
        DontPrintEmptyMessage: false,
        // Enable to remove spaces between parameters (default: false)
        RemoveSpace: false,
        // If true will hide all datetime and prefix arguments (default: false)
        HideAllArgs: false,
        // If true it will hide the datetime arguments (default: false)
        HideArgDatetime: false,
        // If true, it will hide the caller arguments (default: false)
        HideArgCaller: false,
        // If true, it will disable all argument and prefix colors (default: false)
        DisablePrefixColors: false,
        // If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
        //
        // IMPORTANT: If true, the format parameter will not work
        EnableJsonMsgFieldForSlice: false,
    }
    basicMsg := getBasicMsg()
    logger.InfoOpts(opts, basicMsg...)
    // field msg json to slice
    opts.EnableJsonMsgFieldForSlice = true
    logger.DebugOpts(opts, basicMsg...)
    logger.Warning(basicMsg...)
    logger.Error(basicMsg...)
}

func getBasicMsg() []any {
    return []any{"basic example with empty any values", nil, "", 0, map[string]any{}, []any{}}
}
```

Outputs:

    {"level":"INFO","datetime":"2023/12/11 08:17:09","file":"main.go","func":"main.main","line":"50","msg":"basic example with any values text string 1 12.213 true 2023-12-11T08:17:14-03:00 {\"int\":1,\"float\":12.213,\"string\":\"text string\",\"bool\":true,\"time\":\"2023-12-11T08:17:14-03:00\",\"nilValue\":null} [\"text string\",1,12.213,true,null,\"2023-12-11T08:17:14-03:00\"]"}
    {"level":"DEBUG","datetime":"2023/12/11 08:17:09","file":"main.go","func":"main.main","line":"53","msg":["basic example with any values","text string",1,12.213,true,"2023-12-11T08:17:14-03:00",{"bool":true,"float":12.213,"int":1,"nilValue":null,"string":"text string","time":"2023-12-11T08:17:14-03:00"},["text string",1,12.213,true,null,"2023-12-11T08:17:14-03:00"]]}
    [WARNING 2023/12/11 08:17:09] main.go:54: basic example with any values text string 1 12.213 true 2023-12-11T08:17:14-03:00 {"float":12.213,"string":"text string","bool":true,"time":"2023-12-11T08:17:14-03:00","nilValue":null,"int":1} ["text string",1,12.213,true,null,"2023-12-11T08:17:14-03:00"]
    [ERROR 2023/12/11 08:17:09] main.go:55: basic example with any values text string 1 12.213 true 2023-12-11T08:17:14-03:00 {"time":"2023-12-11T08:17:14-03:00","nilValue":null,"int":1,"float":12.213,"string":"text string","bool":true} ["text string",1,12.213,true,null,"2023-12-11T08:17:14-03:00"]

In go-logger we have the same **format** pattern as a parameter in some native functions,
This value formats the parameters passed in **v**, works both in standard mode and in default mode.
JSON, just call the functions with the **f** pattern, see some examples below:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    customOpts := logger.Options{
        // Print mode (default: ModeDefault)
        Mode: logger.ModeJson,
        // Argument date format (default: DateFormatFull24h)
        DateFormat: logger.DateFormatAbbrTextDatetime12h,
        // Custom prefix text
        CustomPrefixText: "",
        // Custom after prefix text (only if Mode is ModeDefault)
        CustomAfterPrefixText: "",
        // Enable asynchronous printing mode (default: false)
        EnableAsynchronousMode: false,
        // Enable argument date to be UTC (default: false)
        UTC: false,
        // Enable to not print empty message (default: false)
        DontPrintEmptyMessage: false,
        // Enable to remove spaces between parameters (default: false)
        RemoveSpace: false,
        // If true will hide all datetime and prefix arguments (default: false)
        HideAllArgs: false,
        // If true it will hide the datetime arguments (default: false)
        HideArgDatetime: false,
        // If true, it will hide the caller arguments (default: false)
        HideArgCaller: false,
        // If true, it will disable all argument and prefix colors (default: false)
        DisablePrefixColors: false,
        // If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
        //
        // IMPORTANT: If true, the format parameter will not work
        EnableJsonMsgFieldForSlice: false,
    }
    format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
    msg := getBasicMsg()
    logger.Infof(format, msg...)
    logger.DebugOptsf(format, customOpts, msg...)
    logger.WarningfMS(format, msg...)
    logger.ErrorOptsfMS(format, customOpts, msg...)
}

func getBasicMsg() []any {
    s := []any{
        "text string",
        1,
        12.213,
        true,
        nil,
        time.Now(),
    }
    m := map[string]any{
        "int":      1,
        "float":    12.213,
        "string":   "text string",
        "bool":     true,
        "nilValue": nil,
        "time":     time.Now(),
    }
    return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}
```

Outputs: 

    [INFO 2023/12/11 07:53:09] main.go:9: basic example with any values, text string, 1, 12.213, true, 2023-12-11T07:53:41-03:00, {"nilValue":null,"int":1,"float":12.213,"string":"text string","bool":true,"time":"2023-12-11T07:53:41-03:00"}, last is ["text string",1,12.213,true,null,"2023-12-11T07:53:41-03:00"]
    {"level":"DEBUG","datetime":"Dec 11 Mon 7:53:41 AM","file":"types.go","func":"main.main","line":"10","msg":"*****************************, ***********, *, ******, ****, *************************, {\"time\":\"*************************\",\"nilValue\":null,\"int\":\"*\",\"float\":\"******\",\"string\":\"***********\",\"bool\":\"****\"}, last is [\"***********\",\"*\",\"******\",\"****\",null,\"*************************\"]"}
    [WARNING 2023/12/11 07:53:09] main.go:11: **************with any values, *****string, *, ***213, **ue, ************7:53:41-03:00, {"nilValue":null,"int":"*","float":"***213","string":"*****string","bool":"**ue","time":"************7:53:41-03:00"}, last is ["*****string","*","***213","**ue",null,"************7:53:41-03:00"]
    {"level":"ERROR","datetime":"Dec 11 Mon 7:53:41 AM","file":"types.go","func":"main.main","line":"12","msg":"**************with any values, *****string, *, ***213, **ue, ************7:53:41-03:00, {\"float\":\"***213\",\"string\":\"*****string\",\"bool\":\"**ue\",\"time\":\"************7:53:41-03:00\",\"nilValue\":null,\"int\":\"*\"}, last is [\"*****string\",\"*\",\"***213\",\"**ue\",null,\"************7:53:41-03:00\"]"}

For a clearer debug, we have in our argument the caller's name and line, to be more precise in some cases
sub calls, for example, you can call functions with the **SkipCaller** pattern and pass the number you want
as a parameter to skip the caller, see an example below:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    subFunc()
}

func subFunc() {
    customOpts := logger.Options{
        // Print mode (default: ModeDefault)
        Mode: logger.ModeJson,
        // Argument date format (default: DateFormatFull24h)
        DateFormat: logger.DateFormatAbbrTextDatetime12h,
        // Custom prefix text
        CustomPrefixText: "",
        // Custom after prefix text (only if Mode is ModeDefault)
        CustomAfterPrefixText: "",
        // Enable asynchronous printing mode (default: false)
        EnableAsynchronousMode: false,
        // Enable argument date to be UTC (default: false)
        UTC: false,
        // Enable to not print empty message (default: false)
        DontPrintEmptyMessage: false,
        // Enable to remove spaces between parameters (default: false)
        RemoveSpace: false,
        // If true will hide all datetime and prefix arguments (default: false)
        HideAllArgs: false,
        // If true it will hide the datetime arguments (default: false)
        HideArgDatetime: false,
        // If true, it will hide the caller arguments (default: false)
        HideArgCaller: false,
        // If true, it will disable all argument and prefix colors (default: false)
        DisablePrefixColors: false,
        // If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
        //
        // IMPORTANT: If true, the format parameter will not work
        EnableJsonMsgFieldForSlice: false,
    }
    format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
    msg := getBasicMsg()
    logger.InfoSkipCaller(1, msg...)
    logger.DebugSkipCallerOpts(1, customOpts, msg...)
    logger.WarningSkipCallerOptsf(format, 2, customOpts, msg...)
    logger.ErrorSkipCaller(2, msg...)
}

func getBasicMsg() []any {
    s := []any{
        "text string",
        1,
        12.213,
        true,
        nil,
        time.Now(),
    }
    m := map[string]any{
        "int":      1,
        "float":    12.213,
        "string":   "text string",
        "bool":     true,
        "nilValue": nil,
        "time":     time.Now(),
    }
    return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}
```

Outputs:
    
    [INFO 2023/12/11 08:37:09] main.go:23: basic example with any values text string 1 12.213 true 2023-12-11T08:37:01-03:00 {"nilValue":null,"int":1,"float":12.213,"string":"text string","bool":true,"time":"2023-12-11T08:37:01-03:00"} ["text string",1,12.213,true,null,"2023-12-11T08:37:01-03:00"]
    {"level":"DEBUG","datetime":"2023/12/11 8:37:01AM","file":"main.go","func":"main.subFunc","line":"24","msg":"basic example with any values text string 1 12.213 true 2023-12-11T08:37:01-03:00 {\"time\":\"2023-12-11T08:37:01-03:00\",\"nilValue\":null,\"int\":1,\"float\":12.213,\"string\":\"text string\",\"bool\":true} [\"text string\",1,12.213,true,null,\"2023-12-11T08:37:01-03:00\"]"}
    {"level":"WARNING","datetime":"2023/12/11 8:37:01AM","file":"main.go","func":"main.main","line":"4","msg":"basic example with any values, text string, 1, 12.213, true, 2023-12-11T08:37:01-03:00, {\"nilValue\":null,\"int\":1,\"float\":12.213,\"string\":\"text string\",\"bool\":true,\"time\":\"2023-12-11T08:37:01-03:00\"}, last is [\"text string\",1,12.213,true,null,\"2023-12-11T08:37:01-03:00\"]"}
    [ERROR 2023/12/11 08:37:09] main.go:4: basic example with any values text string 1 12.213 true 2023-12-11T08:37:01-03:00 {"bool":true,"time":"2023-12-11T08:37:01-03:00","nilValue":null,"int":1,"float":12.213,"string":"text string"} ["text string",1,12.213,true,null,"2023-12-11T08:37:01-03:00"]

We have a complete solution to hide/mask values that works for all types, example
Below we will use a structure, where we will add the **logger** tag with the values **hide** | **mask_start** | **mask_end**
indicating how we want the value to be hidden. See the example below:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

type test struct {
    Name          string    `json:"name,omitempty"`
    BirthDate     time.Time `json:"birthDate,omitempty"`
    Document      string    `json:"document,omitempty" logger:"mask_start"`
    Emails        []string  `json:"emails,omitempty"`
    Balances      []float32 `json:"balances,omitempty" logger:"mask_end"`
    TotalBalances []float64 `json:"totalBalances,omitempty"`
    Booleans      []bool    `json:"booleans,omitempty"`
    Bank          bankTest  `json:"bank,omitempty"`
}

type bankTest struct {
    AccountDigits string  `json:"accountDigits,omitempty"`
    Account       string  `json:"account,omitempty"`
    Balance       float32 `json:"balance,omitempty" logger:"hide"`
    TotalBalance  float64 `json:"totalBalance,omitempty"`
}

func main() {
    testStruct := getTestStruct()
    msg := []any{"test mask/hide struct:", testStruct}
    logger.Info(msg...)
    logger.Debug(msg...)
    logger.Warning(msg...)
    logger.Error(msg...)
}

func getTestStruct() test {
    bank := bankTest{
        AccountDigits: "123",
        Account:       "123981023",
        Balance:       30.89,
        TotalBalance:  200.17,
    }
    return test{
        Name:      "Foo Bar",
        BirthDate: time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local),
        Document:  "02104996642",
        Emails:    []string{"gabriel@test.com", "gabrielcataldo.231@gmail.com", "biel@test.org"},
        Balances:  []float32{10.88, 11, 13.99, 12391.23, 23321},
        Bank:      bank,
    }
}    
```
    
Outputs:

    [INFO 2023/12/11 06:21:09] main.go:20: test mask/hide struct: {"name":"Foo Bar","birthDate":"1999-01-21T00:00:00-02:00","document":"*****996642","emails":["gabriel@test.com","gabrielcataldo.231@gmail.com","biel@test.org"],"balances":["10***","1*","13***","1239****","23***"],"bank":{"account":"123981023","accountDigits":"123","balance":"*****","totalBalance":200.17}}
    [DEBUG 2023/12/11 06:21:09] main.go:21: test mask/hide struct: {"name":"Foo Bar","birthDate":"1999-01-21T00:00:00-02:00","document":"*****996642","emails":["gabriel@test.com","gabrielcataldo.231@gmail.com","biel@test.org"],"balances":["10***","1*","13***","1239****","23***"],"bank":{"account":"123981023","accountDigits":"123","balance":"*****","totalBalance":200.17}}
    [WARNING 2023/12/11 06:21:09] main.go:22: test mask/hide struct: {"name":"Foo Bar","birthDate":"1999-01-21T00:00:00-02:00","document":"*****996642","emails":["gabriel@test.com","gabrielcataldo.231@gmail.com","biel@test.org"],"balances":["10***","1*","13***","1239****","23***"],"bank":{"account":"123981023","accountDigits":"123","balance":"*****","totalBalance":200.17}}
    [ERROR 2023/12/11 06:21:09] main.go:23: test mask/hide struct: {"name":"Foo Bar","birthDate":"1999-01-21T00:00:00-02:00","document":"*****996642","emails":["gabriel@test.com","gabrielcataldo.231@gmail.com","biel@test.org"],"balances":["10***","1*","13***","1239****","23***"],"bank":{"account":"123981023","accountDigits":"123","balance":"*****","totalBalance":200.17}}

You can also use functions with the ending **H** to hide all values, **MS** to mask half
initial half of all values, and **ME** to mask the trailing half of all values. Look:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    // also applies to other levels
    msg := getBasicMsg()
    logger.Info(msg...)
    logger.InfoH(msg...)
    logger.InfoMS(msg...)
    logger.InfoME(msg...)
}

func getBasicMsg() []any {
    s := []any{
        "text string",
        1,
        12.213,
        true,
        nil,
        time.Now(),
    }
    m := map[string]any{
        "int":      1,
        "float":    12.213,
        "string":   "text string",
        "bool":     true,
        "nilValue": nil,
        "time":     time.Now(),
    }
    return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}
```

Outputs:

    [INFO 2023/12/11 07:19:09] main.go:5: basic example with any values text string 1 12.213 true 2023-12-11T07:19:42-03:00 {"string":"text string","bool":true,"time":"2023-12-11T07:19:42-03:00","nilValue":null,"int":1,"float":12.213} ["text string",1,12.213,true,null,"2023-12-11T07:19:42-03:00"]
    [INFO 2023/12/11 07:19:09] main.go:6: ***************************** *********** * ****** **** ************************* {"bool":"****","time":"*************************","nilValue":null,"int":"*","float":"******","string":"***********"} ["***********","*","******","****",null,"*************************"]
    [INFO 2023/12/11 07:19:09] main.go:7: **************with any values *****string * ***213 **ue ************7:19:42-03:00 {"bool":"**ue","time":"************7:19:42-03:00","nilValue":null,"int":"*","float":"***213","string":"*****string"} ["*****string","*","***213","**ue",null,"************7:19:42-03:00"]
    [INFO 2023/12/11 07:19:09] main.go:8: basic example *************** text ****** * 12.*** tr** 2023-12-11T0************* {"time":"2023-12-11T0*************","nilValue":null,"int":"*","float":"12.***","string":"text ******","bool":"tr**"} ["text ******","*","12.***","tr**",null,"2023-12-11T0*************"]

For total customization, you can pass opts as a parameter in functions with default **Opts**
as already mentioned above, and you can also pass it as a global variable by calling the **SetOptions** function
and to reset the global options to default, you can call the function **ResetOptionsToDefault**
see the example below:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    logger.SetOptions(getCustomOptionsExample())
    msg := getBasicMsg()
    format := "%v, %v, %v, %v, %v, %v, %v, last is %v"
    logger.Info(msg...)
    logger.ResetOptionsToDefault()
    logger.Debug(msg...)
    logger.WarningSkipCallerOptsf(format, 1, *getCustomOptionsExample(), msg...)
    logger.ErrorOptsf(format, *getCustomOptionsExample(), msg...)
}

func getCustomOptionsExample() *logger.Options {
    return &logger.Options{
        // Print mode (default: ModeDefault)
        Mode:                  logger.RandomMode(),
        // Argument date format (default: DateFormatFull24h)
        DateFormat:            logger.RandomDateFormat(),
        // Custom prefix text
        CustomPrefixText: "",
        // Custom after prefix text (only if Mode is ModeDefault)
        CustomAfterPrefixText: "",
        // Enable argument date to be UTC (default: false)
        UTC:                   util.RandomBool(),
        // Enable to not print empty message (default: false)
        DontPrintEmptyMessage: util.RandomBool(),
        // Enable to remove spaces between parameters (default: false)
        RemoveSpace:           util.RandomBool(),
        // If true will hide all datetime and prefix arguments (default: false)
        HideAllArgs:           util.RandomBool(),
        // If true it will hide the datetime arguments (default: false)
        HideArgDatetime:       util.RandomBool(),
        // If true, it will hide the caller arguments (default: false)
        HideArgCaller:         util.RandomBool(),
        // If true, it will disable all argument and prefix colors (default: false)
        DisablePrefixColors:   util.RandomBool(),
        // If true, json mode msg field becomes slice (default: false, only if Mode is ModeJson)
        //
        // IMPORTANT: If true, the format parameter will not work
        EnableJsonMsgFieldForSlice: false,
    }
}

func getBasicMsg() []any {
    s := []any{
        "text string",
        1,
        12.213,
        true,
        nil,
        time.Now(),
    }
    m := map[string]any{
        "int":      1,
        "float":    12.213,
        "string":   "text string",
        "bool":     true,
        "nilValue": nil,
        "time":     time.Now(),
    }
    return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}
```

Outputs:

    INFO: basic example with any valuestext string1 12.213 true2023-12-11T08:59:08-03:00{"time":"2023-12-11T08:59:08-03:00","nilValue":null,"int":1,"float":12.213,"string":"text string","bool":true}["text string",1,12.213,true,null,"2023-12-11T08:59:08-03:00"]
    [DEBUG 2023/12/11 08:59:09] types.go:194: basic example with any values text string 1 12.213 true 2023-12-11T08:59:08-03:00 {"bool":true,"time":"2023-12-11T08:59:08-03:00","nilValue":null,"int":1,"float":12.213,"string":"text string"} ["text string",1,12.213,true,null,"2023-12-11T08:59:08-03:00"]
    {"level":"WARNING","msg":"basic example with any values, text string, 1, 12.213, true, 2023-12-11T08:59:08-03:00, {\"nilValue\":null,\"int\":1,\"float\":12.213,\"string\":\"text string\",\"bool\":true,\"time\":\"2023-12-11T08:59:08-03:00\"}, last is [\"text string\",1,12.213,true,null,\"2023-12-11T08:59:08-03:00\"]"}
    ERROR: basic example with any values, text string, 1, 12.213, true, 2023-12-11T08:59:08-03:00, {"nilValue":null,"int":1,"float":12.213,"string":"text string","bool":true,"time":"2023-12-11T08:59:08-03:00"}, last is ["text string",1,12.213,true,null,"2023-12-11T08:59:08-03:00"]

Finally, if you want to print messages asynchronously, you can configure your global options
or via parameter filling the **EnableAsynchronousMode** field as **true**, remembering that the information,
file (ex: main.go:23) are not printed in the asynchronous message, see:

```go
import "github.com/GabrielHCataldo/go-logger/logger"

func main() {
    asyncOpt := logger.Options{EnableAsynchronousMode: true}
    msg := getBasicMsg()
    logger.SetOptions(&asyncOpt)
    logger.Info(msg...)
    logger.ResetOptionsToDefault()
    logger.InfoOpts(asyncOpt)
}

func getBasicMsg() []any {
    s := []any{
        "text string",
        1,
        12.213,
        true,
        nil,
        time.Now(),
    }
    m := map[string]any{
        "int":      1,
        "float":    12.213,
        "string":   "text string",
        "bool":     true,
        "nilValue": nil,
        "time":     time.Now(),
    }
    return []any{"basic example with any values", "text string", 1, 12.213, true, time.Now(), m, s}
}
```

#### IMPORTANT
In our go-logger lib we have several calls with all the examples above, always following the standards of
nomenclature, let's explain it better:

```go
logger.InfoSkipCallerOptsfH("format", 1, opts, v...)
```

All our calls have this priority layout by default, eliminating them according to your desire in 
At the moment, the only mandatory name is the level, see:

    logger.{level}{SkipCaller}{Opts}{f}-{H/MS/ME}

Example eliminating just SkipCaller:

```go
logger.InfoOptsfH("format", opts, v...)
```

Example eliminating just Opts:

```go
logger.InfoSkipCallerfH("format", 1, v...)
```

Example eliminating only the f (format):

```go
logger.InfoSkipCallerOptsH(1, opts, v...)
```

Example eliminating only H/MS/ME (masks):

```go
logger.InfoSkipCallerOpts("format", 1, opts, v...)
```

How to contribute
------
Make a pull request, or if you find a bug, open it
an Issues.

License
-------
Distributed under MIT license, see the license file within the code for more details.