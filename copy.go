package datatable

// EmptyCopy copies the structure of datatable (no values)
func (t *DataTable) EmptyCopy() *DataTable {
	cpy := &DataTable{
		name:    t.name,
		dirty:   t.dirty,
		hasExpr: t.hasExpr,
		nrows:   0,
		cols:    make([]*column, len(t.cols)),
	}

	for i, col := range t.cols {
		cpy.cols[i] = col.emptyCopy()
	}

	return cpy
}

// Copy the datatable
func (t *DataTable) Copy() *DataTable {
	cpy := &DataTable{
		name:    t.name,
		dirty:   t.dirty,
		hasExpr: t.hasExpr,
		nrows:   t.nrows,
		cols:    make([]*column, len(t.cols)),
	}

	for i, col := range t.cols {
		cpy.cols[i] = col.copy()
	}

	return cpy
}
