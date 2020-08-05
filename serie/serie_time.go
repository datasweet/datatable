package serie

import (
	"time"

	"github.com/datasweet/cast"
)

// Time to create a time serie
func Time(format ...string) Serie {
	return New(time.Time{}, asTime(format), compareTime)
}

// TimeN to create a time serie with nil value
func TimeN(format ...string) Serie {
	return New(NullTime{}, asNullTime(format), compareNullTime)
}

func compareTime(a, b time.Time) int {
	if a.Equal(b) {
		return Eq
	}
	if a.Before(b) {
		return Lt
	}
	return Gt
}

func asTime(formats []string) func(interface{}) time.Time {
	return func(i interface{}) time.Time {
		t, _ := cast.AsTime(i, formats...)
		return t
	}
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (t NullTime) Interface() interface{} {
	if t.Valid {
		return t.Time
	}
	return nil
}

func asNullTime(formats []string) func(interface{}) NullTime {
	return func(i interface{}) NullTime {
		var ni NullTime
		if i == nil {
			return ni
		}

		if v, ok := i.(NullTime); ok {
			return v
		}

		if v, ok := cast.AsTime(i, formats...); ok {
			ni.Time = v
			ni.Valid = true
		}
		return ni
	}
}

func compareNullTime(a, b NullTime) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return compareTime(a.Time, b.Time)
}
