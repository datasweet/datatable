package value

import (
	"github.com/datasweet/cast"
)

const BoolType = Type("bool")
const trueStr = "true"
const falseStr = "false"

type boolValue struct {
	val  bool
	null bool
}

// Bool to create a new bool value
// Bool() will create the default value, ie "false"
// Bool(v) will parse the v value as bool
func Bool(v ...interface{}) Value {
	bv := &boolValue{}
	if len(v) == 1 {
		bv.Set(v[0])
	}
	return bv
}

func (value *boolValue) Type() Type {
	return BoolType
}

func (value *boolValue) Val() interface{} {
	if value.null {
		return nil
	}
	return value.val
}

func (value *boolValue) Set(v interface{}) Value {
	value.val = false
	value.null = true
	if v == nil {
		return value
	}
	switch val := v.(type) {
	case *boolValue:
		value.val = val.val
		value.null = false
	default:
		if casted, ok := cast.AsBool(v); ok {
			value.val = casted
			value.null = false
		}
	}
	return value
}

func (value *boolValue) IsValid() bool {
	return !value.null
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a Bool < false < true
func (value *boolValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	bv, ok := to.(*boolValue)
	if !ok {
		// try to convert
		bv = Bool(to.Val()).(*boolValue)
	}

	if bv.null {
		if value.null {
			return Eq
		}
		return Gt
	}

	if value.val == bv.val {
		return Eq
	}

	if value.val {
		return Gt
	}

	return Lt
}

func (value *boolValue) Clone() Value {
	return &boolValue{
		val:  value.val,
		null: value.null,
	}
}

func (value *boolValue) String() string {
	if value.null {
		return nullValueStr
	}
	if value.val {
		return trueStr
	}
	return falseStr
}
