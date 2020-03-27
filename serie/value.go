package serie

import (
	"fmt"
	"reflect"
)

// detect system 32 or 64 bits
const wordSize = 32 << (^uint(0) >> 32 & 1)
const nullValueStr = "#NULL!"

// Value is a value in a serie
type Value interface {
	// Type to get the type of the value
	Type() ValueType

	// Set the value
	Set(v interface{})

	// Val to get the value
	// nil if the value is not valid
	Val() interface{}

	// IsValid to known if the value is valid for the current type
	IsValid() bool

	// Clone to clone the value
	Clone() Value

	fmt.Stringer
}

type ValueType string

const Raw = ValueType("raw")

type ValueBuilder func(v interface{}) Value

type ValuePredicate func(val Value) bool

func valueTypeFromReflect(typ reflect.Type) ValueType {
	if typ == nil {
		return Raw
	}
	return ValueType(typ.Name())
}
