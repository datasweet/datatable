package serie

import (
	"errors"

	"github.com/datasweet/datatable/value"
)

func (s *serie) Map(mapper func(v value.Value) value.Value) Serie {
	cpy := s.clone(false)
	if cpy.err != nil {
		return cpy
	}
	if mapper == nil {
		cpy.err = errors.New("argument k")
	}

	return cpy
}
