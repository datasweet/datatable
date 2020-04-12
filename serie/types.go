package serie

import (
	"reflect"

	"github.com/datasweet/datatable/value"
)

// default serie

// Bool to create a boolean serie
func Bool(v ...interface{}) Serie {
	return New(value.Bool, v...)
}

// Int to create a int serie
func Int(v ...interface{}) Serie {
	return New(value.Int, v...)
}

// Int64 to create a int64 serie
func Int64(v ...interface{}) Serie {
	return New(value.Int64, v...)
}

// Int32 to create a int32 serie
func Int32(v ...interface{}) Serie {
	return New(value.Int32, v...)
}

// Int16 to create a int16 serie
func Int16(v ...interface{}) Serie {
	return New(value.Int16, v...)
}

// Int8 to create a int8 serie
func Int8(v ...interface{}) Serie {
	return New(value.Int8, v...)
}

// Uint to create a uint serie
func Uint(v ...interface{}) Serie {
	return New(value.Uint, v...)
}

// Uint64 to create a uint64 serie
func Uint64(v ...interface{}) Serie {
	return New(value.Uint64, v...)
}

// Uint32 to create a uint serie
func Uint32(v ...interface{}) Serie {
	return New(value.Uint32, v...)
}

// Uint16 to create a uint serie
func Uint16(v ...interface{}) Serie {
	return New(value.Uint16, v...)
}

// Uint8 to create a uint serie
func Uint8(v ...interface{}) Serie {
	return New(value.Uint8, v...)
}

// Float32 to create a float32 serie
func Float32(v ...interface{}) Serie {
	return New(value.Float32, v...)
}

// Float64 to create a float64 serie
func Float64(v ...interface{}) Serie {
	return New(value.Float64, v...)
}

// String to create a string serie
func String(v ...interface{}) Serie {
	return New(value.String, v...)
}

// Raw to create a raw serie (no transformation)
// <!> value.Clone() will share memory...
func Raw(v ...interface{}) Serie {
	builder := func(args ...interface{}) value.Value {
		return value.Raw(args, nil)
	}
	return New(builder, v...)
}

// Generic to create a generic serie.
// <!> concretType must not be a ptr.
func Generic(concretType interface{}, v ...interface{}) Serie {
	ctyp := reflect.TypeOf(concretType)
	builder := func(args ...interface{}) value.Value {
		return value.Raw(args, ctyp)
	}
	return New(builder, v...)
}

// Time to create a new time serie
func Time(v ...interface{}) Serie {
	return New(value.Time, v...)
}
