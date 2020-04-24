package serie

import (
	"github.com/datasweet/datatable/value"
)

type CopyMode uint8

const (
	ShallowCopy = iota
	DeepCopy
	EmptyCopy
)

// Copy the serie
// Mode = ShallowCopy: any change of value in original serie will be reflected to the copy. But faster
// Mode = DeepCopy: copy any values
// Mode = EmptyCopy: just copy the container
func (s *serie) Copy(mode CopyMode) Serie {
	cpy := &serie{typ: s.typ, builder: s.builder}

	switch mode {
	case EmptyCopy:
	case DeepCopy:
		cpy.values = make([]value.Value, len(s.values))
		for i, val := range s.values {
			cpy.values[i] = val.Clone()
		}
	default:
		cpy.values = make([]value.Value, len(s.values))
		copy(cpy.values, s.values)
	}
	return cpy
}
