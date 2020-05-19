package serie

import (
	"github.com/spf13/cast"
)

func Bool(v ...interface{}) Serie {
	s, _ := New(false, cast.ToBool, compareBool)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func BoolN(v ...interface{}) Serie {
	s, _ := New(NullBool{}, asNullBool, compareNullBool)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func compareBool(a, b bool) int {
	if a == b {
		return Eq
	}
	if !a {
		return Lt
	}
	return Gt
}

type NullBool struct {
	Bool  bool
	Valid bool
}

func (b NullBool) Interface() interface{} {
	if b.Valid {
		return b.Bool
	}
	return nil
}

func asNullBool(i interface{}) NullBool {
	var ni NullBool
	if i == nil {
		return ni
	}

	if v, ok := i.(NullBool); ok {
		return v
	}

	if v, err := cast.ToBoolE(i); err == nil {
		ni.Bool = v
		ni.Valid = true
	}
	return ni
}

func compareNullBool(a, b NullBool) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	return compareBool(a.Bool, b.Bool)
}
