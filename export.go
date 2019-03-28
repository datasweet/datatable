package datatable

// ToTable returns the table with raw datas
func ToTable(dt DataTable, headers bool) [][]interface{} {
	if dt == nil {
		return nil
	}

	dr := dt.Rows()

	var hr DataRow
	if headers {
		cols := dt.Columns()
		hr = make(DataRow, len(cols))
		for i, col := range cols {
			hr[i] = col.Label()
		}
	}

	var rows [][]interface{}
	rows = append(rows, hr)
	for _, r := range dr {
		rows = append(rows, r)
	}
	return rows
}

// ToMap returns the table as an array of map
func ToMap(dt DataTable) []map[string]interface{} {
	if dt == nil {
		return nil
	}

	dr := dt.Rows()
	nrows := len(dr)

	if nrows == 0 {
		return nil
	}

	rows := make([]map[string]interface{}, nrows)
	cols := dt.Columns()
	ncols := len(cols)

	for i := 0; i < nrows; i++ {
		rows[i] = make(map[string]interface{}, ncols)
		for j, col := range cols {
			rows[i][col.Name()] = col.GetAt(j)
		}
	}

	return rows
}
