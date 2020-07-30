package serie

import (
	"strings"

	"github.com/spf13/cast"
)

// String to create a new string serie
func String(v ...interface{}) Serie {
	s := New("", cast.ToString, strings.Compare)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

// StringN to create a new serie with null value handling
func StringN(v ...interface{}) Serie {
	s := New(NullString{}, asNullString, compareNullString)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

type NullString struct {
	String string
	Valid  bool
}

func (s NullString) Interface() interface{} {
	if s.Valid {
		return s.String
	}
	return nil
}

func asNullString(i interface{}) NullString {
	var ns NullString
	if i == nil {
		return ns
	}

	if v, ok := i.(NullString); ok {
		return v
	}

	if v, err := cast.ToStringE(i); err == nil {
		ns.String = v
		ns.Valid = true
	}
	return ns
}

func compareNullString(a, b NullString) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return strings.Compare(a.String, b.String)
}
