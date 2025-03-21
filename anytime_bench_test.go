package anytime_test

import (
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/timematic/anytime"
)

var dateonly = time.Now().Format(time.DateOnly)

func BenchmarkStdParse_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		time.Parse(time.DateOnly, dateonly)
	}
}

func BenchmarkAnytimeParse_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		anytime.Parse(dateonly)
	}
}

func BenchmarkAraddonParseAny_DateOnly(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseAny(dateonly)
	}
}

var rfc3339 = time.Now().Format(time.RFC3339)

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
func BenchmarkAraddonParseAny_RFC3339(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		dateparse.ParseAny(rfc3339)
	}
}
