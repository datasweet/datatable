package datatable

// Where filters the datatable based on a predicate
func (t *DataTable) Where(predicate func(row Row) bool) *DataTable {
	if predicate == nil {
		return t.EmptyCopy()
	}

	if err := t.evaluateExpressions(); err != nil {
		panic(err)
	}

	subset := make([]int, 0, t.nrows) // max

	for i := 0; i < t.nrows; i++ {
		r := make(Row, len(t.cols))
		for _, col := range t.cols {
			r[col.name] = col.serie.Get(i)
		}
		if predicate(r) {
			subset = append(subset, i)
		}
	}

	cpy := t.EmptyCopy()

	if len(subset) == 0 {
		return cpy
	}

	cpy.nrows = len(subset)
	for i, col := range t.cols {
		cpy.cols[i].serie = col.serie.Pick(subset...)
	}

	return cpy
}
