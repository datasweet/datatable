package value

import (
	"github.com/datasweet/cast"
)

const String = Type("string")

type stringValue struct {
	val   string
	valid bool
}

func NewString(v interface{}) Value {
	value := &stringValue{}
	value.Set(v)
	return value
}

func (value *stringValue) Type() Type {
	return String
}

func (value *stringValue) Val() interface{} {
	if value.valid {
		return value.val
	}
	return nil
}

func (value *stringValue) Set(v interface{}) {
	value.val = ""
	value.valid = false

	if casted, ok := cast.AsString(v); ok {
		value.val = casted
		value.valid = true
	}
}

func (value *stringValue) IsValid() bool {
	return value.valid
}

func (value *stringValue) Compare(to Value) int {
	return Lt
}

func (value *stringValue) Clone() Value {
	var val stringValue
	val = *value
	return &val
}

func (value *stringValue) String() string {
	if value.valid {
		return value.val
	}
	return nullValueStr
}
