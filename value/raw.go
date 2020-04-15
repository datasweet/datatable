package value

import (
	"fmt"
	"reflect"
)

const RawType = Type("raw")

type rawValue struct {
	vtyp Type
	typ  reflect.Type
	val  interface{}
	null bool
}

// Raw to create a new 'raw' value
func Raw(v interface{}, typ interface{}) Value {
	value := &rawValue{typ: reflect.TypeOf(typ)}
	value.vtyp = TypeFromReflect(value.typ)
	value.Set(v)
	return value
}

// TypeFromReflect to creates a value type from a reflect.Type
func TypeFromReflect(typ reflect.Type) Type {
	if typ == nil {
		return RawType
	}
	return Type(typ.Name())
}

func (value *rawValue) Type() Type {
	return value.vtyp
}

func (value *rawValue) Set(v interface{}) Value {
	value.val = nil
	value.null = true
	if value.typ == nil {
		value.val = v
		value.null = false
	} else if v != nil && reflect.TypeOf(v).ConvertibleTo(value.typ) {
		value.val = v
		value.null = false
	}

	if value.null {
		value.val = nil
	}
	return value
}

func (value *rawValue) Val() interface{} {
	return value.val
}

func (value *rawValue) IsValid() bool {
	return !value.null
}

func (value *rawValue) Compare(to Value) int {
	panic("not implemented")
}

func (value *rawValue) Clone() Value {
	// FIXME: rawValue => no shared memory between value...
	var val rawValue
	val = *value
	return &val
}

func (value *rawValue) String() string {
	if value.null {
		return nullValueStr
	}
	return fmt.Sprintf("%v", value.val)
}
