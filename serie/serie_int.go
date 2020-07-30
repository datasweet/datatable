package serie

import (
	"github.com/spf13/cast"
)

func Int(v ...interface{}) Serie {
	s := New(0, cast.ToInt, compareInt)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func IntN(v ...interface{}) Serie {
	s := New(NullInt{}, asNullInt, compareNullInt)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func compareInt(a, b int) int {
	if a == b {
		return Eq
	}
	if a < b {
		return Lt
	}
	return Gt
}

type NullInt struct {
	Int   int
	Valid bool
}

func (i NullInt) Interface() interface{} {
	if i.Valid {
		return i.Int
	}
	return nil
}

func asNullInt(i interface{}) NullInt {
	var ni NullInt
	if i == nil {
		return ni
	}

	if v, ok := i.(NullInt); ok {
		return v
	}

	if v, err := cast.ToIntE(i); err == nil {
		ni.Int = v
		ni.Valid = true
	}
	return ni
}

func compareNullInt(a, b NullInt) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return compareInt(a.Int, b.Int)
}
