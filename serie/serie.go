package serie

import (
	"fmt"
	"reflect"
)

type ctr uint8

const (
	unknownCtr = iota
	scalarCtr
	sliceCtr
	mapCtr
)

// Serie describe a serie
type Serie interface {
	Type() ValueType
	Len() int
	Values() []interface{}

	// Mutate
	Append(v ...interface{})
	Prepend(v ...interface{})
	Insert(at int, v ...interface{}) error

	// Select
	Head(size int) Serie
	Tail(size int) Serie
	Subset(at, size int) Serie
	Pick(at ...int) Serie
	FindRows(where ValuePredicate) []int
	Filter(where ValuePredicate) Serie

	// Print
	Print(opts ...PrintOption) string
	fmt.Stringer
}

// Serie implementation
type serie struct {
	typ     ValueType
	builder ValueBuilder
	values  []Value
}

func NewBool(v ...interface{}) Serie {
	s := newSerie(Bool, NewBoolValue)
	s.Append(v...)
	return s
}

func NewInt64(v ...interface{}) Serie {
	s := newSerie(Int64, NewInt64Value)
	s.Append(v...)
	return s
}

func NewInt32(v ...interface{}) Serie {
	s := newSerie(Int32, NewInt32Value)
	s.Append(v...)
	return s
}

func NewInt16(v ...interface{}) Serie {
	s := newSerie(Int16, NewInt16Value)
	s.Append(v...)
	return s
}

func NewInt8(v ...interface{}) Serie {
	s := newSerie(Int8, NewInt8Value)
	s.Append(v...)
	return s
}

func NewUint64(v ...interface{}) Serie {
	s := newSerie(Uint64, NewUint64Value)
	s.Append(v...)
	return s
}

func NewUint32(v ...interface{}) Serie {
	s := newSerie(Uint32, NewUint32Value)
	s.Append(v...)
	return s
}

func NewUint16(v ...interface{}) Serie {
	s := newSerie(Uint16, NewUint16Value)
	s.Append(v...)
	return s
}

func NewUint8(v ...interface{}) Serie {
	s := newSerie(Uint8, NewUint8Value)
	s.Append(v...)
	return s
}

func NewString(v ...interface{}) Serie {
	s := newSerie(String, NewStringValue)
	s.Append(v...)
	return s
}

func NewGeneric(concretType interface{}, v ...interface{}) Serie {
	ctyp := reflect.TypeOf(concretType)
	builder := func(v interface{}) Value {
		return NewRawValue(v, ctyp)
	}
	s := newSerie(valueTypeFromReflect(ctyp), builder)
	s.Append(v...)
	return s
}

func newSerie(typ ValueType, builder ValueBuilder) Serie {
	return &serie{
		typ:     typ,
		builder: builder,
	}
}

func (s *serie) Type() ValueType {
	return s.typ
}

func (s *serie) Len() int {
	return len(s.values)
}

func (s *serie) Values() []interface{} {
	values := make([]interface{}, 0, len(s.values))
	for i, v := range s.values {
		values[i] = v.Val()
	}
	return values
}

func (s *serie) clone(withValues bool) *serie {
	var values []Value

	if withValues {
		values := make([]Value, len(s.values))
		copy(values, s.values)
	}

	return &serie{
		typ:     s.typ,
		builder: s.builder,
		values:  values,
	}
}

/*
switch V := val.(type) {
	case []float64:
		// count how many NaN
		for _, v := range V {
			if isNaN(v) {
				s.nilCount++
			}
		}
		s.Values = append(s.Values[:row], append(V, s.Values[row:]...)...)
		return
	}

	s.Values = append(s.Values, nan())
	copy(s.Values[row+1:], s.Values[row:])

	v := s.valToPointer(val)
	if isNaN(v) {
		s.nilCount++
	}

	s.Values[row] = v
*/

// func (s *serie) InsertAt(at, v ...interface{}) {

// }

// New to create a new serie
// func New(v interface{}, opts ...Option) *Serie {
// 	serie := new(Serie)

// 	// default options
// 	options := Options{
// 		Type:       Raw,
// 		DetectType: true,
// 		Len:        -1,
// 	}

// 	// apply options
// 	for _, o := range opts {
// 		o(&options)
// 	}

// 	serie.Name = options.Name
// 	serie.typ = options.Type

// 	// analyze
// 	val := reflect.ValueOf(v)
// 	kind := val.Kind()

// 	switch kind {
// 	case reflect.Bool:
// 		if options.DetectType {
// 			serie.typ = Bool
// 		}
// 		serie.fromScalar(v, options.Len)
// 		return serie

