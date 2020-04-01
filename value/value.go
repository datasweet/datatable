package value

import (
	"fmt"
	"reflect"
)

const nullValueStr = "#NULL!"

// Value is a value in a serie
type Value interface {
	// Type to get the type of the value
	Type() Type

	// Set the value
	Set(v interface{})

	// Val to get the value
	// nil if the value is not valid
	Val() interface{}

	// IsValid to known if the value is valid for the current type
	IsValid() bool

	// Compare to compare the current value to another
	// returns an integer. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	Compare(v Value) int

	// Clone to clone the value
	Clone() Value

	fmt.Stringer
}

type Type string

type Builder func(v interface{}) Value

type Predicate func(val Value) bool

// TypeFromReflect to creates a value type from a reflect.Type
func TypeFromReflect(typ reflect.Type) Type {
	if typ == nil {
		return Raw
	}
	return Type(typ.Name())
}

const (
	Lt = -1
	Eq = 0
	Gt = 1
)
