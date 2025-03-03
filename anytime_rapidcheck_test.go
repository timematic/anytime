package anytime_test

import (
	"testing"
	"time"

	"github.com/longqimin/anytime"
	"pgregory.net/rapid"
)

func genLocation() *rapid.Generator[string] {
	return rapid.SampledFrom(timezone_names)
}

func genFormat() *rapid.Generator[string] {
	return rapid.SampledFrom(gen_datetime_formats())
}

func genDatetime() *rapid.Generator[time.Time] {
	return rapid.OneOf(genDate(), genTime(), genTimeNano())
}

func genDate() *rapid.Generator[time.Time] {
	return rapid.Custom(func(t *rapid.T) time.Time {
		date := time.Unix(int64(rapid.Uint32().Draw(t, "unix")), 0)
		return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	})
}
func genTime() *rapid.Generator[time.Time] {
	return rapid.Custom(func(t *rapid.T) time.Time {
		return time.Unix(int64(rapid.Uint32().Draw(t, "unix")), 0)
	})
}
func genTimeNano() *rapid.Generator[time.Time] {
	return rapid.Custom(func(t *rapid.T) time.Time {
		return time.Unix(int64(rapid.Uint32().Draw(t, "unix")), rapid.Int64().Draw(t, "nano"))
	})
}

func TestRapidCheckParse(t *testing.T) {
	prop := func(t *rapid.T) {
		layout := genFormat().Draw(t, "layout")
		date := genDatetime().Draw(t, "date")

		str := date.Format(layout)

		expect, err := time.Parse(layout, str)
		if err != nil {
			t.Fatalf("time.Parse('%s', '%s') err=%v", layout, str, err)
		}
		output, err := anytime.Parse(str)
		if err != nil {
			t.Fatalf("anytime.Parse('%s') err=%v", str, err)
		}

		if !output.Equal(expect) {
			t.Fatalf("layout=%s, str=%s, expect=%s, output=%s", layout, str, expect, output)
		}
	}

	rapid.Check(t, prop)
}

func TestRapidCheckParseInLocation(t *testing.T) {
	prop := func(t *rapid.T) {
		layout := genFormat().Draw(t, "layout")
		date := genDatetime().Draw(t, "date")

		loc := genLocation().Draw(t, "loc")
		zone, err := time.LoadLocation(loc)
		if err != nil {
			t.Skip()
		}

		str := date.Format(layout)

		expect, err := time.ParseInLocation(layout, str, zone)
		if err != nil {
			t.Fatalf("time.ParseInLocation('%s', '%s', '%s') err=%v", layout, str, loc, err)
		}
		output, err := anytime.ParseInLocation(str, zone)
		if err != nil {
			t.Fatalf("anytime.ParseInLocation('%s','%s') err=%v", str, loc, err)
		}

		if !output.Equal(expect) {
			t.Fatalf("layout=%s, str=%s, loc=%s, expect=%s, output=%s", layout, str, loc, expect, output)
		}
	}

	rapid.Check(t, prop)
}