// 	case reflect.Int:
// 		if options.DetectType {
// 			if wordSize == 32 {
// 				serie.typ = Int32
// 			} else {
// 				serie.typ = Int64
// 			}
// 			serie.fromScalar(v, options.Len)
// 		}
// 	case reflect.Int8:
// 		if options.DetectType {
// 			serie.typ = Int8
// 		}
// 		serie.fromScalar(v, options.Len)

// 	case reflect.Int16:
// 		if options.DetectType {
// 			serie.typ = Int16
// 		}
// 		serie.fromScalar(v, options.Len)

// 	case reflect.Int32:
// 		if options.DetectType {
// 			serie.typ = Int32
// 		}
// 		serie.fromScalar(v, options.Len)

// 	case reflect.Int64:
// 		if options.DetectType {
// 			serie.typ = Int64
// 		}
// 		serie.fromScalar(v, options.Len)

// 	// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
// 	// 	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64: // FIXME overflow detection
// 	// 	if options.DetectType {
// 	// 		serie.typ = Int
// 	// 	}
// 	// 	serie.fromScalar(v, options.Len)
// 	// 	return serie

// 	// case reflect.Float32, reflect.Float64:
// 	// 	if options.DetectType {
// 	// 		serie.typ = String
// 	// 	}
// 	// 	serie.fromScalar(v, options.Len)
// 	// 	return serie

// 	// case reflect.Complex64, reflect.Complex128:
// 	// 	if options.DetectType {
// 	// 		serie.typ = Raw
// 	// 	}
// 	// 	serie.fromScalar(v, options.Len)
// 	// 	return serie

// 	// case reflect.Array, reflect.Slice:
// 	// 	len := val.Len()
// 	// 	n := options.Len

// 	// 	if options.Len < 0 {
// 	// 		n = len
// 	// 	}

// 	// 	slice := make([]interface{}, n)
// 	// 	ctyp := options.Type
// 	// 	detect := options.DetectType

// 	// 	checkType := func(prev, new ValueType) (ValueType, bool) {
// 	// 		if prev == Raw {
// 	// 			return new, true
// 	// 		}
// 	// 		if prev == new {
// 	// 			return new, true
// 	// 		}
// 	// 		return Raw, false
// 	// 	}

// 	// 	for i := 0; i < n; i++ {
// 	// 		if i >= len {
// 	// 			slice[i] = nil
// 	// 		} else {
// 	// 			rv := val.Index(i)
// 	// 			slice[i] = rv.Interface()

// 	// 			if detect {
// 	// 				switch knd := rv.Kind(); knd {
// 	// 				case reflect.Bool:
// 	// 					ctyp, detect = checkType(ctyp, Bool)
// 	// 				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 	// 					ctyp, detect = checkType(ctyp, Int)
// 	// 				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 	// 					ctyp, detect = checkType(ctyp, Int)
// 	// 				case reflect.Float32, reflect.Float64:
// 	// 					ctyp, detect = checkType(ctyp, Float)
// 	// 				case reflect.String:
// 	// 					ctyp, detect = checkType(ctyp, String)
// 	// 				default:
// 	// 					if slice[i] != nil {
// 	// 						ctyp = Raw
// 	// 						detect = false
// 	// 					}
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}
// 	// 	serie.typ = ctyp
// 	// 	serie.fromSlice(slice)
// 	// 	return serie

// 	default:
// 		return nil

// 		//	Invalid Kind = iota
// 		// Bool
// 		// Int
// 		// Int8
// 		// Int16
// 		// Int32
// 		// Int64
// 		// Uint
// 		// Uint8
// 		// Uint16
// 		// Uint32
// 		// Uint64
// 		// Uintptr
// 		// Float32
// 		// Float64
// 		// Complex64
// 		// Complex128
// 		// Array
// 		// Chan
// 		// Func
// 		// Interface
// 		// Map
// 		// Ptr
// 		// Slice
// 		// String
// 		// Struct
// 		// UnsafePointer

// 		// }
// 	}
// }

// func (serie *Serie) fromScalar(v interface{}, n int) {
// 	l := n
// 	if n < 0 {
// 		l = 1
// 	}
// 	serie.indexes = make([]int, 0, l)
// 	serie.values = make(map[int]Value, l)

// 	builder := valueFactory[serie.typ]
// 	// if builder == nil {
// 	// 	builder = NewValue
// 	// }

// 	for i := 1; i <= l; i++ {
// 		serie.indexes = append(serie.indexes, i)
// 		serie.values[i] = builder(v)
// 	}
// }

// func (serie *Serie) fromSlice(v []interface{}) {
// 	l := len(v)

// 	serie.indexes = make([]int, 0, l)
// 	serie.values = make(map[int]*Value, l)

// 	for i := 1; i <= l; i++ {
// 		serie.indexes = append(serie.indexes, i)
// 		serie.values[i] = NewValue(v[i-1], serie.typ)
// 	}
// }
