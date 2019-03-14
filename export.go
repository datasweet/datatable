package datatable

// Table returns the table with raw datas
func (t *DataTable) Table(headers bool) [][]interface{} {
	var h int
	if headers {
		h = 1
	}
	rows := make([][]interface{}, t.nrows+h)

	if headers {
		rows[0] = make([]interface{}, len(t.cols))
		for i, col := range t.cols {
			rows[0][i] = col.Name()
		}
	}

	for j := 0; j < t.nrows; j++ {
		n := j + h
		rows[n] = make([]interface{}, len(t.cols))
		for i, col := range t.cols {
			rows[n][i] = col.GetAt(j)
		}
	}

	return rows
}

// ToMap returns the table as an array of map
func (t *DataTable) ToMap() []map[string]interface{} {
	rows := make([]map[string]interface{}, t.nrows)

	for j := 0; j < t.nrows; j++ {
		rows[j] = make(map[string]interface{}, len(t.cols))
		for i, col := range t.cols {
			rows[j][col.Name()] = col.GetAt(i)
		}
	}

	return rows
}
