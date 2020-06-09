package serie

import (
	"reflect"

	"github.com/pkg/errors"
)

func (s *serie) asValue(i interface{}) []reflect.Value {
	in := i

	if cs, ok := in.(Serie); ok {
		in = cs.Slice()
	}

	rv := reflect.ValueOf(in)
	kind := rv.Kind()

	switch kind {
	case reflect.Slice, reflect.Array:
		arr := make([]reflect.Value, 0, rv.Len())
		for j := 0; j < rv.Len(); j++ {
			arr = append(arr, s.converter.Call([]reflect.Value{rv.Index(j)})...)
		}
		return arr
	case reflect.Invalid:
		// case "nil"
		return s.converter.Call([]reflect.Value{reflect.Zero(s.typ)})
	default:
		return s.converter.Call([]reflect.Value{rv})
	}

}

// Append values to the serie.
func (s *serie) Append(v ...interface{}) {
	values := make([]reflect.Value, 0, len(v))
	for _, val := range v {
		values = append(values, s.asValue(val)...)
	}
	s.slice = reflect.Append(s.slice, values...)
}

// Prepend values to the serie
func (s *serie) Prepend(v ...interface{}) error {
	return s.Insert(0, v...)
}

// Insert values to the serie at index
func (s *serie) Insert(at int, v ...interface{}) (err error) {
	n := s.Len()

	if at < 0 || ((at > 0 || n > 0) && at >= n) {
		return errors.Errorf("insert at [%d]: index out of range with length %d", at, n)
	}

	values := make([]reflect.Value, 0, len(v))
	for _, val := range v {
		values = append(values, s.asValue(val)...)
	}

	if len(values) == 0 {
		return nil
	}

	for i := 0; i < len(values); i++ {
		s.slice = reflect.Append(s.slice, reflect.Zero(s.typ))
	}

	// Refresh len
	n = s.Len()

	reflect.Copy(s.slice.Slice(at+len(values), n), s.slice.Slice(at, n))

	for i, rv := range values {
		s.slice.Index(i + at).Set(rv)
	}

	return nil
}

// Set a value at index
func (s *serie) Set(at int, v interface{}) error {
	if at < 0 || at >= s.Len() {
		return errors.Errorf("set at [%d]: index out of range with length %d", at, s.Len())
	}
	values := s.asValue(v)

	if len(values) != 1 {
		return errors.Errorf("set at [%d]: can't flatten slice with set", at)
	}

	s.slice.Index(at).Set(values[0])
	return nil
}

// Delete a value at index
func (s *serie) Delete(at int) error {
	cnt := s.Len()
	if at < 0 || at >= cnt {
		return errors.Errorf("delete at [%d]: index out of range with length %d", at, cnt)
	}
	if at < cnt-1 {
		reflect.Copy(s.slice.Slice(at, cnt), s.slice.Slice(at+1, cnt))
	}
	s.slice = s.slice.Slice(0, cnt-1)
	return nil
}

// Grow the serie with size
// Grow will create zero value
func (s *serie) Grow(size int) error {
	if size < 0 {
		return errors.Errorf("grow: size '%d' must be > 0", size)
	}
	for i := 0; i < size; i++ {
		s.slice = reflect.Append(s.slice, reflect.Zero(s.typ))
	}
	return nil
}

// Shrink the serie with size
func (s *serie) Shrink(size int) error {
	if size < 0 {
		return errors.Errorf("shrink: size '%d' must be > 0", size)
	}
	cnt := s.Len()
	if size > cnt {
		return errors.Errorf("shrink: size '%d' must be < length '%d'", size, cnt)
	}
	s.slice = s.slice.Slice(0, cnt-size)
	return nil
}

// Concat the serie (mutate) with others series
// series provided must be the same type as the source serie
func (s *serie) Concat(serie ...Serie) error {
	if len(serie) == 0 {
		return nil
	}

	for i, other := range serie {
		if other.Type() != s.Type() {
			return errors.Errorf("concat: serie #%d is not the same type as source", i)
		}

		s.Append(other.Slice())
	}

	return nil
}

func (s *serie) Clear() {
	s.slice = reflect.MakeSlice(reflect.SliceOf(s.typ), 0, 0)
}
