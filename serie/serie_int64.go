package serie

import (
	"github.com/datasweet/cast"
)

func Int64(v ...interface{}) Serie {
	s := New(int64(0), asInt64, compareInt64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func Int64N(v ...interface{}) Serie {
	s := New(NullInt64{}, asNullInt64, compareNullInt64)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func asInt64(i interface{}) int64 {
	n, _ := cast.AsInt64(i)
	return n
}

func compareInt64(a, b int64) int {
	if a == b {
		return Eq
	}
	if a < b {
		return Lt
	}
	return Gt
}

type NullInt64 struct {
	Int64 int64
	Valid bool
}

func (i NullInt64) Interface() interface{} {
	if i.Valid {
		return i.Int64
	}
	return nil
}

func asNullInt64(i interface{}) NullInt64 {
	var ni NullInt64
	if i == nil {
		return ni
	}

	if v, ok := i.(NullInt64); ok {
		return v
	}

	if v, ok := cast.AsInt64(i); ok {
		ni.Int64 = v
		ni.Valid = true
	}
	return ni
}

func compareNullInt64(a, b NullInt64) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return compareInt64(a.Int64, b.Int64)
}
