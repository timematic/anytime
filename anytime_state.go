package anytime

import (
	"fmt"
	"time"
)

// parsedTime is the struct representing a parsed time value.
type parsedTime struct {
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

func (state *parsedTime) String() string {
	return fmt.Sprintf(
		"ab_bc=%v, year=%d, month=%d, day=%d, day_of_year=%d, hour=%d, minute=%d, second=%d, ms=%d, us=%d, ns=%d, zoned=%v ZoneName=%s ZoneOffsetHour=%d ZoneOffstMinute=%d m=%d\n",
		state.Ad_bc, state.Year, state.Month, state.Day, state.DayOfYear,
		state.Hour, state.Minute, state.Second, state.Millisecond, state.Microsecond, state.Nanosecond,
		state.Zoned,
		state.ZoneName, state.ZoneOffsetHour, state.ZoneOffsetMinute, state.MonotonicOffsetNanosecond)
}

func (state *parsedTime) AsTime(defaultLoc *time.Location, targetLoc *time.Location) (time.Time, error) {
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
