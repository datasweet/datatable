package value

import (
	"math"
	"strconv"

	"github.com/datasweet/cast"
)

const Float32Type = Type("float32")
const Float64Type = Type("float64")

type floatValue struct {
	bitSize int
	val     float64
	null    bool
}

func newFloat(bitSize int, v ...interface{}) Value {
	value := &floatValue{bitSize: bitSize}
	if len(v) == 1 {
		value.Set(v[0])
	}
	return value
}

// Float64 to create a new float64 value
// Float64() will create the default value, ie "0.0"
// Float64(v) will parse the v value as float, returns nil if error
func Float64(v ...interface{}) Value {
	return newFloat(64, v...)
}

// Float32 to create a new float32 value
// Float32() will create the default value, ie "0.0"
// Float32(v) will parse the v value as float32, returns nil if error
func Float32(v ...interface{}) Value {
	return newFloat(32, v...)
}

func (value *floatValue) Type() Type {
	switch value.bitSize {
	default:
		return Float64Type
	case 32:
		return Float32Type
	}
}

func (value *floatValue) Val() interface{} {
	if value.null {
		return nil
	}
	switch value.bitSize {
	default:
		return value.val
	case 32:
		return float32(value.val)
	}
}

func (value *floatValue) Set(v interface{}) Value {
	value.val = 0.0
	value.null = true

	if casted, ok := cast.AsFloat(v); ok {
		switch value.bitSize {
		default:
			value.val = casted
			value.null = false
		case 32:
			if casted > -math.MaxFloat32 && casted < math.MaxFloat32 {
				value.val = casted
				value.null = false
			}
		}
	}
	return value
}

func (value *floatValue) IsValid() bool {
	return !value.null
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a int < value
func (value *floatValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	iv, ok := to.(*floatValue)
	if !ok {
		// try to convert
		iv = Float64(to.Val()).(*floatValue)
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

func (value *floatValue) Clone() Value {
	var cpy floatValue
	cpy = *value
	return &cpy
}

func (value *floatValue) String() string {
	if value.null {
		return nullValueStr
	}
	return strconv.FormatFloat(value.val, 'f', -1, value.bitSize)
}
