package serie

import (
	"fmt"
	"reflect"
)

const Raw = ValueType("raw")

type rawValue struct {
	vtyp  ValueType
	typ   reflect.Type
	val   interface{}
	valid bool
}

// NewRawValue to create a new 'raw' value
func NewRawValue(v interface{}, typ interface{}) Value {
	value := &rawValue{typ: reflect.TypeOf(typ)}
	value.vtyp = valueTypeFromReflect(value.typ)
	value.Set(v)
	return value
}

func (value *rawValue) Type() ValueType {
	return value.vtyp
}

func (value *rawValue) Set(v interface{}) {
	value.val = v
	value.valid = false
	if value.typ == nil {
		value.valid = true
	} else if v != nil {
		value.valid = reflect.TypeOf(v).ConvertibleTo(value.typ)
	}

	if !value.valid {
		value.val = nil
	}
}

func (value *rawValue) Val() interface{} {
	return value.val
}

func (value *rawValue) IsValid() bool {
	return value.valid
}

func (value *rawValue) Clone() Value {
	// FIXME: rawValue => no shared memory between value...
	var val rawValue
	val = *value
	return &val
}

func (value *rawValue) String() string {
	if value.valid {
		return fmt.Sprintf("%v", value.val)
	}
	return nullValueStr
}
