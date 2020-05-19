package datatable

// Concat datatables
func (left *DataTable) Concat(table ...*DataTable) (*DataTable, error) {
	out := left.EmptyCopy()
	out.dirty = true

	tables := make([]*DataTable, 0, 1+len(table))
	tables = append(tables, left)
	tables = append(tables, table...)

	for _, t := range tables {
		if t == nil {
			continue
		}
		for _, tc := range t.cols {
			pos := out.ColumnIndex(tc.name)
			if pos >= 0 {
				oc := out.cols[pos]
				if oc.IsComputed() {
					oc.serie.Grow(out.nrows - oc.serie.Len() + tc.serie.Len())
					continue
				}
				if err := oc.serie.Concat(tc.serie); err != nil {
					return nil, err
				}
			} else {
				out.cols = append(out.cols, tc.emptyCopy())
				oc := out.cols[len(out.cols)-1]
				oc.serie.Grow(out.nrows - oc.serie.Len())
				if oc.IsComputed() {
					oc.serie.Grow(tc.serie.Len())
					continue
				}
				if err := oc.serie.Concat(tc.serie); err != nil {
					return nil, err
				}
			}
		}
		out.nrows += t.nrows
	}

	// check
	for _, oc := range out.cols {
		size := out.nrows - oc.serie.Len()
		if size > 0 {
			oc.serie.Grow(size)
		}
	}

	return out, nil
}
