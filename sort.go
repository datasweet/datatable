package datatable

import (
	"sort"

	"github.com/datasweet/datatable/serie"
)

// SortBy defines a sort to be applied
type SortBy struct {
	Column string
	Desc   bool
	index  int
}

// credits : https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
type sorter struct {
	t  *DataTable
	by []SortBy
}

func (s *sorter) Len() int {
	return s.t.nrows
}

func (s *sorter) Swap(i, j int) {
	s.t.SwapRow(i, j)
}

func (s *sorter) Less(i, j int) bool {
	for _, by := range s.by {
		sr := s.t.cols[by.index].serie
		switch cmp := sr.Compare(i, j); cmp {
		case serie.Eq:
			continue
		case serie.Gt:
			return by.Desc
		case serie.Lt:
			return !by.Desc
		}
	}

	return false
}

// Sort the table
func (t *DataTable) Sort(by ...SortBy) *DataTable {
	cpy := t.Copy()

	if len(by) == 0 {
		return cpy
	}

	for i := range by {
		b := &by[i]
		// Check if column exists
		b.index = t.ColumnIndex(b.Column)
		if b.index < 0 {
			return cpy
		}
	}

	srt := &sorter{
		t:  cpy,
		by: by,
	}

	sort.Sort(srt)
	return cpy
}
