package datatable

// ToTable returns the table with raw datas
func ToTable(dt DataTable, headers bool) [][]interface{} {
	if dt == nil {
		return nil
	}

	dr := dt.Rows()
	cols := dt.Columns()

	var hr []interface{}
	if headers {
		for _, col := range cols {
			if !col.Hidden() {
				label := col.Label()
				if len(label) == 0 {
					label = col.Name()
				}
				hr = append(hr, label)
			}
		}
	}

	var rows [][]interface{}
	rows = append(rows, hr)
	for _, r := range dr {
		var row []interface{}
		for _, col := range cols {
			if !col.Hidden() {
				row = append(row, r[col.Name()])
			}
		}
		rows = append(rows, row)
	}
	return rows
}
