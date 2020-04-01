package serie

import (
	"reflect"

	"github.com/pkg/errors"
)

// serie implementation
type serie struct {
	typ    reflect.Type
	zv     interface{}
	values []interface{}
	caster Caster
	err    error
}

type Caster func(v interface{}) interface{}

func newSerie(zeroValue interface{}, caster Caster) *serie {
	typ := reflect.TypeOf(zeroValue)

	s := &serie{
		typ:    typ,
		zv:     zeroValue,
		caster: caster,
	}

	// Only accept value type // or slice
	switch kind := typ.Kind(); kind {
	case reflect.Invalid:
		s.err = errors.Errorf("%T is invalid type", typ)
	case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		s.err = errors.Errorf("serie called with a non-value type %T", typ)
	}

	return s
}

func (s *serie) Append(v ...interface{}) {
	for _, val := range v {
		s.values = append(s.values, s.caster(val))
	}
}

func (s *serie) Error() error {
	return s.err
}
