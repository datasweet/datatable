package datatable

// Subset selects rows at index with size
func (t *DataTable) Subset(at, size int) *DataTable {
	cpy := t.EmptyCopy()

	for i, col := range t.cols {
		cpy.cols[i].serie = col.serie.Subset(at, size)
	}

	if len(cpy.cols) > 0 {
		cpy.nrows = cpy.cols[0].serie.Len()

	}

	return cpy
}

// Head selects {size} first rows
func (t *DataTable) Head(size int) *DataTable {
	return t.Subset(0, size)
}

// Tail selects {size} last rows
func (t *DataTable) Tail(size int) *DataTable {
	return t.Subset(t.nrows-size, size)
}
