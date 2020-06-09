package serie

import (
	"fmt"
	"strings"
)

func Raw(v ...interface{}) Serie {
	s, _ := New(RawValue{}, asRawValue, compareRawValue)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

type RawValue struct {
	Value interface{}
	Valid bool
}

func (r RawValue) Interface() interface{} {
	if r.Valid {
		return r.Value
	}
	return nil
}

func (r RawValue) String() string {
	return fmt.Sprint(r.Value)
}

func asRawValue(i interface{}) RawValue {
	if rv, ok := i.(RawValue); ok {
		return rv
	}
	var r RawValue
	if i == nil {
		return r
	}
	r.Valid = true
	r.Value = i
	return r
}

func compareRawValue(a, b RawValue) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
	}
	return strings.Compare(a.String(), b.String())
}
