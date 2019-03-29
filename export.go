package datatable

// ToTable returns the table with raw datas
func ToTable(dt DataTable, headers bool) [][]interface{} {
	if dt == nil {
		return nil
	}

	dr := dt.Rows()
	cols := dt.Columns()
	ncols := len(cols)

	var hr []interface{}
	if headers {
		hr = make([]interface{}, ncols)
		for i, col := range cols {
			hr[i] = col.Label()
		}
	}

	var rows [][]interface{}
	rows = append(rows, hr)
	for _, r := range dr {
		row := make([]interface{}, ncols)
		for i, col := range cols {
			if v, ok := r[col.Name()]; ok {
				row[i] = v
			}
		}
		rows = append(rows, row)
	}
	return rows
}
