package value

import (
	"github.com/datasweet/cast"
)

const StringType = Type("string")

type stringValue struct {
	val  string
	null bool
}

func String(v ...interface{}) Value {
	value := &stringValue{}
	if len(v) == 1 {
		value.Set(v[0])
	}
	return value
}

func (value *stringValue) Type() Type {
	return StringType
}

func (value *stringValue) Val() interface{} {
	if value.null {
		return nil
	}
	return value.val
}

func (value *stringValue) Set(v interface{}) Value {
	value.val = ""
	value.null = true

	if casted, ok := cast.AsString(v); ok {
		value.val = casted
		value.null = false
	}
	return value
}

func (value *stringValue) IsValid() bool {
	return !value.null
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
	if value.null {
		return nullValueStr
	}
	return value.val
}
