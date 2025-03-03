package ragel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ADBC int32

const (
	ADBC_AD ADBC = 0
	ADBC_BC ADBC = 1
)

type AMPM int32

const (
	AMPM_AM AMPM = 0
	AMPM_PM AMPM = 1
)

func parse_ampm(s string) (AMPM, error) {
	if len(s) < 2 {
		return AMPM_AM, errors.New("parse_ampm too short")
	}
	upper := strings.ToUpper(s[:2])
	if upper == "AM" {
		return AMPM_AM, nil
	}
	if upper == "PM" {
		return AMPM_PM, nil
	}
	return AMPM_AM, errors.New("parse_ampm invalid")
}

func parse_digits(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

// ParsedDatetime is the struct representing a parsed time value.
type ParsedDatetime struct {
	Ad_bc                                ADBC
	Year                                 int
	Month                                int
	Day                                  int
	DayOfYear                            int
	Hour, Minute, Second                 int
	Millisecond, Microsecond, Nanosecond int

	Zoned             bool
	ZoneOffsetHour    int
	ZoneOffsetMinute  int
	NegtiveZoneOffset bool
	ZoneName          string // e.g., "MST"

}

func (state *ParsedDatetime) String() string {
	return fmt.Sprintf(
		"ab_bc=%v, year=%d, month=%d, day=%d, day_of_year=%d, hour=%d, minute=%d, second=%d, ms=%d, us=%d, ns=%d, zoned=%v ZoneName=%s ZoneOffsetHour=%d ZoneOffstMinute=%d\n",
		state.Ad_bc, state.Year, state.Month, state.Day, state.DayOfYear,
		state.Hour, state.Minute, state.Second, state.Millisecond, state.Microsecond, state.Nanosecond,
		state.Zoned,
		state.ZoneName, state.ZoneOffsetHour, state.ZoneOffsetMinute)
}

func (state *ParsedDatetime) AsTime(defaultLoc *time.Location, targetLoc *time.Location) (time.Time, error) {
	// fmt.Printf("%s\n", state.String())

	day_of_year_fmt := state.DayOfYear != 0
	if day_of_year_fmt {
		state.DayOfYear -= 1
		if state.Day == 0 {
			state.Day = 1
		}
		if state.Month == 0 {
			state.Month = 1
		}
	}

	var date time.Time
	ns := state.Millisecond*1000000 + state.Microsecond*1000 + state.Nanosecond
	if !state.Zoned {
		date = time.Date(state.Year, time.Month(state.Month), state.Day,
			state.Hour, state.Minute, state.Second,
			ns, defaultLoc)
		if day_of_year_fmt {
			return date.AddDate(0, 0, state.DayOfYear), nil
		}
		return date, nil
	}

	if len(state.ZoneName) != 0 {
		if zone, err := time.LoadLocation(state.ZoneName); err == nil {
			date = time.Date(state.Year, time.Month(state.Month), state.Day,
				state.Hour, state.Minute, state.Second,
				ns, zone)
			if day_of_year_fmt {
				return date.AddDate(0, 0, state.DayOfYear), nil
			}
			return date, nil
		} else { // fallback to time.Parse() or time.ParseInLocation()
			if defaultLoc != targetLoc {
				if day_of_year_fmt {
					date, err = time.Parse("2006002 15:04:05.000000000 "+state.ZoneName, fmt.Sprintf("%04d%03d %02d:%02d:%02d.%09d %s", state.Year, state.DayOfYear,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName))
				} else {
					date, err = time.Parse("2006-01-02 15:04:05.000000000 "+state.ZoneName, fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName))
				}

				if err == nil {
					date = date.In(targetLoc)
				}
				return date, err
			} else {
				if day_of_year_fmt {
					str := fmt.Sprintf("%04d%03d %02d:%02d:%02d.%09d %s", state.Year, state.DayOfYear,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName)
					layout := "2006002 15:04:05.000000000 " + state.ZoneName
					date, err = time.ParseInLocation(layout, str, targetLoc)
				} else {
					str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName)
					layout := "2006-01-02 15:04:05.000000000 " + state.ZoneName
					date, err = time.ParseInLocation(layout, str, targetLoc)
				}

				if err == nil {
					date = date.AddDate(0, 0, state.DayOfYear)
				}
				return date, err
			}
		}
	}

	utcdate := time.Date(state.Year, time.Month(state.Month), state.Day,
		state.Hour, state.Minute, state.Second,
		ns, time.UTC)

	offset := time.Duration(state.ZoneOffsetHour*int(time.Hour) + state.ZoneOffsetMinute*int(time.Minute))
	if state.NegtiveZoneOffset {
		utcdate = utcdate.Add(offset)
	} else {
		utcdate = utcdate.Add(-offset)
	}

	return utcdate.AddDate(0, 0, state.DayOfYear), nil
}

func parse_year_2_digits(str string) int {
	year, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	if year >= 69 { // Unix time starts Dec 31 1969 in some time zones
		year += 1900
	} else {
		year += 2000
	}
	return year
}
