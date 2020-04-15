package serie

import (
	"reflect"

	"github.com/datasweet/datatable/value"
)

// default serie

// Bool to create a boolean serie
func Bool(v ...interface{}) Serie {
	s, _ := New(value.Bool, v...)
	return s
}

// Int to create a int serie
func Int(v ...interface{}) Serie {
	s, _ := New(value.Int, v...)
	return s
}

// Int64 to create a int64 serie
func Int64(v ...interface{}) Serie {
	s, _ := New(value.Int64, v...)
	return s
}

// Int32 to create a int32 serie
func Int32(v ...interface{}) Serie {
	s, _ := New(value.Int32, v...)
	return s
}

// Int16 to create a int16 serie
func Int16(v ...interface{}) Serie {
	s, _ := New(value.Int16, v...)
	return s
}

// Int8 to create a int8 serie
func Int8(v ...interface{}) Serie {
	s, _ := New(value.Int8, v...)
	return s
}

// Uint to create a uint serie
func Uint(v ...interface{}) Serie {
	s, _ := New(value.Uint, v...)
	return s
}

// Uint64 to create a uint64 serie
func Uint64(v ...interface{}) Serie {
	s, _ := New(value.Uint64, v...)
	return s
}

// Uint32 to create a uint serie
func Uint32(v ...interface{}) Serie {
	s, _ := New(value.Uint32, v...)
	return s
}

// Uint16 to create a uint serie
func Uint16(v ...interface{}) Serie {
	s, _ := New(value.Uint16, v...)
	return s
}

// Uint8 to create a uint serie
func Uint8(v ...interface{}) Serie {
	s, _ := New(value.Uint8, v...)
	return s
}

// Float32 to create a float32 serie
func Float32(v ...interface{}) Serie {
	s, _ := New(value.Float32, v...)
	return s
}

// Float64 to create a float64 serie
func Float64(v ...interface{}) Serie {
	s, _ := New(value.Float64, v...)
	return s
}

// String to create a string serie
func String(v ...interface{}) Serie {
	s, _ := New(value.String, v...)
	return s
}

// Raw to create a raw serie (no transformation)
// <!> value.Clone() will share memory...
func Raw(v ...interface{}) Serie {
	builder := func(args ...interface{}) value.Value {
		return value.Raw(args, nil)
	}
	s, _ := New(builder, v...)
	return s
}

// Generic to create a generic serie.
// <!> concretType must not be a ptr.
func Generic(concretType interface{}, v ...interface{}) Serie {
	ctyp := reflect.TypeOf(concretType)
	builder := func(args ...interface{}) value.Value {
		return value.Raw(args, ctyp)
	}
	s, _ := New(builder, v...)
	return s
}

// Time to create a new time serie
func Time(v ...interface{}) Serie {
	s, _ := New(value.Time, v...)
	return s
}
