package value

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Int8Type = Type("int8")
const Int16Type = Type("int16")
const Int32Type = Type("int32")
const Int64Type = Type("int64")
const IntType = Type("int")

type intValue struct {
	bitSize int
	val     int64
	null    bool
}

func newInt(bitSize int, v ...interface{}) Value {
	value := &intValue{bitSize: bitSize}
	if len(v) == 1 {
		value.Set(v[0])
	}
	return value
}

// Int to create a new int value
// Int() will create the default value, ie "0"
// Int(v) will parse the v value as int, returns nil if error
func Int(v ...interface{}) Value {
	return newInt(0, v...)
}

// Int64 to create a new int64 value
// Int64() will create the default value, ie "0"
// Int64(v) will parse the v value as int64, returns nil if error
func Int64(v ...interface{}) Value {
	return newInt(64, v...)
}

// Int32 to create a new int32 value
// Int32() will create the default value, ie "0"
// Int32(v) will parse the v value as int32, returns nil if error
func Int32(v ...interface{}) Value {
	return newInt(32, v...)
}

// Int16 to create a new int16 value
// Int16() will create the default value, ie "0"
// Int16(v) will parse the v value as int16, returns nil if error
func Int16(v ...interface{}) Value {
	return newInt(16, v...)
}

// Int8 to create a new int8 value
// Int8() will create the default value, ie "0"
// Int8(v) will parse the v value as int8, returns nil if error
func Int8(v ...interface{}) Value {
	return newInt(8, v...)
}

func (value *intValue) Type() Type {
	switch value.bitSize {
	default:
		return IntType
	case 64:
		return Int64Type
	case 32:
		return Int32Type
	case 16:
		return Int16Type
	case 8:
		return Int8Type
	}
}

func (value *intValue) Val() interface{} {
	if value.null {
		return nil
	}
	switch value.bitSize {
	default:
		return int(value.val)
	case 64:
		return value.val
	case 32:
		return int32(value.val)
	case 16:
		return int16(value.val)
	case 8:
		return int8(value.val)
	}
}

func (value *intValue) Set(v interface{}) Value {
	value.val = 0
	value.null = true

	if casted, ok := cast.AsInt(v); ok {
		switch value.bitSize {
		default:
			if int64(int(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 64:
			value.val = casted
			value.null = false
		case 32:
			if int64(int32(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 16:
			if int64(int16(casted)) == casted {
				value.val = casted
				value.null = false
			}
		case 8:
			if int64(int8(casted)) == casted {
				value.val = casted
				value.null = false
			}
		}
	}
	return value
}

func (value *intValue) IsValid() bool {
	return !value.null
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a int < value
func (value *intValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	iv, ok := to.(*intValue)
	if !ok {
		// try to convert
		iv = Int64(to.Val()).(*intValue)
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

func (value *intValue) Clone() Value {
	var cpy intValue
	cpy = *value
	return &cpy
}

func (value *intValue) String() string {
	if value.null {
		return nullValueStr
	}
	return strconv.FormatInt(value.val, 10)
}
