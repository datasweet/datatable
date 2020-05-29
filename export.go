package datatable

func (t *DataTable) ToMap() []map[string]interface{} {
	if t == nil {
		return nil
	}

	if err := t.evaluateExpressions(); err != nil {
		panic(err)
	}

	// visible columns
	cols := make(map[string]int)
	for i, col := range t.cols {
		if col.IsVisible() {
			cols[col.Name()] = i
		}
	}

	rows := make([]map[string]interface{}, 0, t.nrows)
	for i := 0; i < t.nrows; i++ {
		r := make(map[string]interface{}, len(cols))
		for name, pos := range cols {
			r[name] = t.cols[pos].serie.Get(i)
		}
		rows = append(rows, r)
	}
	return rows
}

func (t *DataTable) ToTable() [][]interface{} {
	if t == nil {
		return nil
	}

	if err := t.evaluateExpressions(); err != nil {
		panic(err)
	}

	rows := make([][]interface{}, 0, t.nrows+1)

	// visible columns
	var headers []interface{}
	var cols []int
	for i, col := range t.cols {
		if col.IsVisible() {
			cols = append(cols, i)
			headers = append(headers, col.Name())
		}
	}

	rows = append(rows, headers)

	for i := 0; i < t.nrows; i++ {
		r := make([]interface{}, 0, len(cols))
		for _, pos := range cols {
			r = append(r, t.cols[pos].serie.Get(i))
		}
		rows = append(rows, r)
	}
	return rows
}

type Schema struct {
	Name    string          `json:"name"`
	Columns []SchemaColumn  `json:"cols"`
	Rows    [][]interface{} `json:"rows"`
}

type SchemaColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (t *DataTable) ToSchema() *Schema {
	if t == nil {
		return nil
	}

	if err := t.evaluateExpressions(); err != nil {
		panic(err)
	}

	schema := &Schema{
		Name: t.name,
		Rows: make([][]interface{}, 0, t.nrows),
	}

	// visible columns
	var cols []int
	for i, col := range t.cols {
		if col.IsVisible() {
			cols = append(cols, i)
			schema.Columns = append(schema.Columns, SchemaColumn{Type: col.Type().Name(), Name: col.Name()})
		}
	}

	for i := 0; i < t.nrows; i++ {
		r := make([]interface{}, 0, len(cols))
		for _, pos := range cols {
			r = append(r, t.cols[pos].serie.Get(i))
		}
		schema.Rows = append(schema.Rows, r)
	}

	return schema
}
