package datatable

import (
	"fmt"
	"strings"
)

// New creates a new datatable
func New(name string) *DataTable {
	return &DataTable{name: name}
}

// DataTable is our main struct
type DataTable struct {
	name    string
	cols    []*column
	nrows   int
	dirty   bool
	hasExpr bool
}

// Name returns the datatable's name
func (t *DataTable) Name() string {
	return t.name
}

// Rename the datatable
func (t *DataTable) Rename(name string) {
	t.name = name
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

// ColumnIndex gets the index of the column with name
// returns -1 if not found
func (t *DataTable) ColumnIndex(name string) int {
	for i, col := range t.cols {
		if col.Name() == name {
			return i
		}
	}
	return -1
}

// Records returns the rows in datatable as string
// Computes all expressions.
func (t *DataTable) Records() [][]string {
	if t == nil {
		return nil
	}

	if err := t.evaluateExpressions(); err != nil {
		panic(err)
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
		for _, pos := range cols {
			r = append(r, fmt.Sprintf("%v", t.cols[pos].serie.Get(i)))
		}
		rows = append(rows, r)
	}
	return rows
}

// Rows returns the rows in datatable
// Computes all expressions.
func (t *DataTable) Rows() []Row {
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

	rows := make([]Row, 0, t.nrows)
	for i := 0; i < t.nrows; i++ {
		r := make(Row, len(cols))
		for name, pos := range cols {
			r[name] = t.cols[pos].serie.Get(i)
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

// Row gets the row at index
func (t *DataTable) Row(at int) Row {
	t.evaluateExpressions()
	r := make(Row, len(t.cols))
	for _, col := range t.cols {
		r[col.name] = col.serie.Get(at)
	}
	return r
}
