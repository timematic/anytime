package ragel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

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

// ParsedDatetime is the struct representing a parsed time value.
type ParsedDatetime struct {
	Ad_bc                                ADBC
	Year                                 int
	Month                                int
	Day                                  int
	DayOfYear                            int
	Hour, Minute, Second                 int
	Millisecond, Microsecond, Nanosecond int

	Zoned                     bool
	ZoneOffsetHour            int
	ZoneOffsetMinute          int
	NegtiveZoneOffset         bool
	ZoneOffsetIsValid         bool
	ZoneName                  string // e.g., "MST"
	MonotonicOffsetNanosecond int64
}

func (state *ParsedDatetime) String() string {
	return fmt.Sprintf(
		"ab_bc=%v, year=%d, month=%d, day=%d, day_of_year=%d, hour=%d, minute=%d, second=%d, ms=%d, us=%d, ns=%d, zoned=%v ZoneName=%s ZoneOffsetHour=%d ZoneOffstMinute=%d m=%d\n",
		state.Ad_bc, state.Year, state.Month, state.Day, state.DayOfYear,
		state.Hour, state.Minute, state.Second, state.Millisecond, state.Microsecond, state.Nanosecond,
		state.Zoned,
		state.ZoneName, state.ZoneOffsetHour, state.ZoneOffsetMinute, state.MonotonicOffsetNanosecond)
}

func (state *ParsedDatetime) AsTime(defaultLoc *time.Location, targetLoc *time.Location) (time.Time, error) {
	// fmt.Printf("%s\n", state.String())

	day_of_year_fmt := state.DayOfYear != 0
	if day_of_year_fmt {
		state.Month = 1
	}

	var date time.Time
	ns := state.Millisecond*1000000 + state.Microsecond*1000 + state.Nanosecond + int(state.MonotonicOffsetNanosecond)
	if !state.Zoned {
		date = time.Date(state.Year, time.Month(state.Month), state.Day,
			state.Hour, state.Minute, state.Second,
			ns, defaultLoc)
		if day_of_year_fmt {
			return date.AddDate(0, 0, state.DayOfYear), nil
		}
		return date, nil
	}

	if state.ZoneOffsetIsValid {
		state.ZoneName = ""
	}

	if len(state.ZoneName) != 0 {
		_, ambiguous := ambiguousTimeZoneAbbrs[state.ZoneName]
		if zone, err := time.LoadLocation(state.ZoneName); err == nil && !ambiguous {
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
					date, err = time.Parse("2006.002 15:04:05.000000000 MST", fmt.Sprintf("%04d.%03d %02d:%02d:%02d.%09d %s", state.Year, state.DayOfYear,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName))
				} else {
					date, err = time.Parse("2006-01-02 15:04:05.000000000 MST", fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName))
				}
				return date, err
			} else {
				if day_of_year_fmt {
					str := fmt.Sprintf("%04d.%03d %02d:%02d:%02d.%09d %s", state.Year, state.DayOfYear,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName)
					layout := "2006.002 15:04:05.000000000 MST"
					date, err = time.ParseInLocation(layout, str, targetLoc)
				} else {
					str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
						state.Hour, state.Minute, state.Second, ns,
						state.ZoneName)
					layout := "2006-01-02 15:04:05.000000000 MST"
					date, err = time.ParseInLocation(layout, str, targetLoc)
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

	if day_of_year_fmt {
		return utcdate.AddDate(0, 0, state.DayOfYear), nil
	}
	return utcdate, nil
}

var ambiguousTimeZoneAbbrs = map[string]bool{
	"CST":  true, // Central Standard Time (UTC-6), China Standard Time (UTC+8), Cuba Standard Time (UTC-5)
	"PST":  true, // Pacific Standard Time (UTC-8), Philippine Standard Time (UTC+8)
	"EST":  true, // Eastern Standard Time (UTC-5), Eastern Standard Time (Australia) (UTC+10)
	"MST":  true, // Mountain Standard Time (UTC-7), Moscow Standard Time (UTC+3)
	"IST":  true, // Indian Standard Time (UTC+5:30), Irish Standard Time (UTC+1), Israel Standard Time (UTC+2)
	"BST":  true, // British Summer Time (UTC+1), Bangladesh Standard Time (UTC+6), Bougainville Standard Time (UTC+11)
	"GST":  true, // Gulf Standard Time (UTC+4), South Georgia Standard Time (UTC-2)
	"AST":  true, // Atlantic Standard Time (UTC-4), Arabia Standard Time (UTC+3), Alaska Standard Time (UTC-9)
	"SST":  true, // Samoa Standard Time (UTC-11), Singapore Standard Time (UTC+8)
	"ACT":  true, // Acre Time (UTC-5), ASEAN Common Time (UTC+8)
	"CAT":  true, // Central Africa Time (UTC+2), Central America Time (UTC-6)
	"EAT":  true, // East Africa Time (UTC+3), East Antarctica Time (UTC+10)
	"WAT":  true, // West Africa Time (UTC+1), West Antarctica Time (UTC-1)
	"HST":  true, // Hawaii-Aleutian Standard Time (UTC-10), Heure Standard de Tahiti (UTC-10)
	"AKST": true, // Alaska Standard Time (UTC-9), Anadyr Standard Time (UTC+12)
}
