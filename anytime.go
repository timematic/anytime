package anytime

import (
	"time"
)

// Parse parses a time string and returns the corresponding time.Time value.
// It attempts to parse the input string in multiple common time formats
// without requiring a specific format layout string.
// The returned time is in UTC timezone.
//
// Supported formats include various date/time formats, Unix timestamps,
// and formats with timezone information.
func Parse(str string) (time.Time, error) {
	return parse(str, time.UTC, time.Local)
}

// ParseInLocation parses a time string and returns the corresponding time.Time value.
// It is similar to Parse but interprets ambiguous timezone abbreviations
// (like "EST", "CST") using the provided location.
//
// If the input string contains an explicit timezone offset or IANA timezone name,
// that will be used instead of the provided location.
func ParseInLocation(str string, loc *time.Location) (time.Time, error) {
	return parse(str, loc, loc)
}

func parse(str string, defaultLoc *time.Location, targetLoc *time.Location) (time.Time, error) {
	state, err := ragelParse(str)
	if err != nil {
		return time.Time{}, err
	}

	return state.AsTime(defaultLoc, targetLoc)

}
