package serie

import (
	"github.com/spf13/cast"
)

func Int32(v ...interface{}) Serie {
	s, _ := New(int32(0), cast.ToInt32, compareInt32)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func Int32N(v ...interface{}) Serie {
	s, _ := New(NullInt32{}, asNullInt32, compareNullInt32)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func compareInt32(a, b int32) int {
	if a == b {
		return Eq
	}
	if a < b {
		return Lt
	}
	return Gt
}

type NullInt32 struct {
	Int32 int32
	Valid bool
}

func (i NullInt32) Interface() interface{} {
	if i.Valid {
		return i.Int32
	}
	return nil
}

func asNullInt32(i interface{}) NullInt32 {
	var ni NullInt32
	if i == nil {
		return ni
	}

	if v, ok := i.(NullInt32); ok {
		return v
	}

	if v, err := cast.ToInt32E(i); err == nil {
		ni.Int32 = v
		ni.Valid = true
	}
	return ni
}

func compareNullInt32(a, b NullInt32) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	return compareInt32(a.Int32, b.Int32)
}