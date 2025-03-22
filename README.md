![CI Status](https://github.com/timematic/anytime/actions/workflows/test.yml/badge.svg)

# anytime

Fast, Zero-Allocs and User-Friendly `time.Time` parser which no need specify the time `Layout`.

## Example

```
package main

import (
	"fmt"
	"github.com/timematic/anytime"
)

func main() {
	datetime, err := anytime.Parse("2006-01-02T15:04:05")
	if err != nil {
		panic(err)
	}
	fmt.Println(datetime) // 2006-01-02 15:04:05 +0000 UTC

	loc, _ := time.LoadLocation("America/New_York")
	datetime, err = anytime.ParseInLocation("2006-01-02T15:04:05", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(datetime) // 2006-01-02 15:04:05 -0500 EST
}
```

## Benchmark

goos: darwin
goarch: arm64
pkg: github.com/timematic/anytime
cpu: Apple M1 Pro

|                                      | ns/op       | allocation bytes | allocation times |
| ------------------------------------ | ----------- | ---------------- | ---------------- |
| time.Parse date_only                 | 56.82 ns/op | 0 B/op           | 0 allocs/op      |
| anytime.Parse date_only              | 34.11 ns/op | 0 B/op           | 0 allocs/op      |
| araddon/dateparse.ParseAny date_only | 210.3 ns/op | 288 B/op         | 3 allocs/op      |
| time.Parse rfc3339                   | 38.26 ns/op | 0 B/op           | 0 allocs/op      |
| anytime.Parse rfc3339                | 67.56 ns/op | 0 B/op           | 0 allocs/op      |
| araddon/dateparse.ParseAny rfc3339   | 304.7 ns/op | 320 B/op         | 3 allocs/op      |

## Support lots of time layout

If you need support other layouts, feel free to share Issues or PRs, both are welcomed.

```
var anytime_layouts = map[string]string{ // map[value]layout
	// Multiple Date Layout
	"1970-01-01":      "2006-01-02",
	"1970-Jan-01":     "2006-Jan-02",
	"1970/01/01":      "2006/01/02",
	"1970/Jan/01":     "2006/Jan/02",
	"1970/jan/01":     "2006/Jan/02",
	"1970/January/01": "2006/January/02",
	"01/Jan/1970":     "02/Jan/2006",
	"Jan/01/1970":     "Jan/02/2006",
	"19700101":        "20060102",
	"1970.01.01":      "2006.01.02",
	"1970-01-01T":     "2006-01-02T",
	"1970.001":        "2006.002",
	"1970-001":        "2006-002",
	"1970001":         "2006002",
	"1970 01 23":      "2006 01 02",
	"1970 Jan 23":     "2006 Jan 02",
	"Jan 23 1970":     "Jan 02 2006",
	"23 Jan 1970":     "02 Jan 2006",
	"Mon 30 Sep 2018": "Mon 02 Jan 2006",
	"Fri Jul 03 2015": "Mon Jan 02 2006",
	"2014年04月08日":     "2006年01月02日",
	"2014年4月8日":       "2006年1月2日",
	"1970-05-13":      "2006-01-02",
	"05-13-1970":      "01-02-2006",
	"13-05-1970":      "02-01-2006",
	"05-05-1970":      "02-01-2006",

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

	// ignore zone_abbrvs if contains zone_offset
	"1970-01-01 PST-0300":                     "2006-01-02 PST-0700",
	"1970-01-01 UTC-0300":                     "2006-01-02 UTC-0700",
	"1970-01-01 GMT-0300":                     "2006-01-02 GMT-0700",
	"1970-01-01 -0300 UTC":                    "2006-01-02 -0700 UTC",
	"1970-01-01 -0300 GMT":                    "2006-01-02 -0700 GMT",
	"1970-01-01 -0300 PST":                    "2006-01-02 -0700 PST",
	"1970-01-01 PST-0000":                     "2006-01-02 PST-0700",
	"1970-01-01 UTC-0000":                     "2006-01-02 UTC-0700",
	"1970-01-01 GMT-0000":                     "2006-01-02 GMT-0700",
	"1970-01-01 -0000 UTC":                    "2006-01-02 -0700 UTC",
	"1970-01-01 -0000 GMT":                    "2006-01-02 -0700 GMT",
	"1970-01-01 -0000 PST":                    "2006-01-02 -0700 PST",
	"1970-01-01 PST+0300":                     "2006-01-02 PST-0700",
	"1970-01-01 UTC+0300":                     "2006-01-02 UTC-0700",
	"1970-01-01 GMT+0300":                     "2006-01-02 GMT-0700",
	"1970-01-01 GMT+03":                       "2006-01-02 GMT-07",
	"1970-01-01 UTC+03":                       "2006-01-02 UTC-07",
	"1970-01-01 +0300 UTC":                    "2006-01-02 -0700 UTC",
	"1970-01-01 +0300 GMT":                    "2006-01-02 -0700 GMT",
	"1970-01-01 +0300 PST":                    "2006-01-02 -0700 PST",
	"1970-01-01 PST+0000":                     "2006-01-02 PST-0700",
	"1970-01-01 UTC+0000":                     "2006-01-02 UTC-0700",
	"1970-01-01 GMT+0000":                     "2006-01-02 GMT-0700",
	"1970-01-01 +0000 UTC":                    "2006-01-02 -0700 UTC",
	"1970-01-01 +0000 GMT":                    "2006-01-02 -0700 GMT",
	"1970-01-01 +0000 PST":                    "2006-01-02 -0700 PST",
	"1970-01-01 GMT+0100 (GMT Daylight Time)": "2006-01-02 GMT-0700 (GMT Daylight Time)",

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
	"1970-01-01 00:00:00,123456789":   "2006-01-02 15:04:05,000000000",
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
	"1970-01-01T00":          "2006-01-02T15",
	"1970-01-01t00":          "2006-01-02t15",
	"1970-01-01_00":          "2006-01-02_15",
	"1970-01-01 at 13:04:05": "2006-01-02 at 15:04:05",
	"07 Feb 2004, 09:07:07":  "02 Jan 2006, 15:04:05",

	// Some Special Layout
	"Wed Dec  3 07:37:16 1997":               "Mon Jan _2 15:04:05 2006",                // time.ANSIC
	"Wed Dec  3 07:37:16 PST 1997":           "Mon Jan _2 15:04:05 MST 2006",            // time.UnixDate
	"2006-01-02 15:04:05 +0800 CST":          "2006-01-02 15:04:05.999999999 -0700 MST", // time.String
	"02 Jan 70 00:01 UTC":                    "02 Jan 06 15:04 MST",
	"Mon, 02-Jan-70 00:01:02 UTC":            "Mon, 02-Jan-06 15:04:05 MST",
	"Mon Jan 02 00:01:02 -0500 1970":         "Mon Jan 02 15:04:05 -0700 2006",
	"Wed Dec 17 07:37:16 1997 PST":           "Mon Jan 02 15:04:05 2006 MST", // pg style
	"Dec 17 07:37:16 1997 PST":               "Jan 02 15:04:05 2006 MST",     // pg style
	"19700102T030405.123456":                 "20060102T150405.000000",
	"January 02, 2006, 15:04:05":             "January 02, 2006, 15:04:05",
	"January 02, 2006 15:04:05":              "January 02, 2006 15:04:05",
	"January 02, 2006 03:04:05PM":            "January 02, 2006 03:04:05PM",
	"September 17, 2012 at 5:00PM UTC":       "January 02, 2006 at 3:04PM MST",
	"Tue, 11 Jul 2017 04:08:03 +0200 (CEST)": "Mon, 02 Jan 2006 15:04:05 -0700 (CEST)",
	"Thu May 08 17:57:51 CEST 2009":          "Mon Jan 02 15:04:05 MST 2006",
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

## Alternatives:

- [github.com/araddon/dateparse](https://github.com/araddon/dateparse) `dateparse` support little more formats, while `anytime` does not support.
  - `anytime` is ~4x faster than `dateparse`
