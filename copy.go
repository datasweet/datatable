package datatable

import (
	"github.com/datasweet/datatable/serie"
)

// Copy the datatable
// Mode = ShallowCopy: any change of value in original datatable will be reflected to the copy. But faster
// Mode = DeepCopy: copy any values
// Mode = EmptyCopy: just copy the columns with no values
func (t *DataTable) Copy(mode serie.CopyMode) *DataTable {
	cpy := &DataTable{
		name:    t.name,
		dirty:   t.dirty,
		hasExpr: t.hasExpr,
		nrows:   t.nrows,
		cols:    make([]*column, len(t.cols)),
	}

	if mode == serie.EmptyCopy {
		cpy.nrows = 0
	}

	for i, col := range t.cols {
		cpy.cols[i] = col.copy(mode)
	}

	return cpy
}
