package serie

import (
	"sort"

	"github.com/datasweet/datatable/value"
)

func (s *serie) Len() int {
	return len(s.values)
}

func (s *serie) Swap(i, j int) {
	if i == j {
		return
	}
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s *serie) Less(i, j int) bool {
	return s.values[i].Compare(s.values[j]) == value.Lt
}

func (s *serie) SortAsc() {
	sort.Sort(s)
}

func (s *serie) SortDesc() {
	sort.Sort(sort.Reverse(s))
}
