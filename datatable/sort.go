package datatable

import (
	"sort"

	"github.com/datasweet/datatable/value"
)

// By defines a sort to be applied
type By struct {
	Column string
	Desc   bool
	index  int
}

// credits : https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
type sorter struct {
	t  *DataTable
	by []By
}

func (s *sorter) Len() int {
	return s.t.nrows
}

func (s *sorter) Swap(i, j int) {
	s.t.SwapRow(i, j)
}

func (s *sorter) Less(i, j int) bool {
	for _, by := range s.by {
		serie := s.t.cols[by.index].serie
		a, b := serie.Value(i), serie.Value(j)

		switch cmp := a.Compare(b); cmp {
		case value.Eq:
			continue
		case value.Gt:
			return by.Desc
		case value.Lt:
			return !by.Desc
		}
	}

	return false
}

// Sort the table
func (t *DataTable) Sort(by ...By) {
	if len(by) == 0 {
		return
	}

	for i := range by {
		b := &by[i]
		// Check if column exists
		b.index = t.ColumnIndex(b.Column)
		if b.index < 0 {
			return
		}
	}

	srt := &sorter{
		t:  t,
		by: by,
	}

	sort.Sort(srt)
}
