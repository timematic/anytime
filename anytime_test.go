package anytime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/longqimin/anytime"

	"github.com/stretchr/testify/assert"
)

var golang_time_formats = []string{
	// time.Layout,
	// time.ANSIC,
	// time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.DateTime,
	time.DateOnly,
}

var date_fmts = []string{
	"2006-01-02", "2006-1-02", "2006-01-2", "2006-1-2", "2006-Jan-02", "2006-Jan-2", // yyyy-mm-dd
	"2006/01/02", "2006/1/02", "2006/01/2", "2006/1/2", "2006/Jan/02", "2006/Jan/2", // yyyy/mm/dd
	"20060102",    // yyyymmdd
	"2006.002",    // yyyy.day_of_year
	"2006-002",    // yyyy.day_of_year
	"2006002",     // yyyy.day_of_year
	"Jan/02/2006", // mm/dd/yyy
	"Jan-02-2006", // mm-dd-yyy
}

var time_fmts = []string{
	"15", "1504", "150405", "150405.0", "150405.00", "150405.000", "150405.0000", "150405.00000", "150405.000000",
	"15:04", "15:04:05", "15:04:05.0", "15:04:05.00", "15:04:05.000", "15:04:05.0000", "15:04:05.00000", "15:04:05.000000",
	"03PM", "03:04PM", "03:04:05PM", "03:04:05.0PM", "03:04:05.00PM", "03:04:05.000PM", "03:04:05.0000PM", "03:04:05.00000PM", "03:04:05.000000PM",
}
var zone_fmts = []string{
	"", "-07", "-0700", "-07:00", "Z07:00", "MST",
}

func gen_date_only_formats() []string {
	fmts := date_fmts

	for _, date_fmt := range date_fmts {
		for _, zone_fmt := range zone_fmts {
			fmts = append(fmts, fmt.Sprintf("%s%s", date_fmt, zone_fmt))
			fmts = append(fmts, fmt.Sprintf("%s %s", date_fmt, zone_fmt))
		}
	}
	return fmts
}

func gen_datetime_formats() []string {
	fmts := gen_date_only_formats()

	for _, date_fmt := range date_fmts {
		for _, time_fmt := range time_fmts {
			for _, sp := range []string{" ", "T"} {
				datetime_fmt := fmt.Sprintf("%s%s%s", date_fmt, sp, time_fmt)
				fmts = append(fmts, datetime_fmt)
				for _, zone_fmt := range zone_fmts {
					fmts = append(fmts, fmt.Sprintf("%s %s", datetime_fmt, zone_fmt))
				}
			}
		}
	}

	for _, layout := range anytime_layouts {
		if layout != "" {
			fmts = append(fmts, layout)
		}
	}
	return append(fmts, golang_time_formats...)
}

func checkParse(t *testing.T, layout, value string) {
	expect, err := time.Parse(layout, value)
	assert.Nil(t, err)

	date, err := anytime.Parse(value)
	assert.Nil(t, err)
	assert.True(t, expect.Equal(date), fmt.Sprintf("layout=%s, value=%s,  expect=%v, output=%v", layout, value, expect, date))
}

func checkParseInLocation(t *testing.T, layout, value, location string) {
	loc := time.Local
	var err error
	if location != "Local" {
		if loc, err = time.LoadLocation(location); err != nil {
			return
		}
	}

	expect, err := time.ParseInLocation(layout, value, loc)
	assert.Nil(t, err)

	date, err := anytime.ParseInLocation(value, loc)
	assert.Nil(t, err)
	assert.True(t, expect.Equal(date), fmt.Sprintf("layout=%s, value=%s, loc=%s, expect=%v, output=%v", layout, value, location, expect, date))
}

func TestIssues(t *testing.T) {
	// issues
	checkParseInLocation(t, "2006.002", "2088.001", "Africa/Casablanca")
	checkParseInLocation(t, "2006002", "2025063", "UTC")
	checkParseInLocation(t, "2006/January/02", "1970/June/29", "UTC")
	checkParseInLocation(t, "Mon Jan 02 15:04:05 2006 MST", "Wed Dec 17 07:37:16 1997 PST", "UTC")
	checkParseInLocation(t, "Jan 02 15:04:05 2006 MST", "Dec 17 07:37:16 1997 PST", "UTC")

}

func TestParse(t *testing.T) {
	now := time.Now()
	for _, layout := range gen_datetime_formats() {
		str := now.Format(layout)
		checkParse(t, layout, str)
	}
}

func TestParseInLocation(t *testing.T) {
	now := time.Now()

	for _, layout := range gen_datetime_formats() {
		str := now.Format(layout)

		checkParseInLocation(t, layout, str, "Local")
		checkParseInLocation(t, layout, str, "UTC")
		checkParseInLocation(t, layout, str, "Asia/Shanghai")
		checkParseInLocation(t, layout, str, "America/New_York")

	}

	layout := "2006-01-02 15:04:05"
	str := now.Format(layout)
	for _, location := range append(iana_timezones, iana_timezone_abbrvs...) {
		checkParseInLocation(t, layout, str, location)
	}

	layout = "2006-01-02 15:04:05 Z0700"
	str = now.Format(layout)
	for _, location := range append(iana_timezones, iana_timezone_abbrvs...) {
		checkParseInLocation(t, layout, str, location)
	}

	layout = "2006-01-02 15:04:05 MST"
	str = now.Format(layout)
	for _, location := range append(iana_timezones, iana_timezone_abbrvs...) {
		checkParseInLocation(t, layout, str, location)
	}
}
