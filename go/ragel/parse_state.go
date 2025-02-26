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

type Datetime struct {
	Ad_bc       ADBC
	Year        int
	Month       int
	Day         int
	Day_of_year int
	Hour        int
	Minute      int
	Second      int
	Millisecond int
	Microsecond int

	Zoned bool
	// zone or offset
	Zone_name_or_abbrev string
	Negative_offset     bool
	Offset_hour         int
	Offset_minute       int
}

func (state *Datetime) String() string {
	return fmt.Sprintf(
		"ab_bc=%v, year=%d, month=%d, day=%d, day_of_year=%d, hour=%d, minute=%d, second=%d, millisecond=%d, microsecond=%d\n",
		state.Ad_bc, state.Year, state.Month, state.Day, state.Day_of_year, state.Hour, state.Minute, state.Second, state.Millisecond, state.Microsecond)
}
