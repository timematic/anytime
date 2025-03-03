package anytime_test

import (
	"fmt"
	"strings"
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
	// time.RFC822,
	time.RFC822Z,
	// time.RFC850,
	// time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.DateTime,
	time.DateOnly,
}

var timezone_names = []string{
	"America/New_York",    //	北美东部时间（纽约） UTC-05:00 / UTC-04:00 (夏令时)
	"America/Chicago",     //	北美中部时间（芝加哥） UTC-06:00 / UTC-05:00 (夏令时)
	"America/Denver",      //	北美山区时间（丹佛） UTC-07:00 / UTC-06:00 (夏令时)
	"America/Los_Angeles", //	北美太平洋时间（洛杉矶） UTC-08:00 / UTC-07:00 (夏令时)
	"Europe/London",       //	格林尼治时间（伦敦） UTC+00:00 / UTC+01:00 (夏令时)
	"Europe/Paris",        //	中欧时间（巴黎） UTC+01:00 / UTC+02:00 (夏令时)
	"Asia/Shanghai",       //	中国标准时间（上海） UTC+08:00
	"Asia/Tokyo",          //	日本标准时间（东京） UTC+09:00
	"Asia/Kolkata",        //	印度标准时间（加尔各答） UTC+05:30
	"Australia/Sydney",    //	澳大利亚东部时间（悉尼） UTC+10:00 / UTC+11:00 (夏令时)
}

var timezone_abbrnames = []string{
	"UTC",  // 协调世界时（基准时区）	UTC+00:00
	"GMT",  //	格林尼治标准时间（等同于 UTC）	UTC+00:00
	"EST",  //	北美东部标准时间	UTC-05:00
	"EDT",  //	北美东部夏令时间	UTC-04:00
	"CST",  //	北美中部标准时间	UTC-06:00
	"CDT",  //	北美中部夏令时间	UTC-05:00
	"PST",  //	北美太平洋标准时间	UTC-08:00
	"PDT",  //	北美太平洋夏令时间	UTC-07:00
	"CET",  //	中欧时间	UTC+01:00
	"CEST", //	中欧夏令时间	UTC+02:00
	"EET",  //	东欧时间	UTC+02:00
	"EEST", //	东欧夏令时间	UTC+03:00
	"AEST", //	澳大利亚东部标准时间	UTC+10:00
	"AEDT", //	澳大利亚东部夏令时间	UTC+11:00
	"IST",  //	印度标准时间	UTC+05:30
	"JST",  //	日本标准时间	UTC+09:00
	"CST",  //	中国标准时间	UTC+08:00
	"HKT",  //	香港时间	UTC+08:00
	"SGT",  //	新加坡时间	UTC+08:00
}

func TestParseInLocation(t *testing.T) {
	for _, fmtstr := range golang_time_formats {
		now := time.Now()
		str := now.Format(fmtstr)
		if strings.HasSuffix(str, "Z") {
			str = strings.TrimSuffix(str, "Z") + "+00:00"
		}

		expect, err := time.ParseInLocation(fmtstr, str, time.Local)
		assert.Nil(t, err)
		date, err := anytime.ParseInLocation(str, time.Local)
		assert.Nil(t, err)
		assert.True(t, expect.Equal(date), fmt.Sprintf("loc=%s format=%v, str=%v, expect=%v, output=%v", time.Local, fmtstr, str, expect, date))

		expect, err = time.ParseInLocation(fmtstr, str, time.UTC)
		assert.Nil(t, err)
		date, err = anytime.ParseInLocation(str, time.UTC)
		assert.Nil(t, err)
		assert.True(t, expect.Equal(date), fmt.Sprintf("loc=%s format=%v, str=%v, expect=%v, output=%v", time.UTC, fmtstr, str, expect, date))

		for _, name := range timezone_names {
			zone, err := time.LoadLocation(name)
			assert.Nil(t, err)
			expect, err := time.ParseInLocation(fmtstr, str, zone)
			assert.Nil(t, err)

			date, err := anytime.ParseInLocation(str, zone)
			assert.Nil(t, err)
			assert.True(t, expect.Equal(date), fmt.Sprintf("loc=%s format=%v, str=%v, expect=%v, output=%v", zone, fmtstr, str, expect, date))
		}
	}
}

func TestParse(t *testing.T) {
	for _, fmtstr := range golang_time_formats {
		now := time.Now()
		str := now.Format(fmtstr)

		expect, err := time.Parse(fmtstr, str)
		assert.Nil(t, err)
		date, err := anytime.Parse(str)
		assert.Nil(t, err)
		assert.True(t, expect.Equal(date), fmt.Sprintf("format=%v, str=%v, expect=%v, output=%v", fmtstr, str, expect, date))
	}
}

var date_fmts = []string{
	"2006-01-02", "2006-1-02", "2006-01-2", "2006-1-2", "2006-Jan-02", "2006-Jan-2", // yyyy-mm-dd
	"2006/01/02", "2006/1/02", "2006/01/2", "2006/1/2", "2006/Jan/02", "2006/Jan/2", // yyyy/mm/dd
	"20060102", // yyyymmdd

}

var time_fmts = []string{
	"15", "1504", "150405", "150405.0", "150405.00", "150405.000", "150405.0000", "150405.00000", "150405.000000",
	"15:04", "15:04:05", "15:04:05.0", "15:04:05.00", "15:04:05.000", "15:04:05.0000", "15:04:05.00000", "15:04:05.000000",
	"03PM", "03:04PM", "03:04:05PM", "03:04:05.0PM", "03:04:05.00PM", "03:04:05.000PM", "03:04:05.0000PM", "03:04:05.00000PM", "03:04:05.000000PM",
}
var zone_fmts = []string{
	"", "-07", "-0700", "-07:00", "Z07:00",
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
	return fmts
}

func TestParseDateOnly(t *testing.T) {
	now := time.Now()

	for _, layout := range gen_date_only_formats() {
		str := now.Format(layout)
		expect, err := time.ParseInLocation(layout, str, time.Local)
		assert.Nil(t, err)

		date, err := anytime.ParseInLocation(str, time.Local)
		assert.Nil(t, err)
		assert.True(t, expect.Equal(date), fmt.Sprintf("input=%s, layout=%s, expect=%s, output=%s", str, layout, expect, date))

	}
}

func TestParseDateTime(t *testing.T) {
	now := time.Now()

	for _, layout := range gen_datetime_formats() {
		str := now.Format(layout)
		expect, err := time.ParseInLocation(layout, str, time.Local)
		assert.Nil(t, err)

		date, err := anytime.ParseInLocation(str, time.Local)
		assert.Nil(t, err)
		assert.True(t, expect.Equal(date), fmt.Sprintf("input=%s, layout=%s, expect=%s, output=%s", str, layout, expect, date))

	}
}

var dateonly = time.Now().Format(time.DateOnly)

func BenchmarkStdParse_DateOnly(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Parse(time.DateOnly, dateonly)
	}
}

func BenchmarkAnytimeParse_DateOnly(b *testing.B) {
	for n := 0; n < b.N; n++ {
		anytime.Parse(dateonly)
	}
}

var rfc3339 = time.Now().Format(time.RFC3339)

func BenchmarkStdParse_RFC3339(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Parse(time.RFC3339, rfc3339)
	}
}

func BenchmarkAnytimeParse_RFC3339(b *testing.B) {
	for n := 0; n < b.N; n++ {
		anytime.Parse(rfc3339)
	}
}
