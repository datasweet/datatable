package serie

import (
	"fmt"
	"reflect"
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
		for i, val := range s.values {
			values[i] = val.Clone()
		}
	}

	return &serie{
		typ:     s.typ,
		builder: s.builder,
		values:  values,
	}
}
