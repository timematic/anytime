package anytime

import (
	"fmt"
	"time"

	"github.com/longqimin/anytime/ragel"
)

func Parse(str string) (time.Time, error) {
	return parse(str, time.UTC, time.Local)
}

func ParseInLocation(str string, loc *time.Location) (time.Time, error) {
	return parse(str, loc, loc)
}

func parse(str string, defaultLoc *time.Location, targetLoc *time.Location) (time.Time, error) {
	state, err := ragel.Parse(str)
	if err != nil {
		return time.Time{}, err
	}

	// fmt.Printf("%s\n", state.String())
	ns := state.Millisecond*1000000 + state.Microsecond*1000 + state.Nanosecond
	if !state.Zoned {
		return time.Date(state.Year, time.Month(state.Month), state.Day,
			state.Hour, state.Minute, state.Second,
			ns, defaultLoc), nil
	}

	if len(state.ZoneName) != 0 {
		if zone, err := time.LoadLocation(state.ZoneName); err == nil {
			return time.Date(state.Year, time.Month(state.Month), state.Day,
				state.Hour, state.Minute, state.Second,
				ns, zone), nil
		} else { // fallback to time.Parse() or time.ParseInLocation()
			if defaultLoc != targetLoc {
				date, err := time.Parse("2006-01-02 15:04:05.000000000 "+state.ZoneName, fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
					state.Hour, state.Minute, state.Second, ns,
					state.ZoneName))
				if err == nil {
					date = date.In(targetLoc)
				}
				return date, err
			} else {
				str = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%09d %s", state.Year, state.Month, state.Day,
					state.Hour, state.Minute, state.Second, ns,
					state.ZoneName)
				layout := "2006-01-02 15:04:05.000000000 " + state.ZoneName
				date, err := time.ParseInLocation(layout, str, targetLoc)
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

	return utcdate, nil

}
