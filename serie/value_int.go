package serie

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Int8 = ValueType("int8")
const Int16 = ValueType("int16")
const Int32 = ValueType("int32")
const Int64 = ValueType("int64")

type intValue struct {
	bitSize int
	val     int64
	valid   bool
}

func NewIntValue(v interface{}) Value {
	value := &intValue{bitSize: wordSize}
	value.Set(v)
	return value
}

func NewInt64Value(v interface{}) Value {
	value := &intValue{bitSize: 64}
	value.Set(v)
	return value
}

func NewInt32Value(v interface{}) Value {
	value := &intValue{bitSize: 32}
	value.Set(v)
	return value
}

func NewInt16Value(v interface{}) Value {
	value := &intValue{bitSize: 16}
	value.Set(v)
	return value
}

func NewInt8Value(v interface{}) Value {
	value := &intValue{bitSize: 8}
	value.Set(v)
	return value
}

func (value *intValue) Type() ValueType {
	switch value.bitSize {
	case 8:
		return Int8
	case 16:
		return Int16
	case 32:
		return Int32
	default:
		return Int64
	}
}

func (value *intValue) Val() interface{} {
	if value.valid {
		switch value.bitSize {
		case 8:
			return int8(value.val)
		case 16:
			return int16(value.val)
		case 32:
			return int32(value.val)
		default:
			return value.val
		}
	}
	return nil
}

func (value *intValue) Set(v interface{}) {
	value.val = 0
	value.valid = false

	if casted, ok := cast.AsInt(v); ok {
		value.val = casted
		switch value.bitSize {
		case 8:
			value.valid = int64(int8(casted)) == casted
		case 16:
			value.valid = int64(int16(casted)) == casted
		case 32:
			value.valid = int64(int32(casted)) == casted
		default:
			value.valid = true
		}
	}
}

func (value *intValue) IsValid() bool {
	return value.valid
}

func (value *intValue) Clone() Value {
	var val intValue
	val = *value
	return &val
}

func (value *intValue) String() string {
	if value.valid {
		return strconv.FormatInt(value.val, 10)
	}
	return nullValueStr
}
