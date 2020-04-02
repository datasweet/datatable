package value

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Uint8Type = Type("uint8")
const Uint16Type = Type("uint16")
const Uint32Type = Type("uint32")
const Uint64Type = Type("uint64")
const UintType = Type("uint")

type uintValue struct {
	bitSize int
	val     uint64
	null    bool
}

func newUint(bitSize int, v ...interface{}) Value {
	value := &uintValue{bitSize: bitSize}
	if len(v) == 1 {
		value.Set(v[0])
	}
	return value
}

// Uint to create a new uint value
// Uint() will create the default value, ie "0"
// Uint(v) will parse the v value as uint, returns nil if error
func Uint(v ...interface{}) Value {
	return newUint(0, v...)
}

// Uint64 to create a new uint64 value
// Uint64() will create the default value, ie "0"
// Uint64(v) will parse the v value as uint64, returns nil if error
func Uint64(v ...interface{}) Value {
	return newUint(64, v...)
}

// Uint32 to create a new uint32 value
// Uint32() will create the default value, ie "0"
// Uint32(v) will parse the v value as uint32, returns nil if error
func Uint32(v ...interface{}) Value {
	return newUint(32, v...)
}

// Uint16 to create a new uint16 value
// Uint16() will create the default value, ie "0"
// Uint16(v) will parse the v value as uint16, returns nil if error
func Uint16(v ...interface{}) Value {
	return newUint(16, v...)
}

// Uint8 to create a new uint8 value
// Uint8() will create the default value, ie "0"
// Uint8(v) will parse the v value as uint8, returns nil if error
func Uint8(v ...interface{}) Value {
	return newUint(8, v...)
}

func (value *uintValue) Type() Type {
	switch value.bitSize {
	default:
		return UintType
	case 64:
		return Uint64Type
	case 32:
		return Uint32Type
	case 16:
		return Uint16Type
	case 8:
		return Uint8Type
	}
}

func (value *uintValue) Val() interface{} {
	if value.null {
		return nil
	}
	switch value.bitSize {
	default:
		return uint(value.val)
	case 64:
		return value.val
	case 32:
		return uint32(value.val)
	case 16:
		return uint16(value.val)
	case 8:
		return uint8(value.val)
	}
}

func (value *uintValue) Set(v interface{}) Value {
	value.val = 0
	value.null = true

	if casted, ok := cast.AsUint(v); ok {
		switch value.bitSize {
		default:
			if uint64(uint(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 64:
			value.val = casted
			value.null = false
		case 32:
			if uint64(uint32(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 16:
			if uint64(uint16(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 8:
			if uint64(uint8(casted)) == casted {
				value.val = casted
				value.null = false
			}
		}
	}
	return value
}

func (value *uintValue) IsValid() bool {
	return !value.null
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a int < value
func (value *uintValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	iv, ok := to.(*uintValue)
	if !ok {
		// try to convert
		iv = Uint64(to.Val()).(*uintValue)
	}

	if iv.null {
		if value.null {
			return Eq
		}
		return Gt
	}

	if value.val == iv.val {
		return Eq
	}

	if value.val > iv.val {
		return Gt
	}

	return Lt
}

func (value *uintValue) Clone() Value {
	var cpy uintValue
	cpy = *value
	return &cpy
}

func (value *uintValue) String() string {
	if value.null {
		return nullValueStr
	}
	return strconv.FormatUint(value.val, 10)
}
