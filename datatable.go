package datatable

// DataTable is our main struct
type DataTable struct {
	name  string
	cols  []*Column
	nrows int
}

// New creates a new datatable
func New(name string) *DataTable {
	return &DataTable{name: name}
}

// Name returns the datatable's name
func (t *DataTable) Name() string {
	return t.name
}

// Cols returns the columns in datatable
func (t *DataTable) Cols() []string {
	var cols []string
	for _, c := range t.cols {
		cols = append(cols, c.Name())
	}
	return cols
}

// Len returns the number of rows in datatable
func (t *DataTable) Len() int {
	return t.nrows
}

// AddColumn adds a new column
func (t *DataTable) AddColumn(name string, ctyp ColumnType, values ...interface{}) {
	col := NewColumn(name, ctyp)
	t.cols = append(t.cols, col)

	// Auto extends the table if needed
	if col.Set(values...) && !t.extend() {
		col.Size(t.nrows)
	}
}

// AddRow to add a new row
func (t *DataTable) AddRow(values ...interface{}) bool {
	cnt := len(t.cols)
	if len(values) != cnt {
		return false
	}

	for i, v := range values {
		col := t.cols[i]
		col.Append(v)
	}
	t.nrows++
	return true
}

// extends the rows of datatable if needed
// return true if the datatable has been extended.
func (t *DataTable) extend() bool {
	max := t.nrows
	for _, col := range t.cols {
		if size := col.Len(); size > max {
			max = size
		}
	}

	// extends if needed
	if max > t.nrows {
		t.nrows = max
		for _, col := range t.cols {
			col.Size(max)
		}
		return true
	}

	return false
}

// Find the column index by col name
func (t *DataTable) findColIndex(name string) int {
	for i, c := range t.cols {
		if c.Name() == name {
			return i
		}
	}

	return -1
}
