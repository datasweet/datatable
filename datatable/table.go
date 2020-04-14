package datatable

import (
	"strings"
)

// New creates a new datatable
func New(name string) *DataTable {
	return &DataTable{name: name}
}

// table is our main struct
type DataTable struct {
	name  string
	cols  []*column
	nrows int
	dirty bool
}

// Name returns the datatable's name
func (t *DataTable) Name() string {
	return t.name
}

// NumRows returns the number of rows in datatable
func (t *DataTable) NumRows() int {
	return t.nrows
}

// NumCols returns the number of visible columns in datatable
func (t *DataTable) NumCols() int {
	return len(t.Columns())
}

// Columns returns the visible column names in datatable
func (t *DataTable) Columns() []string {
	var cols []string
	for _, col := range t.cols {
		if col.IsVisible() {
			cols = append(cols, col.Name())
		}
	}
	return cols
}

// HiddenColumns returns the hidden column names in datatable
func (t *DataTable) HiddenColumns() []string {
	var cols []string
	for _, col := range t.cols {
		if !col.IsVisible() {
			cols = append(cols, col.Name())
		}
	}
	return cols
}

// Column gets the column with name
// returns nil if not found
func (t *DataTable) Column(name string) Column {
	for _, col := range t.cols {
		if col.Name() == name {
			return col
		}
	}
	return nil
}

// Records returns the rows in datatable as string
// Computes all expressions.
func (t *DataTable) Records() [][]string {
	if t.dirty {
		if err := t.evaluateExpressions(); err != nil {
			panic(err)
		}
	}

	// visible columns
	var cols []int
	for i, col := range t.cols {
		if col.IsVisible() {
			cols = append(cols, i)
		}
	}

	rows := make([][]string, 0, t.nrows)
	for i := 0; i < t.nrows; i++ {
		r := make([]string, 0, len(cols))
		for pos := range cols {
			r = append(r, t.cols[pos].serie.Value(i).String())
		}
		rows = append(rows, r)
	}
	return rows
}

// Rows returns the rows in datatable
// Computes all expressions.
func (t *DataTable) Rows() []Row {
	if t.dirty {
		if err := t.evaluateExpressions(); err != nil {
			panic(err)
		}
	}

	// visible columns
	cols := make(map[string]int)
	for i, col := range t.cols {
		if col.IsVisible() {
			cols[col.Name()] = i
		}
	}

	rows := make([]Row, 0, t.nrows)
	for i := 0; i < t.nrows; i++ {
		r := make(Row, len(cols))
		for name, pos := range cols {
			r[name] = t.cols[pos].serie.Value(i).Val()
		}
		rows = append(rows, r)
	}
	return rows
}

func (t *DataTable) String() string {
	var sb strings.Builder
	t.Print(&sb)
	return sb.String()
}