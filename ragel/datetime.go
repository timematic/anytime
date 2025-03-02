package ragel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
