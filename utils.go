package anytime

import (
	"errors"
	"strconv"
	"strings"
)

// ADBC represents AD (Anno Domini) or BC (Before Christ) era for years.
type ADBC int32

const (
	ADBC_AD ADBC = 0
	ADBC_BC ADBC = 1
)

// AMPM represents AM or PM indicator for 12-hour time format.
type AMPM int32

const (
	AMPM_AM AMPM = 0
	AMPM_PM AMPM = 1
)

// parse_ampm parses AM/PM indicator from string.
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

// parse_digits parses a string into an integer.
// Panics on error - should only be called with validated digit strings.
func parse_digits(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

// parse_year_2_digits parses a 2-digit year into a full year.
// Years >= 69 are treated as 1969-1999, others as 2000-2068.
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

// ambiguousTimeZoneAbbrs contains timezone abbreviations that have multiple meanings
// and should be handled with caution.
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
