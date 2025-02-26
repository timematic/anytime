package anytime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/longqimin/anytime"

	"github.com/stretchr/testify/assert"
	_ "pgregory.net/rapid"
)

// func TestSortStrings(t *testing.T) {
// 	rapid.Check(t, func(t *rapid.T) {

// 		year := rapid.Int64Range(0, 200).Draw(t, "year")
// 		month := year*12 + rapid.Int64Range(0, 12).Draw(t, "month")
// 		day := month*30 + rapid.Int64Range(0, 31).Draw(t, "day")
// 		hour := day*24 + rapid.Int64().Draw(t, "hour")
// 		minute := hour*60 + rapid.Int64Range(0, 60).Draw(t, "minute")
// 		second := minute*60 + rapid.Int64Range(0, 60).Draw(t, "second")
// 		nanosecond := second * (10 ^ 9)

// 		zero_unix := time.Unix(0, 0)
// 		dt := zero_unix.Add(time.Duration(nanosecond - nanosecond%1000))
// 		fmt.Println(dt)
// 		fmt.Println(dt.Format("2006-01-02"))
// 		fmt.Println(dt.Format("2006-01-02T15"))
// 		fmt.Println(dt.Format("2006-01-02T15:04"))
// 		fmt.Println(dt.Format("2006-01-02T15:04:05"))
// 		fmt.Println(dt.Format("2006-01-02T15:04:05.000000"))
// 		// s := rapid.SliceOf(rapid.String()).Draw(t, "s")
// 		// sort.Strings(s)
// 		// if !sort.StringsAreSorted(s) {
// 		// 	t.Fatalf("unsorted after sort: %v", s)
// 		// }
// 	})
// }

func TestParse(t *testing.T) {
	now := time.Now()

	date_fmts := []string{
		"2006-01-02", "2006/01/02", "20060102", "2006-Jan-02", "2006/Jan/02",
	}
	time_fmts := []string{
		"15", "1504", "150405", "150405.0", "150405.00", "150405.000", "150405.0000", "150405.00000", "150405.000000",
		"15:04", "15:04:05", "15:04:05.0", "15:04:05.00", "15:04:05.000", "15:04:05.0000", "15:04:05.00000", "15:04:05.000000",
		"03PM", "03:04PM", "03:04:05PM", "03:04:05.0PM", "03:04:05.00PM", "03:04:05.000PM", "03:04:05.0000PM", "03:04:05.00000PM", "03:04:05.000000PM",
	}
	zone_fmts := []string{
		"Z",
	}

	for _, date_fmt := range date_fmts {
		{
			str := now.Format(date_fmt)
			expect, err := time.Parse(date_fmt, str)
			assert.Nil(t, err)

			date, err := anytime.Parse(str)
			assert.Nil(t, err)
			// fmt.Println(date)
			assert.Equal(t, expect, date)
		}

		for _, zone_fmt := range zone_fmts {
			fmtstr := fmt.Sprintf("%s %s", date_fmt, zone_fmt)
			str := now.Format(fmtstr)
			expect, err := time.Parse(fmtstr, str)
			assert.Nil(t, err)

			date, err := anytime.Parse(str)
			assert.Nil(t, err)
			// fmt.Println(date)
			assert.Equal(t, expect, date)
		}
	}

	for _, date_fmt := range date_fmts {
		for _, time_fmt := range time_fmts {
			for _, sp := range []string{" ", "T"} {
				fmtstr := fmt.Sprintf("%s%s%s", date_fmt, sp, time_fmt)

				str := now.Format(fmtstr)
				fmt.Println(str)
				expect, err := time.Parse(fmtstr, str)
				assert.Nil(t, err)

				date, err := anytime.Parse(str)
				assert.Nil(t, err)
				// fmt.Println(date)
				assert.Equal(t, expect, date)

			}

		}
	}

}
