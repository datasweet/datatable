package value

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Int8 = Type("int8")
const Int16 = Type("int16")
const Int32 = Type("int32")
const Int64 = Type("int64")
const Int = Type("int")

type intValue struct {
	bitSize int
	val     *int64
}

func NewInt(v interface{}) Value {
	value := &intValue{}
	value.Set(v)
	return value
}

func NewInt64(v interface{}) Value {
	value := &intValue{bitSize: 64}
	value.Set(v)
	return value
}

func NewInt32(v interface{}) Value {
	value := &intValue{bitSize: 32}
	value.Set(v)
	return value
}

func NewInt16(v interface{}) Value {
	value := &intValue{bitSize: 16}
	value.Set(v)
	return value
}

func NewInt8(v interface{}) Value {
	value := &intValue{bitSize: 8}
	value.Set(v)
	return value
}

func (value *intValue) Type() Type {
	switch value.bitSize {
	default:
		return Int
	case 64:
		return Int64
	case 32:
		return Int32
	case 16:
		return Int16
	case 8:
		return Int8
	}
}

func (value *intValue) Val() interface{} {
	if value.val == nil {
		return nil
	}
	switch value.bitSize {
	default:
		return int(*value.val)
	case 64:
		return *value.val
	case 32:
		return int32(*value.val)
	case 16:
		return int16(*value.val)
	case 8:
		return int8(*value.val)
	}
}

func (value *intValue) Set(v interface{}) {
	value.val = nil

	if casted, ok := cast.AsInt(v); ok {
		switch value.bitSize {
		default:
			if int64(int(casted)) == casted {
				value.val = &casted
			}
		case 64:
			value.val = &casted
		case 32:
			if int64(int32(casted)) == casted {
				value.val = &casted
			}
		case 16:
			if int64(int16(casted)) == casted {
				value.val = &casted
			}
		case 8:
			if int64(int8(casted)) == casted {
				value.val = &casted
			}
		}
	}
}

func (value *intValue) IsValid() bool {
	return value.val != nil
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
		iv = NewInt64(to.Val()).(*intValue)
	}

	if iv.val == nil {
		if value.val == nil {
			return Eq
		}
		return Gt
	}

	a, b := *value.val, *iv.val

	if a == b {
		return Eq
	}

	if a > b {
		return Gt
	}

	return Lt
}

func (value *intValue) Clone() Value {
	var cpy *int64
	if value.val != nil {
		cpy = new(int64)
		*cpy = *value.val
	}
	return &intValue{
		val:     cpy,
		bitSize: value.bitSize,
	}
}

func (value *intValue) String() string {
	if value.val == nil {
		return nullValueStr
	}
	return strconv.FormatInt(*value.val, 10)
}
