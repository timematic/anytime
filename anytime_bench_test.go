package anytime_test

import (
	"testing"
	"time"

	"github.com/longqimin/anytime"
)

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
