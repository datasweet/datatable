package serie

import (
	"github.com/spf13/cast"
)

func Float64(v ...interface{}) Serie {
	s, _ := New(float64(0), cast.ToFloat64, compareFloat64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func Float64N(v ...interface{}) Serie {
	s, _ := New(NullFloat64{}, asNullFloat64, compareNullFloat64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func compareFloat64(a, b float64) int {
	if a == b {
		return Eq
	}
	if a < b {
		return Lt
	}
	return Gt
}

type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

func (f NullFloat64) Interface() interface{} {
	if f.Valid {
		return f.Float64
	}
	return nil
}

func asNullFloat64(i interface{}) NullFloat64 {
	var ni NullFloat64
	if i == nil {
		return ni
	}

	if v, ok := i.(NullFloat64); ok {
		return v
	}

	if v, err := cast.ToFloat64E(i); err == nil {
		ni.Float64 = v
		ni.Valid = true
	}
	return ni
}

func compareNullFloat64(a, b NullFloat64) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return compareFloat64(a.Float64, b.Float64)
}
