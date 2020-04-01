package serie

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/datasweet/datatable/value"
)

// Serie describe a serie
type Serie interface {
	Type() value.Type
	Error() error
	Value(at int) interface{}
	Values() []interface{}

	// Mutate
	Append(v ...interface{})
	Prepend(v ...interface{})
	Insert(at int, v ...interface{})
	Update(at int, v interface{})
	Delete(at int)

	// Select
	Head(size int) Serie
	Tail(size int) Serie
	Subset(at, size int) Serie
	Pick(at ...int) Serie
	FindRows(where value.Predicate) []int
	Filter(where value.Predicate) Serie
	Distinct() Serie

	// Clone
	Clone() Serie

	// Sort
	sort.Interface
	SortAsc() Serie
	SortDesc() Serie

	// Print
	Print(opts ...PrintOption) string
	fmt.Stringer
}

func NewBool(v ...interface{}) Serie {
	s := newSerie(value.Bool, value.NewBool)
	s.Append(v...)
	return s
}

func NewInt(v ...interface{}) Serie {
	s := newSerie(value.Int, value.NewInt)
	s.Append(v...)
	return s
}

func NewInt64(v ...interface{}) Serie {
	s := newSerie(value.Int64, value.NewInt64)
	s.Append(v...)
	return s
}

func NewInt32(v ...interface{}) Serie {
	s := newSerie(value.Int32, value.NewInt32)
	s.Append(v...)
	return s
}

func NewInt16(v ...interface{}) Serie {
	s := newSerie(value.Int16, value.NewInt16)
	s.Append(v...)
	return s
}

func NewInt8(v ...interface{}) Serie {
	s := newSerie(value.Int8, value.NewInt8)
	s.Append(v...)
	return s
}

func NewUint(v ...interface{}) Serie {
	s := newSerie(value.Uint, value.NewUint)
	s.Append(v...)
	return s
}

func NewUint64(v ...interface{}) Serie {
	s := newSerie(value.Uint64, value.NewUint64)
	s.Append(v...)
	return s
}

func NewUint32(v ...interface{}) Serie {
	s := newSerie(value.Uint32, value.NewUint32)
	s.Append(v...)
	return s
}

func NewUint16(v ...interface{}) Serie {
	s := newSerie(value.Uint16, value.NewUint16)
	s.Append(v...)
	return s
}

func NewUint8(v ...interface{}) Serie {
	s := newSerie(value.Uint8, value.NewUint8)
	s.Append(v...)
	return s
}

func NewString(v ...interface{}) Serie {
	s := newSerie(value.String, value.NewString)
	s.Append(v...)
	return s
}

func NewRaw(v ...interface{}) Serie {
	builder := func(v interface{}) value.Value {
		return value.NewRaw(v, nil)
	}
	s := newSerie(value.Raw, builder)
	s.Append(v...)
	return s
}

func NewGeneric(concretType interface{}, v ...interface{}) Serie {
	ctyp := reflect.TypeOf(concretType)
	builder := func(v interface{}) value.Value {
		return value.NewRaw(v, ctyp)
	}
	s := newSerie(value.TypeFromReflect(ctyp), builder)
	s.Append(v...)
	return s
}

// Serie implementation
type serie struct {
	typ     value.Type
	builder value.Builder
	values  []value.Value
	err     error
}

func newSerie(typ value.Type, builder value.Builder) Serie {
	return &serie{
		typ:     typ,
		builder: builder,
	}
}

func (s *serie) Type() value.Type {
	return s.typ
}

func (s *serie) Error() error {
	return s.err
}

func (s *serie) Value(at int) interface{} {
	if s.err != nil {
		return nil
	}
	if at < 0 || at >= len(s.values) {
		return nil
	}
	return s.values[at].Val()
}

func (s *serie) Values() []interface{} {
	values := make([]interface{}, 0, len(s.values))
	for _, v := range s.values {
		values = append(values, v.Val())
	}
	return values
}
