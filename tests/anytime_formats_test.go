package anytime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/timematic/anytime"
)

type testdata struct {
	input  string
	layout string
	expect string
}

var anytime_tests = []testdata{
	// Multiple Date Layout
	{input: "1970-01-01", layout: "2006-01-02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-Jan-01", layout: "2006-Jan-02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970/01/01", layout: "2006/01/02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970/Jan/01", layout: "2006/Jan/02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970/jan/01", layout: "2006/Jan/02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970/January/01", layout: "2006/January/02", expect: "1970-01-01T00:00:00Z"},
	{input: "01/Jan/1970", layout: "02/Jan/2006", expect: "1970-01-01T00:00:00Z"},
	{input: "Jan/01/1970", layout: "Jan/02/2006", expect: "1970-01-01T00:00:00Z"},
	{input: "8/8/1964", layout: "1/2/2006", expect: "1964-08-08T00:00:00Z"},
	{input: "8/08/1964", layout: "1/02/2006", expect: "1964-08-08T00:00:00Z"},
	{input: "08/8/1964", layout: "01/2/2006", expect: "1964-08-08T00:00:00Z"},
	{input: "08/08/1964", layout: "01/02/2006", expect: "1964-08-08T00:00:00Z"},
	{input: "19700101", layout: "20060102", expect: "1970-01-01T00:00:00Z"},
	{input: "1970.01.01", layout: "2006.01.02", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01T", layout: "2006-01-02T", expect: "1970-01-01T00:00:00Z"},
	{input: "1970.001", layout: "2006.002", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-001", layout: "2006-002", expect: "1970-01-01T00:00:00Z"},
	{input: "1970005", layout: "2006002", expect: "1970-01-05T00:00:00Z"},
	{input: "1970 01 23", layout: "2006 01 02", expect: "1970-01-23T00:00:00Z"},
	{input: "1970 Jan 23", layout: "2006 Jan 02", expect: "1970-01-23T00:00:00Z"},
	{input: "Jan 23 1970", layout: "Jan 02 2006", expect: "1970-01-23T00:00:00Z"},
	{input: "23 Jan 1970", layout: "02 Jan 2006", expect: "1970-01-23T00:00:00Z"},
	{input: "23 Jan '70", layout: "02 Jan '06", expect: "1970-01-23T00:00:00Z"},
	{input: "Mon 30 Sep 2018", layout: "Mon 02 Jan 2006", expect: "2018-09-30T00:00:00Z"},
	{input: "Fri Jul 03 2015", layout: "Mon Jan 02 2006", expect: "2015-07-03T00:00:00Z"},
	{input: "2014年04月08日", layout: "2006年01月02日", expect: "2014-04-08T00:00:00Z"},
	{input: "2014年4月8日", layout: "2006年1月2日", expect: "2014-04-08T00:00:00Z"},
	{input: "1970-05-13", layout: "2006-01-02", expect: "1970-05-13T00:00:00Z"},
	{input: "05-13-1970", layout: "01-02-2006", expect: "1970-05-13T00:00:00Z"},
	{input: "13-05-1970", layout: "02-01-2006", expect: "1970-05-13T00:00:00Z"},
	{input: "05-05-1970", layout: "02-01-2006", expect: "1970-05-05T00:00:00Z"},
	{input: "May 1st 2012", layout: "Jan 2st 2006", expect: "2012-05-01T00:00:00Z"},
	{input: "2004", layout: "2006", expect: "2004-01-01T00:00:00Z"},
	{input: "2004.05", layout: "2006.01", expect: "2004-05-01T00:00:00Z"},
	{input: "2004/05", layout: "2006/01", expect: "2004-05-01T00:00:00Z"},
	{input: "2004-05", layout: "2006-01", expect: "2004-05-01T00:00:00Z"},

	// Multiple Zone Layout
	{input: "1970-01-01Z", layout: "2006-01-02Z07:00", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01+05", layout: "2006-01-02-07", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01T+05", layout: "2006-01-02T-07", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 +05", layout: "2006-01-02 -07", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 +05:00", layout: "2006-01-02 -07:00", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 +0500", layout: "2006-01-02 -0700", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 -05", layout: "2006-01-02 -07", expect: "1970-01-01T05:00:00Z"},
	{input: "1970-01-01 -05:00", layout: "2006-01-02 -07:00", expect: "1970-01-01T05:00:00Z"},
	{input: "1970-01-01 -0500", layout: "2006-01-02 -0700", expect: "1970-01-01T05:00:00Z"},
	{input: "1970-01-01 UTC", layout: "2006-01-02 MST", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 America/New_York", layout: "", expect: "1970-01-01T00:00:00-05:00"},
	{input: "1970-01-01 +5", layout: "", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 +5:0", layout: "", expect: "1969-12-31T19:00:00Z"},
	{input: "1970-01-01 -8:00", layout: "", expect: "1970-01-01T08:00:00Z"},
	{input: "2018-06-29 +0300 +03", layout: "2006-01-02 -0700 -07", expect: "2018-06-28T21:00:00Z"},

	// ignore zone_abbrvs if contains zone_offset
	{input: "1970-01-01 PST-0300", layout: "2006-01-02 PST-0700", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 UTC-0300", layout: "2006-01-02 UTC-0700", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 GMT-0300", layout: "2006-01-02 GMT-0700", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 -0300 UTC", layout: "2006-01-02 -0700 UTC", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 -0300 GMT", layout: "2006-01-02 -0700 GMT", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 -0300 PST", layout: "2006-01-02 -0700 PST", expect: "1970-01-01T03:00:00Z"},
	{input: "1970-01-01 PST-0000", layout: "2006-01-02 PST-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 UTC-0000", layout: "2006-01-02 UTC-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 GMT-0000", layout: "2006-01-02 GMT-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 -0000 UTC", layout: "2006-01-02 -0700 UTC", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 -0000 GMT", layout: "2006-01-02 -0700 GMT", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 -0000 PST", layout: "2006-01-02 -0700 PST", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 PST+0300", layout: "2006-01-02 PST-0700", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 UTC+0300", layout: "2006-01-02 UTC-0700", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 GMT+0300", layout: "2006-01-02 GMT-0700", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 GMT+03", layout: "2006-01-02 GMT-07", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 UTC+03", layout: "2006-01-02 UTC-07", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 +0300 UTC", layout: "2006-01-02 -0700 UTC", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 +0300 GMT", layout: "2006-01-02 -0700 GMT", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 +0300 PST", layout: "2006-01-02 -0700 PST", expect: "1969-12-31T21:00:00Z"},
	{input: "1970-01-01 PST+0000", layout: "2006-01-02 PST-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 UTC+0000", layout: "2006-01-02 UTC-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 GMT+0000", layout: "2006-01-02 GMT-0700", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 +0000 UTC", layout: "2006-01-02 -0700 UTC", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 +0000 GMT", layout: "2006-01-02 -0700 GMT", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 +0000 PST", layout: "2006-01-02 -0700 PST", expect: "1970-01-01T00:00:00Z"},
	{input: "1970-01-01 GMT+0100 (GMT Daylight Time)", layout: "2006-01-02 GMT-0700 (GMT Daylight Time)", expect: "1969-12-31T23:00:00Z"},

	// Multiple Time Layout
	{input: "1970-01-01 01", layout: "2006-01-02 15", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01 01:00", layout: "2006-01-02 15:04", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01 01:02:03", layout: "2006-01-02 15:04:05", expect: "1970-01-01T01:02:03Z"},
	{input: "1970-01-01 01:02:03.1", layout: "2006-01-02 15:04:05.0", expect: "1970-01-01T01:02:03.1Z"},
	{input: "1970-01-01 01:02:03.12", layout: "2006-01-02 15:04:05.00", expect: "1970-01-01T01:02:03.12Z"},
	{input: "1970-01-01 01:02:03.123", layout: "2006-01-02 15:04:05.000", expect: "1970-01-01T01:02:03.123Z"},
	{input: "1970-01-01 01:02:03.1234", layout: "2006-01-02 15:04:05.0000", expect: "1970-01-01T01:02:03.1234Z"},
	{input: "1970-01-01 01:02:03.12345", layout: "2006-01-02 15:04:05.00000", expect: "1970-01-01T01:02:03.12345Z"},
	{input: "1970-01-01 01:02:03.123456", layout: "2006-01-02 15:04:05.000000", expect: "1970-01-01T01:02:03.123456Z"},
	{input: "1970-01-01 01:02:03.1234567", layout: "2006-01-02 15:04:05.0000000", expect: "1970-01-01T01:02:03.1234567Z"},
	{input: "1970-01-01 01:02:03.12345678", layout: "2006-01-02 15:04:05.00000000", expect: "1970-01-01T01:02:03.12345678Z"},
	{input: "1970-01-01 01:02:03.123456789", layout: "2006-01-02 15:04:05.000000000", expect: "1970-01-01T01:02:03.123456789Z"},
	{input: "1970-01-01 01:02:03,123456789", layout: "2006-01-02 15:04:05,000000000", expect: "1970-01-01T01:02:03.123456789Z"},
	{input: "1970-01-01 1", layout: "2006-01-02 15", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01 1:2", layout: "2006-01-02 15:4", expect: "1970-01-01T01:02:00Z"},
	{input: "1970-01-01 1:2:3", layout: "2006-01-02 15:4:5", expect: "1970-01-01T01:02:03Z"},
	{input: "1970-01-01 1:2:03", layout: "2006-01-02 15:4:05", expect: "1970-01-01T01:02:03Z"},
	{input: "1970-01-01 0102", layout: "2006-01-02 1504", expect: "1970-01-01T01:02:00Z"},
	{input: "1970-01-01 010203", layout: "2006-01-02 150405", expect: "1970-01-01T01:02:03Z"},
	{input: "1970-01-01 010203.1", layout: "2006-01-02 150405.0", expect: "1970-01-01T01:02:03.1Z"},
	{input: "1970-01-01 10:00:00.123456789AM", layout: "2006-01-02 03:04:05.000000000PM", expect: "1970-01-01T10:00:00.123456789Z"},
	{input: "1970-01-01 10:00:00.123456789PM", layout: "2006-01-02 03:04:05.000000000PM", expect: "1970-01-01T22:00:00.123456789Z"},

	// Multiple Date/Time Separator
	{input: "1970-01-01T01", layout: "2006-01-02T15", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01t01", layout: "2006-01-02t15", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01_01", layout: "2006-01-02_15", expect: "1970-01-01T01:00:00Z"},
	{input: "1970-01-01 at 13:04:05", layout: "2006-01-02 at 15:04:05", expect: "1970-01-01T13:04:05Z"},
	{input: "07 Feb 2004, 09:07:07", layout: "02 Jan 2006, 15:04:05", expect: "2004-02-07T09:07:07Z"},

	// Some Special Layout
	{input: "01/20 03:04:05PM '06 -0700", layout: "01/02 03:04:05PM '06 -0700", expect: "2006-01-20T22:04:05Z"},                 // time.Layout
	{input: "Wed Dec  3 07:37:16 1997", layout: "Mon Jan _2 15:04:05 2006", expect: "1997-12-03T07:37:16Z"},                     // time.ANSIC
	{input: "Wed Dec  3 07:37:16 PST 1997", layout: "Mon Jan _2 15:04:05 MST 2006", expect: "1997-12-03T07:37:16Z"},             // time.UnixDate
	{input: "2006-01-02 15:04:05 +0800 CST", layout: "2006-01-02 15:04:05.999999999 -0700 MST", expect: "2006-01-02T07:04:05Z"}, // time.String
	{input: "02 Jan 70 00:01 UTC", layout: "02 Jan 06 15:04 MST", expect: "1970-01-02T00:01:00Z"},
	{input: "Mon, 02-Jan-70 00:01:02 UTC", layout: "Mon, 02-Jan-06 15:04:05 MST", expect: "1970-01-02T00:01:02Z"},
	{input: "Mon Jan 02 00:01:02 -0500 1970", layout: "Mon Jan 02 15:04:05 -0700 2006", expect: "1970-01-02T05:01:02Z"},
	{input: "Wed Dec 17 07:37:16 1997 PST", layout: "Mon Jan 02 15:04:05 2006 MST", expect: "1997-12-17T07:37:16Z"}, // pg style
	{input: "Dec 17 07:37:16 1997 PST", layout: "Jan 02 15:04:05 2006 MST", expect: "1997-12-17T07:37:16Z"},         // pg style
	{input: "Dec 17 07:37:16 '97 PST", layout: "Jan 02 15:04:05 '06 MST", expect: "1997-12-17T07:37:16Z"},           // pg style
	{input: "19700102T030405.123456", layout: "20060102T150405.000000", expect: "1970-01-02T03:04:05.123456Z"},
	{input: "January 02, 2006, 15:04:05", layout: "January 02, 2006, 15:04:05", expect: "2006-01-02T15:04:05Z"},
	{input: "January 02, '06, 15:04:05", layout: "January 02, '06, 15:04:05", expect: "2006-01-02T15:04:05Z"},
	{input: "January 02, 2006 15:04:05", layout: "January 02, 2006 15:04:05", expect: "2006-01-02T15:04:05Z"},
	{input: "January 02, 2006 03:04:05PM", layout: "January 02, 2006 03:04:05PM", expect: "2006-01-02T15:04:05Z"},
	{input: "September 17, 2012 at 5:00PM UTC", layout: "January 02, 2006 at 3:04PM MST", expect: "2012-09-17T17:00:00Z"},
	{input: "Tue, 11 Jul 2017 04:08:03 +0200 (CEST)", layout: "Mon, 02 Jan 2006 15:04:05 -0700 (CEST)", expect: "2017-07-11T02:08:03Z"},
	{input: "Thu May 08 17:57:51 CEST 2009", layout: "Mon Jan 02 15:04:05 MST 2006", expect: "2009-05-08T17:57:51Z"},
	{input: "Thu May 08 17:57:51 CEST '09", layout: "Mon Jan 02 15:04:05 MST '06", expect: "2009-05-08T17:57:51Z"},
	{input: "2015-02-08 03:02:00 +0300 MSK m=+0.000000001", layout: "2006-01-02 15:04:05 -0700 MST m=+0.000000000", expect: "2015-02-08T00:02:00.000000001Z"},
	{input: "2015-02-08 03:02:00 m=-0.000000001", layout: "2006-01-02 15:04:05 m=-0.000000000", expect: "2015-02-08T03:02:00.000000001Z"},
	{input: "1332151919", layout: "", expect: "2012-03-19T10:11:59Z"},                    // unix timestamp second
	{input: "1384216367111", layout: "", expect: "2013-11-12T00:32:47.111Z"},             // unix timestamp ms
	{input: "1384216367111222", layout: "", expect: "2013-11-12T00:32:47.111222Z"},       // unix timestamp us
	{input: "1384216367111222333", layout: "", expect: "2013-11-12T00:32:47.111222333Z"}, // unix timestamp ns
	{input: "20140722105203", layout: "20060102150405", expect: "2014-07-22T10:52:03Z"},
}

func TestAnytimeFormats(t *testing.T) {
	for _, item := range anytime_tests {
		str, layout, expect := item.input, item.layout, item.expect

		if layout == "" { // supported by antime.Parse, not by time.Parse
			output, err := anytime.ParseInLocation(str, time.UTC)
			assert.Nil(t, err)
			assert.Equal(t, expect, output.Format(time.RFC3339Nano))
			continue
		}

		date, err := anytime.ParseInLocation(str, time.UTC)
		assert.Nil(t, err)

		datestr := date.Format(time.RFC3339Nano)
		assert.Equal(t, expect, datestr, fmt.Sprintf("str=%s, expect=%s, date=%s", str, expect, datestr))

		output, err := time.ParseInLocation(layout, str, time.UTC)
		assert.Nil(t, err)
		assert.True(t, output.Equal(date), fmt.Sprintf("layout=%s, str=%s, expect=%s, date=%s", layout, str, expect, date))
	}
}
