package value

import (
	"fmt"
	"reflect"
)

const Raw = Type("raw")

type rawValue struct {
	vtyp  Type
	typ   reflect.Type
	val   interface{}
	valid bool
}

// NewRaw to create a new 'raw' value
func NewRaw(v interface{}, typ interface{}) Value {
	value := &rawValue{typ: reflect.TypeOf(typ)}
	value.vtyp = TypeFromReflect(value.typ)
	value.Set(v)
	return value
}

func (value *rawValue) Type() Type {
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

func (value *rawValue) Compare(to Value) int {
	return Lt
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
