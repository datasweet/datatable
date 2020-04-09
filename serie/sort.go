package serie

import (
	"sort"

	"github.com/datasweet/datatable/value"
)

func (s *serie) Len() int {
	return len(s.values)
}

func (s *serie) Swap(i, j int) {
	if i < 0 || i >= len(s.values) || j < 0 || j >= len(s.values) || i == j {
		return
	}
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s *serie) Less(i, j int) bool {
	return i < 0 || i >= len(s.values) ||
		s.values[i].Compare(s.values[j]) == value.Lt
}

func (s *serie) SortAsc() Serie {
	cpy := s.Clone()
	sort.Sort(cpy)
	return cpy
}

func (s *serie) SortDesc() Serie {
	cpy := s.Clone()
	sort.Sort(sort.Reverse(cpy))
	return cpy
}
