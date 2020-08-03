package serie

import (
	"fmt"
	"reflect"
	"sort"
)

type Serie interface {
	Type() reflect.Type
	Slice() interface{}     // Underlying slice
	Get(at int) interface{} // T[i]. If T is an interfacer, returns Interfaced value
	All() []interface{}

	// Iterate
	Iterator() Iterator

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
	Distinct() Serie
	Pick(at ...int) Serie
	Where(predicate func(interface{}) bool) Serie
	NonNils() Serie

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

	// Statistics
	Avg(opt ...StatOption) float64
	Count(opt ...StatOption) int64
	CountDistinct(opt ...StatOption) int64
	Cusum(opt ...StatOption) []float64
	Max(opt ...StatOption) float64
	Min(opt ...StatOption) float64
	Median(opt ...StatOption) float64
	Stddev(opt ...StatOption) float64
	Sum(opt ...StatOption) float64
	Variance(opt ...StatOption) float64
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

type serie struct {
	typ        reflect.Type
	slice      reflect.Value
	converter  reflect.Value
	comparer   reflect.Value
	interfacer bool
}

func New(typ interface{}, converter interface{}, comparer interface{}) Serie {
	if typ == nil {
		panic("arg 'typ' is not a concrete type")
	}
	if converter == nil {
		panic("nil converter")
	}
	if comparer == nil {
		panic("nil comparer")
	}

	rv := reflect.ValueOf(typ)
	kind := rv.Kind()

	if kind == reflect.Invalid {
		panic(fmt.Sprintf("type %T is invalid", rv))
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
		panic(fmt.Sprintf("wrong converter signature, must be func(i interface{}) %s", serie.typ.Name()))
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
		panic("wrong comparer signature, must be func(i, j T) int")
	}
	serie.comparer = cmpValue

	// analyse interfacer
	if serie.typ.Implements(reflect.TypeOf((*Interfacer)(nil)).Elem()) {
		serie.interfacer = true
	}

	return serie
}

// Len returns the len of the serie
func (s *serie) Len() int {
	return s.slice.Len()
}

// Type returns the underlying type of serie
func (s *serie) Type() reflect.Type {
	return s.typ
}

// Slice returns the underlying slice
func (s *serie) Slice() interface{} {
	return s.slice.Interface()
}

// Get returns the value at index
// If the serie is an interfacer, ie, values have custom Interface() func,
// the Interface() func will be called.
// So you can have difference between serie.Slice()[at] and serie.Get(at)
func (s *serie) Get(at int) interface{} {
	if s.interfacer {
		return s.slice.Index(at).Interface().(Interfacer).Interface()
	}
	return s.slice.Index(at).Interface()
}

// All to get all values
// <!> Better to use serie.Iterator() if you want to work on values
func (s *serie) All() []interface{} {
	all := make([]interface{}, 0, s.Len())
	for it := s.Iterator(); it.Next(); {
		all = append(all, it.Current())
	}
	return all
}

func (s *serie) String() string {
	return fmt.Sprintf("%+v", s.Slice())
}
