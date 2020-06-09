package serie

import (
	"reflect"
	"sort"
)

func (s *serie) Swap(i, j int) {
	tmp := reflect.New(s.typ).Elem()
	a, b := s.slice.Index(i), s.slice.Index(j)
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)
}

func (s *serie) Less(i, j int) bool {
	return s.Compare(i, j) == Lt
}

// Compare values at indexes i, j
// panic if out of range
func (s *serie) Compare(i, j int) int {
	return s.comparer.Call([]reflect.Value{
		s.slice.Index(i),
		s.slice.Index(j),
	})[0].Interface().(int)
}

func (s *serie) SortAsc() {
	sort.Sort(s)
}

func (s *serie) SortDesc() {
	sort.Sort(sort.Reverse(s))
}
