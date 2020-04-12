package datatable

// Row gets the row at index
func (t *DataTable) Row(at int) Row {
	r := make(Row, len(t.cols))
	for _, col := range t.cols {
		r[col.name] = col.serie.Value(at).Val()
	}
	return r
}
