package datatable

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

func (t *DataTable) Head(size int) *DataTable {
	return t.Subset(0, size)
}

func (t *DataTable) Tail(size int) *DataTable {
	return t.Subset(t.nrows-size, size)
}
