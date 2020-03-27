package serie

import (
	"github.com/datasweet/cast"
)

const Bool = ValueType("bool")
const trueStr = "true"
const falseStr = "false"

type boolValue struct {
	val   bool
	valid bool
}

func NewBoolValue(v interface{}) Value {
	value := &boolValue{}
	value.Set(v)
	return value
}

func (value *boolValue) Type() ValueType {
	return Bool
}

func (value *boolValue) Val() interface{} {
	if value.valid {
		return value.val
	}
	return nil
}

func (value *boolValue) Set(v interface{}) {
	value.val = false
	value.valid = false

	if casted, ok := cast.AsBool(v); ok {
		value.val = casted
		value.valid = true
	}
}

func (value *boolValue) IsValid() bool {
	return value.valid
}

func (value *boolValue) Clone() Value {
	var val boolValue
	val = *value
	return &val
}

func (value *boolValue) String() string {
	if value.valid {
		if value.val {
			return trueStr
		}
		return falseStr
	}
	return nullValueStr
}
