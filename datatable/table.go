package datatable

import (
	"fmt"
	"io"
	"strings"

	"github.com/datasweet/datatable/serie"
)

type DataTable interface {
	Name() string
	NumCols() int
	NumRows() int
	Columns() []string
	HiddenColumns() []string
	Rows() []Row
	Column(name string) Column
	Row(at int) Row

	// Mutate
	AddColumn(name string, serie serie.Serie) error
	AddExprColumn(name string, typ ExprType, formulae string) error
	NewRow() Row
	Append(r ...Row)
	AppendRow(v ...interface{}) error
	//Swap(colA, colB string) bool

	// DeleteRow(at) bool
	// UpdateRow(at, values...) bool
	// DeleteColumn(name string)
	// SortBy(colName ...string)

	// Print
	Print(writer io.Writer, opt ...PrintOption)
	fmt.Stringer
}

// New creates a new datatable
func New(name string) DataTable {
	return &table{name: name}
}

// table is our main struct
type table struct {
	name  string
	cols  []*column
	nrows int
	dirty bool
}

// Name returns the datatable's name
func (t *table) Name() string {
	return t.name
}

// NumRows returns the number of rows in datatable
func (t *table) NumRows() int {
	return t.nrows
}

// NumCols returns the number of visible columns in datatable
func (t *table) NumCols() int {
	return len(t.Columns())
}

// Columns returns the visible column names in datatable
func (t *table) Columns() []string {
	var cols []string
	for _, col := range t.cols {
		if col.IsVisible() {
			cols = append(cols, col.Name())
		}
	}
	return cols
}

// HiddenColumns returns the hidden column names in datatable
func (t *table) HiddenColumns() []string {
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
func (t *table) Column(name string) Column {
	for _, col := range t.cols {
		if col.Name() == name {
			return col
		}
	}
	return nil
}

// Records returns the rows in datatable as string
// Computes all expressions.
func (t *table) Records() [][]string {
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
func (t *table) Rows() []Row {
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

func (t *table) String() string {
	var sb strings.Builder
	t.Print(&sb)
	return sb.String()
}
