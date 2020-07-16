package serie

import (
	"reflect"

	"github.com/pkg/errors"
)

// Head returns the first {size} rows of the serie
func (s *serie) Head(size int) Serie {
	return s.Subset(0, size)
}

// Head returns the last {size} rows of the serie
func (s *serie) Tail(size int) Serie {
	return s.Subset(s.Len()-size, size)
}

// Subset returns the a subset {at} index and with {size}
func (s *serie) Subset(at, size int) Serie {
	cpy := s.EmptyCopy().(*serie)
	to := at + size
	if at >= 0 && size > 0 && to <= s.Len() {
		cpy.slice = s.slice.Slice(at, to)
	}
	return cpy
}

// Filter the series with a predicate
// Predicate must be func(T) bool
func (s *serie) Filter(predicate interface{}) (Serie, error) {
	// Check predicate
	// must be func(T) bool

	if predicate == nil {
		return s.EmptyCopy(), errors.New("no predicate")
	}

	pv := reflect.ValueOf(predicate)
	pt := pv.Type()
	if pt.Kind() != reflect.Func ||
		pt.NumIn() != 1 ||
		pt.NumOut() != 1 ||
		pt.In(0) != s.typ ||
		pt.Out(0).Kind() != reflect.Bool {
		return s.EmptyCopy(), errors.New("wrong converter signature, must be func(T) bool")
	}

	cnt := s.Len()

	cpy := &serie{
		typ:       s.typ,
		converter: s.converter,
		slice:     reflect.MakeSlice(reflect.SliceOf(s.typ), 0, cnt),
	}

	for i := 0; i < cnt; i++ {
		v := s.slice.Index(i)
		ok := pv.Call([]reflect.Value{v})[0].Interface().(bool)
		if ok {
			cpy.slice = reflect.Append(cpy.slice, v)
		}
	}

	return cpy, nil
}

// Distinct remove duplicate values
func (s *serie) Distinct() Serie {
	cnt := s.Len()

	cpy := &serie{
		typ:       s.typ,
		converter: s.converter,
		slice:     reflect.MakeSlice(reflect.SliceOf(s.typ), 0, cnt),
	}

	m := make(map[interface{}]bool)

	for i := 0; i < cnt; i++ {
		v := s.slice.Index(i)

		if _, ok := m[v.Interface()]; !ok {
			cpy.slice = reflect.Append(cpy.slice, v)
			m[v.Interface()] = true
		}
	}

	return cpy
}

// Pick picks some indexes {at} to create a new serie
// If {at} is out of range, Pick will fill with a "zero" value
func (s *serie) Pick(at ...int) Serie {
	cpy := s.EmptyCopy().(*serie)
	cnt := s.Len()

	for _, pos := range at {
		if pos >= 0 && pos < cnt {
			cpy.slice = reflect.Append(cpy.slice, s.slice.Index(pos))
		} else {
			cpy.slice = reflect.Append(cpy.slice, s.converter.Call([]reflect.Value{reflect.Zero(s.typ)})...)
		}
	}
	return cpy
}
