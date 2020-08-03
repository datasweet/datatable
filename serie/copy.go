package serie

import (
	"reflect"
)

func (s *serie) makeEmptyCopy(capacity int) *serie {
	return &serie{
		typ:        s.typ,
		converter:  s.converter,
		comparer:   s.comparer,
		interfacer: s.interfacer,
		slice:      reflect.MakeSlice(reflect.SliceOf(s.typ), 0, capacity),
	}
}

func (s *serie) EmptyCopy() Serie {
	return s.makeEmptyCopy(0)
}

func (s *serie) Copy() Serie {
	cnt := s.Len()
	cpy := &serie{
		typ:        s.typ,
		converter:  s.converter,
		comparer:   s.comparer,
		interfacer: s.interfacer,
		slice:      reflect.MakeSlice(reflect.SliceOf(s.typ), cnt, cnt),
	}
	reflect.Copy(cpy.slice, s.slice)
	return cpy
}
