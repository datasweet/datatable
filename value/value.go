package value

import (
	"fmt"
)

const nullValueStr = "#NULL!"

// Value is a value in a serie
type Value interface {
	// Type to get the type of the value
	Type() Type

	// Set the value
	Set(v interface{}) Value

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

const (
	Lt = -1
	Eq = 0
	Gt = 1
)

type Builder func(v ...interface{}) Value
