![CI Status](https://github.com/longqimin/anytime/actions/workflows/test.yml/badge.svg)

# anytime

a user friendly `time.Time` parser which no need specify the time `Layout`.

## Support lots of time layout

```
var anytime_layouts = map[string]string{ // map[value]layout
	// Multiple Date Layout
	"1970-01-01":      "2006-01-02",
	"1970-Jan-01":     "2006-Jan-02",
	"1970/01/01":      "2006/01/02",
	"1970/Jan/01":     "2006/Jan/02",
	"1970/January/01": "2006/January/02",
	"19700101":        "20060102",
	"Jan/01/1970":     "Jan/02/2006",
	"1970-01-01T":     "2006-01-02T",
	"1970.001":        "2006.002",
	"1970-001":        "2006-002",
	"1970001":         "2006002",

	// Multiple Zone Layout
	"1970-01-01Z":                 "2006-01-02Z07:00",
	"1970-01-01+05":               "2006-01-02-07",
	"1970-01-01T+05":              "2006-01-02T-07",
	"1970-01-01 +05":              "2006-01-02 -07",
	"1970-01-01 +05:00":           "2006-01-02 -07:00",
	"1970-01-01 +0500":            "2006-01-02 -0700",
	"1970-01-01 -05":              "2006-01-02 -07",
	"1970-01-01 -05:00":           "2006-01-02 -07:00",
	"1970-01-01 -0500":            "2006-01-02 -0700",
	"1970-01-01 UTC":              "2006-01-02 MST",
	"1970-01-01 America/New_York": "", // not supported by time.Parse
	"1970-01-01 +5":               "", // not supported by time.Parse
	"1970-01-01 +5:0":             "", // not supported by time.Parse
	"1970-01-01 -8:00":            "", // not supported by time.Parse

	// Multiple Time Layout
	"1970-01-01 00":                   "2006-01-02 15",
	"1970-01-01 00:00":                "2006-01-02 15:04",
	"1970-01-01 00:00:00":             "2006-01-02 15:04:05",
	"1970-01-01 00:00:00.1":           "2006-01-02 15:04:05.0",
	"1970-01-01 00:00:00.12":          "2006-01-02 15:04:05.00",
	"1970-01-01 00:00:00.123":         "2006-01-02 15:04:05.000",
	"1970-01-01 00:00:00.1234":        "2006-01-02 15:04:05.0000",
	"1970-01-01 00:00:00.12345":       "2006-01-02 15:04:05.00000",
	"1970-01-01 00:00:00.123456":      "2006-01-02 15:04:05.000000",
	"1970-01-01 00:00:00.1234567":     "2006-01-02 15:04:05.0000000",
	"1970-01-01 00:00:00.12345678":    "2006-01-02 15:04:05.00000000",
	"1970-01-01 00:00:00.123456789":   "2006-01-02 15:04:05.000000000",
	"1970-01-01 0":                    "2006-01-02 15",
	"1970-01-01 0:1":                  "2006-01-02 15:4",
	"1970-01-01 0:1:2":                "2006-01-02 15:4:5",
	"1970-01-01 0:1:02":               "2006-01-02 15:4:05",
	"1970-01-01 0001":                 "2006-01-02 1504",
	"1970-01-01 000102":               "2006-01-02 150405",
	"1970-01-01 000102.1":             "2006-01-02 150405.0",
	"1970-01-01 10:00:00.123456789AM": "2006-01-02 03:04:05.000000000PM",
	"1970-01-01 10:00:00.123456789PM": "2006-01-02 03:04:05.000000000PM",

	// Multiple Date/Time Separator
	"1970-01-01T00": "2006-01-02T15",
	"1970-01-01t00": "2006-01-02t15",
	"1970-01-01_00": "2006-01-02_15",

	// Some Special Layout
	"02 Jan 70 00:01 UTC":            "02 Jan 06 15:04 MST",
	"Mon, 02-Jan-70 00:01:02 UTC":    "Mon, 02-Jan-06 15:04:05 MST",
	"Mon Jan 02 00:01:02 -0500 1970": "Mon Jan 02 15:04:05 -0700 2006",
	"Wed Dec 17 07:37:16 1997 PST":   "Mon Jan 02 15:04:05 2006 MST", // pg style
	"Dec 17 07:37:16 1997 PST":       "Jan 02 15:04:05 2006 MST",     // pg style
	"19700102T030405.123456":         "20060102T150405.000000",
}
```

## Example

```
package main

import (
	"fmt"

	"github.com/longqimin/anytime"
)

func main() {
	datetime, err := anytime.Parse("2006-01-02T15:04:05+08")
	if err != nil {
		panic(err)
	}
	fmt.Println(datetime) // 2006-01-02 15:04:05 +0800 CST
}
```

## APIs

### `anytime.Parse(value string) (time.Time, error)`

automatically figure out the time `Layout` and parse to `time.Time`.
**possible implementation**:

```go
Parse(value string) (time.Time, error){
    layout, err := ExtractLayout(str)
    if err != nil {
        return time.Time{}, err
    }
    return time.Parse(layout, value)
}
```

### `anytime.ParseInLocation(value string, loc *time.Location) (time.Time, error)`

**possible implementation**:

```go
ParseInLocation(value string, loc *time.Location) (time.Time, error){
    layout, err := ExtractLayout(str)
    if err != nil {
        return time.Time{}, err
    }
    return time.ParseInLocation(layout, value, loc)
}
```

## Benchmark

|                         | ns/op       |
| ----------------------- | ----------- |
| time.Parse date_only    | 38.98 ns/op |
| anytime.Parse date_only | 28.14 ns/op |
| time.Parse rfc3339      | 26.79 ns/op |
| anytime.Parse rfc3339   | 52.53 ns/op |
