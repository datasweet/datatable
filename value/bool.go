package value

import (
	"github.com/datasweet/cast"
)

const Bool = Type("bool")
const trueStr = "true"
const falseStr = "false"

type boolValue struct {
	val *bool
}

func NewBool(v interface{}) Value {
	value := &boolValue{}
	value.Set(v)
	return value
}

func (value *boolValue) Type() Type {
	return Bool
}

func (value *boolValue) Val() interface{} {
	if value.val == nil {
		return nil
	}
	return *value.val
}

func (value *boolValue) Set(v interface{}) {
	value.val = nil
	if v == nil {
		return
	}
	switch val := v.(type) {
	case *boolValue:
		value.val = val.val
	default:
		if casted, ok := cast.AsBool(v); ok {
			value.val = &casted
		}
	}
}

func (value *boolValue) IsValid() bool {
	return value.val != nil
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
		bv = NewBool(to.Val()).(*boolValue)
	}

	if bv.val == nil {
		if value.val == nil {
			return Eq
		}
		return Gt
	}

	a, b := *value.val, *bv.val

	if a == b {
		return Eq
	}

	if a {
		return Gt
	}

	return Lt
}

func (value *boolValue) Clone() Value {
	var cpy *bool
	if value.val != nil {
		cpy = new(bool)
		*cpy = *value.val
	}
	return &boolValue{
		val: cpy,
	}
}

func (value *boolValue) String() string {
	if value.val == nil {
		return nullValueStr
	}
	if *value.val {
		return trueStr
	}
	return falseStr
}
