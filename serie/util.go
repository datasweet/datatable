package serie

import (
	"github.com/datasweet/datatable/value"
)

func (s *serie) Clone() Serie {
	return s.clone(true)
}

func (s *serie) clone(includeValues bool) *serie {
	var values []value.Value

	if includeValues {
		values = make([]value.Value, len(s.values))
		for i, val := range s.values {
			values[i] = val.Clone()
		}
	}

	return &serie{
		typ:     s.typ,
		builder: s.builder,
		values:  values,
		err:     s.err,
	}
}
