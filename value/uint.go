package value

import (
	"strconv"

	"github.com/datasweet/cast"
)

const Uint8 = Type("uint8")
const Uint16 = Type("uint16")
const Uint32 = Type("uint32")
const Uint64 = Type("uint64")
const Uint = Type("uint")

type uintValue struct {
	bitSize int
	val     *uint64
}

func NewUint(v interface{}) Value {
	value := &uintValue{}
	value.Set(v)
	return value
}

func NewUint64(v interface{}) Value {
	value := &uintValue{bitSize: 64}
	value.Set(v)
	return value
}

func NewUint32(v interface{}) Value {
	value := &uintValue{bitSize: 32}
	value.Set(v)
	return value
}

func NewUint16(v interface{}) Value {
	value := &uintValue{bitSize: 16}
	value.Set(v)
	return value
}

func NewUint8(v interface{}) Value {
	value := &uintValue{bitSize: 8}
	value.Set(v)
	return value
}

func (value *uintValue) Type() Type {
	switch value.bitSize {
	default:
		return Uint
	case 64:
		return Uint64
	case 32:
		return Uint32
	case 16:
		return Uint16
	case 8:
		return Uint8
	}
}

func (value *uintValue) Val() interface{} {
	if value.val == nil {
		return nil
	}
	switch value.bitSize {
	default:
		return uint(*value.val)
	case 64:
		return *value.val
	case 32:
		return uint32(*value.val)
	case 16:
		return uint16(*value.val)
	case 8:
		return uint8(*value.val)
	}
}

func (value *uintValue) Set(v interface{}) {
	value.val = nil

	if casted, ok := cast.AsUint(v); ok {
		switch value.bitSize {
		default:
			if uint64(uint(casted)) == casted {
				value.val = &casted
			}
		case 64:
			value.val = &casted
		case 32:
			if uint64(uint32(casted)) == casted {
				value.val = &casted
			}
		case 16:
			if uint64(uint16(casted)) == casted {
				value.val = &casted
			}
		case 8:
			if uint64(uint8(casted)) == casted {
				value.val = &casted
			}
		}
	}
}

func (value *uintValue) IsValid() bool {
	return value.val != nil
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a uint < value
func (value *uintValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	iv, ok := to.(*uintValue)
	if !ok {
		// try to convert
		iv = NewUint64(to.Val()).(*uintValue)
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

func (value *uintValue) Clone() Value {
	var cpy *uint64
	if value.val != nil {
		cpy = new(uint64)
		*cpy = *value.val
	}
	return &uintValue{
		val:     cpy,
		bitSize: value.bitSize,
	}
}

func (value *uintValue) String() string {
	if value.val == nil {
		return nullValueStr
	}
	return strconv.FormatUint(*value.val, 10)
}
