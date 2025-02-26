package anytime

import (
	"time"

	"github.com/longqimin/anytime/ragel"
)

func Parse(str string) (time.Time, error) {
	state, err := ragel.Parse(str)
	if err != nil {
		return time.Time{}, err
	}

	// TODO(lqm)
	return time.Date(state.Year, time.Month(state.Month), state.Day, state.Hour, state.Minute, state.Second, state.Millisecond*1000000+state.Microsecond*1000, time.UTC), nil
}
