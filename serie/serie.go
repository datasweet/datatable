package serie

import (
	"reflect"
	"sort"

	"github.com/pkg/errors"
)

type Serie interface {
	Type() reflect.Type
	All() []interface{}
	Get(at int) interface{}
	Slice() interface{}

	// Mutate
	Append(v ...interface{})
	Prepend(v ...interface{}) error
	Insert(at int, v ...interface{}) error
	Set(at int, v interface{}) error
	Delete(at int) error
	Grow(size int) error
	Shrink(size int) error
	Concat(serie ...Serie) error
	Clear()

	// Select
	Head(size int) Serie
	Tail(size int) Serie
	Subset(at, size int) Serie
	Filter(where interface{}) (Serie, error)
	Distinct() Serie

	// Copy
	EmptyCopy() Serie
	Copy() Serie

	// Sort
	sort.Interface
	Compare(i, j int) int
	SortAsc()
	SortDesc()

	// // Print
	// Print(opts ...PrintOption) string
	// fmt.Stringer
}

type serie struct {
	typ        reflect.Type
	slice      reflect.Value
	converter  reflect.Value
	comparer   reflect.Value
	interfacer bool
}

// Interfacer to convert a value of serie to interface{}
// Used with serie.Get(at) serie.All()
type Interfacer interface {
	Interface() interface{}
}

const (
	Lt = -1
	Eq = 0
	Gt = 1
)

func New(typ interface{}, converter interface{}, comparer interface{}) (Serie, error) {
	if typ == nil {
		return nil, errors.New("not a concrete type")
	}
	if converter == nil {
		return nil, errors.New("nil converter")
	}
	if comparer == nil {
		return nil, errors.New("nil comparer")
	}

	rv := reflect.ValueOf(typ)
	kind := rv.Kind()

	if kind == reflect.Invalid {
		return nil, errors.Errorf("type %T is invalid", rv)
	}

	serie := &serie{}

	if kind == reflect.Slice {
		serie.slice = rv
		serie.typ = rv.Type().Elem()
	} else {
		serie.typ = rv.Type()
		serie.slice = reflect.MakeSlice(reflect.SliceOf(serie.typ), 0, 0)
	}

	// analyse converter
	convValue := reflect.ValueOf(converter)
	convType := convValue.Type()
	if convType.Kind() != reflect.Func ||
		convType.NumIn() != 1 ||
		convType.NumOut() != 1 ||
		convType.In(0).Kind() != reflect.Interface ||
		convType.Out(0) != serie.typ {
		return nil, errors.Errorf("wrong converter signature, must be func(i interface{}) %s", serie.typ.Name())
	}
	serie.converter = convValue

	// analyse comparer
	cmpValue := reflect.ValueOf(comparer)
	cmpType := cmpValue.Type()
	if cmpType.Kind() != reflect.Func ||
		cmpType.NumIn() != 2 ||
		cmpType.NumOut() != 1 ||
		cmpType.In(0) != serie.typ ||
		cmpType.In(1) != serie.typ ||
		cmpType.Out(0).Kind() != reflect.Int {
		return nil, errors.New("wrong comparer signature, must be func(i, j T) int")
	}
	serie.comparer = cmpValue

	// analyse interfacer
	if serie.typ.Implements(reflect.TypeOf((*Interfacer)(nil)).Elem()) {
		serie.interfacer = true
	}

	return serie, nil
}

// Len returns the len of the serie
func (s *serie) Len() int {
	return s.slice.Len()
}

// Type returns the underlying type of serie
func (s *serie) Type() reflect.Type {
	return s.typ
}

// All returns all values in serie
func (s *serie) All() []interface{} {
	cnt := s.slice.Len()
	values := make([]interface{}, 0, cnt)
	if s.interfacer {
		for i := 0; i < cnt; i++ {
			values = append(values, s.slice.Index(i).Interface().(Interfacer).Interface())
		}
	} else {
		for i := 0; i < cnt; i++ {
			values = append(values, s.slice.Index(i).Interface())
		}
	}
	return values
}

// Get returns the value at index
func (s *serie) Get(at int) interface{} {
	if s.interfacer {
		return s.slice.Index(at).Interface().(Interfacer).Interface()
	}
	return s.slice.Index(at).Interface()
}

func (s *serie) Slice() interface{} {
	return s.slice.Interface()
}
