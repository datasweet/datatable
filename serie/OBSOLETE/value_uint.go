package serie

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Uint8 = ValueType("uint8")
const Uint16 = ValueType("uint16")
const Uint32 = ValueType("uint32")
const Uint64 = ValueType("uint64")

type uintValue struct {
	bitSize int
	val     uint64
	valid   bool
}

func NewUintValue(v interface{}) Value {
	value := &uintValue{bitSize: wordSize}
	value.Set(v)
	return value
}

func NewUint64Value(v interface{}) Value {
	value := &uintValue{bitSize: 64}
	value.Set(v)
	return value
}

func NewUint32Value(v interface{}) Value {
	value := &uintValue{bitSize: 32}
	value.Set(v)
	return value
}

func NewUint16Value(v interface{}) Value {
	value := &uintValue{bitSize: 16}
	value.Set(v)
	return value
}

func NewUint8Value(v interface{}) Value {
	value := &uintValue{bitSize: 8}
	value.Set(v)
	return value
}

func (value *uintValue) Type() ValueType {
	switch value.bitSize {
	case 8:
		return Uint8
	case 16:
		return Uint16
	case 32:
		return Uint32
	default:
		return Uint64
	}
}

func (value *uintValue) Val() interface{} {
	if value.valid {
		switch value.bitSize {
		case 8:
			return uint8(value.val)
		case 16:
			return uint16(value.val)
		case 32:
			return uint32(value.val)
		default:
			return value.val
		}
	}
	return nil
}

func (value *uintValue) Set(v interface{}) {
	value.val = 0
	value.valid = false

	if casted, ok := cast.AsUint(v); ok {
		value.val = casted
		switch value.bitSize {
		case 8:
			value.valid = uint64(uint8(casted)) == casted
		case 16:
			value.valid = uint64(uint16(casted)) == casted
		case 32:
			value.valid = uint64(uint32(casted)) == casted
		default:
			value.valid = true
		}
	}
}

func (value *uintValue) IsValid() bool {
	return value.valid
}

func (value *uintValue) Clone() Value {
	var val uintValue
	val = *value
	return &val
}

func (value *uintValue) String() string {
	if value.valid {
		return strconv.FormatUint(value.val, 10)
	}
	return nullValueStr
}
