package anytime

import (
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

	return state.AsTime(defaultLoc, targetLoc)

}
