package serie

import (
	"github.com/datasweet/cast"
)

func Float64(v ...interface{}) Serie {
	s := New(float64(0), asFloat64, compareFloat64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func Float64N(v ...interface{}) Serie {
	s := New(NullFloat64{}, asNullFloat64, compareNullFloat64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func asFloat64(i interface{}) float64 {
	f, _ := cast.AsFloat64(i)
	return f
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

	if v, ok := cast.AsFloat64(i); ok {
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
