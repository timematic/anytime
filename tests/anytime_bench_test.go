package anytime_test

import (
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/timematic/anytime"
)

var (
	dateOnly    = time.Now().Format(time.DateOnly)
	dateWithTZ  = time.Now().In(time.FixedZone("PST", -8*3600)).Format("2006-01-02 -0800")
	rfc3339    = time.Now().Format(time.RFC3339)
	rfc3339Nano = time.Now().Format(time.RFC3339Nano)
	unixTs      = "1714000000"
	timeOnly    = "15:04:05"
	datetime     = "2006-01-02 15:04:05"
	datetimeNano = "2006-01-02 15:04:05.123456789"
)

func BenchmarkStdParse_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse(time.DateOnly, dateOnly)
	}
}

func BenchmarkAnytimeParse_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(dateOnly)
	}
}

func BenchmarkAraddonParseStrict_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(dateOnly)
	}
}

func BenchmarkStdParse_DateWithTZ(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse("2006-01-02 -0700", dateWithTZ)
	}
}

func BenchmarkAnytimeParse_DateWithTZ(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(dateWithTZ)
	}
}

func BenchmarkAraddonParseStrict_DateWithTZ(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(dateWithTZ)
	}
}

func BenchmarkStdParse_RFC3339(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse(time.RFC3339, rfc3339)
	}
}

func BenchmarkAnytimeParse_RFC3339(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(rfc3339)
	}
}

func BenchmarkAraddonParseStrict_RFC3339(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(rfc3339)
	}
}

func BenchmarkStdParse_RFC3339Nano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse(time.RFC3339Nano, rfc3339Nano)
	}
}

func BenchmarkAnytimeParse_RFC3339Nano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(rfc3339Nano)
	}
}

func BenchmarkAraddonParseStrict_RFC3339Nano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(rfc3339Nano)
	}
}

func BenchmarkStdParse_Layout(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse(time.RFC1123, "Mon, 02 Jan 2006 15:04:05 UTC")
	}
}

func BenchmarkAnytimeParse_Layout(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse("Mon, 02 Jan 2006 15:04:05 UTC")
	}
}

func BenchmarkAraddonParseStrict_Layout(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict("Mon, 02 Jan 2006 15:04:05 UTC")
	}
}

func BenchmarkStdParse_UnixTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Unix(1714000000, 0)
	}
}

func BenchmarkAnytimeParse_UnixTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(unixTs)
	}
}

func BenchmarkAraddonParseStrict_UnixTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(unixTs)
	}
}

func BenchmarkStdParse_TimeOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse("15:04:05", timeOnly)
	}
}

func BenchmarkAnytimeParse_TimeOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(timeOnly)
	}
}

func BenchmarkAraddonParseStrict_TimeOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(timeOnly)
	}
}

func BenchmarkStdParse_Datetime(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse("2006-01-02 15:04:05", datetime)
	}
}

func BenchmarkAnytimeParse_Datetime(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(datetime)
	}
}

func BenchmarkAraddonParseStrict_Datetime(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(datetime)
	}
}

func BenchmarkStdParse_DatetimeNano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse("2006-01-02 15:04:05.000000000", datetimeNano)
	}
}

func BenchmarkAnytimeParse_DatetimeNano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(datetimeNano)
	}
}

func BenchmarkAraddonParseStrict_DatetimeNano(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseStrict(datetimeNano)
	}
}
