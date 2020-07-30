package serie

import (
	"time"

	"github.com/spf13/cast"
)

func Time(v ...interface{}) Serie {
	s := New(time.Time{}, cast.ToTime, compareTime)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func TimeN(v ...interface{}) Serie {
	s := New(NullTime{}, asNullTime, compareNullTime)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
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

func asNullTime(i interface{}) NullTime {
	var ni NullTime
	if i == nil {
		return ni
	}

	if v, ok := i.(NullTime); ok {
		return v
	}

	if v, err := cast.ToTimeE(i); err == nil {
		ni.Time = v
		ni.Valid = true
	}
	return ni
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
