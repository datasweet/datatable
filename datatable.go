package datatable

import (
	"github.com/pkg/errors"
)

type DataTable interface {
	Name() string
	NumCols() int
	NumRows() int
	Columns() []DataColumn
	AddColumn(name string, ctyp ColumnType, values ...interface{}) (DataColumn, error)
	AddExprColumn(name string, formulae string) (DataColumn, error)
	GetColumn(name string) (int, DataColumn)
	Rows() []DataRow
	GetRow(at int) DataRow
	NewRow() DataRow
	AddRow(dr DataRow) bool
	AppendRow(v ...interface{}) bool
	Swap(colA, colB string) bool

	// DeleteRow(at) bool
	// UpdateRow(at, values...) bool
	// DeleteColumn(name string)
	// SortBy(colName ...string)

}

// table is our main struct
type table struct {
	name   string
	cols   []*column
	cindex map[string]int
	nrows  int
}

// New creates a new datatable
func New(name string) DataTable {
	return &table{
		name:   name,
		cindex: make(map[string]int, 0),
	}
}

// Name returns the datatable's name
func (t *table) Name() string {
	return t.name
}

// NumRows returns the number of rows in datatable
func (t *table) NumRows() int {
	return t.nrows
}

// NumCols returns the number of cols in datatable
func (t *table) NumCols() int {
	return len(t.cols)
}

// Columns returns the columns in datatable
func (t *table) Columns() []DataColumn {
	cols := make([]DataColumn, len(t.cols))
	for i, c := range t.cols {
		cols[i] = c
	}
	return cols
}

// AddColumn adds a new column
func (t *table) AddColumn(name string, ctyp ColumnType, values ...interface{}) (DataColumn, error) {
	if _, c := t.GetColumn(name); c != nil {
		return nil, errors.Errorf("column '%s' already exists", name)
	}

	col := newColumn(name, ctyp)
	t.cols = append(t.cols, col)
	t.cindex[name] = len(t.cols) - 1

	col.Set(values...)

	// Auto extends the table if needed
	l := col.Len()
	if l < t.nrows {
		col.Size(t.nrows)
	} else if l > t.nrows {
		t.size(l)
	}

	return col, nil
}

// AddExprColumn to adds a column with a binded expression
func (t *table) AddExprColumn(name string, formulae string) (DataColumn, error) {
	if _, c := t.GetColumn(name); c != nil {
		return nil, errors.Errorf("column '%s' already exists", name)
	}

	col, err := newExprColumn(name, formulae)
	if err != nil {
		return nil, errors.Wrap(err, "can't create expr column")
	}
	t.cols = append(t.cols, col)
	t.cindex[name] = len(t.cols) - 1

	return col, nil
}

// GetColumn returns the column index and the column itself
// If not exists returns -1, nil
func (t *table) GetColumn(name string) (int, DataColumn) {
	if i, ok := t.cindex[name]; ok {
		if i < len(t.cols) {
			return i, t.cols[i]
		}
	}
	return -1, nil
}

// Rows returns the rows in datatable
// Computes all expressions.
func (t *table) Rows() []DataRow {
	t.evaluateExpressions()

	rows := make([]DataRow, t.nrows)

	for i := 0; i < t.nrows; i++ {
		rows[i] = make(DataRow, len(t.cols))
		for _, col := range t.cols {
			rows[i][col.Name()] = col.GetAt(i)
		}
	}

	return rows
}

func (t *table) NewRow() DataRow {
	dr := make(map[string]interface{}, len(t.cols))

	for _, c := range t.cols {
		dr[c.Name()] = c.ZeroValue()
	}

	return dr
}

// AddRow to add
func (t *table) AddRow(dr DataRow) bool {
	if dr == nil {
		return false
	}
	for k, v := range dr {
		if i, _ := t.GetColumn(k); i >= 0 {
			if col := t.cols[i]; !col.IsExpr() {
				col.Append(v)
			}
		}
	}

	t.nrows++
	return true
}

// AppendRow add a new row to our table
// Must faster than dt.AddRow(dt.NewRow()) when you know the structure of datatable
// <!> a expr col will ignore the passed value
func (t *table) AppendRow(v ...interface{}) bool {
	lv := len(v)
	if lv == 0 {
		return false
	}

	for i, col := range t.cols {
		if !col.IsExpr() && i < lv {
			col.Append(v[i])
		} else {
			col.Append(col.ZeroValue())
		}
	}

	t.nrows++

	return true
}

// GetRow returns the datarow at index
func (t *table) GetRow(at int) DataRow {
	if at < 0 || at >= t.NumRows() {
		return nil
	}
	row := make(DataRow, len(t.cols))
	for _, c := range t.cols {
		row[c.Name()] = c.GetAt(at)
	}
	return row
}

// Size sets the numbers of rows in our datatabe
// Extend or shrink the rows
func (t *table) size(size int) bool {
	if size < 0 {
		return false
	}
	ok := true
	for _, c := range t.cols {
		ok = ok && c.Size(size)
	}
	t.nrows = size
	return ok
}

// Swap swap 2 columns
func (t *table) Swap(colA, colB string) bool {
	a, _ := t.GetColumn(colA)
	if a < 0 {
		return false
	}

	b, _ := t.GetColumn(colB)
	if b < 0 {
		return false
	}

	tmp := t.cols[a]
	t.cols[a] = t.cols[b]
	t.cols[b] = tmp
	return true
}

// hasColWithExpr to check if the datatable has at least one
// column with an expression to be evaluate
func (t *table) exprColsIndex() []int {
	var indexes []int
	for i, c := range t.cols {
		if c.expr != nil {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// evaluateExpressions to evaluate all columns with a binded expression
func (t *table) evaluateExpressions() error {
	exprCols := t.exprColsIndex()
	lec := len(exprCols)
	if lec == 0 {
		return nil
	}

	// Initialize params
	params := make(map[string]interface{}, len(t.cols))
	j := 0
	nextExpr := exprCols[j]

	for i, c := range t.cols {
		if i == nextExpr {
			j++
			if j >= lec {
				nextExpr = -1
			} else {
				nextExpr = exprCols[j]
			}
			continue
		}
		params[c.name] = c.rows
	}

	// Evaluate
	for _, ic := range exprCols {
		col := t.cols[ic]
		res, err := col.expr.Eval(params)
		if err != nil {
			return err
		}

		// clear
		col.Size(0)

		if arr, ok := res.([]interface{}); ok {
			col.Set(arr...)

			// Sync size
			l := len(arr)
			if l < t.nrows {
				col.Size(t.nrows)
			} else if l > t.nrows {
				t.size(l)
			}

		} else {
			// duplicate res on all row
			ar := make([]interface{}, t.nrows)
			for i := 0; i < t.nrows; i++ {
				ar[i] = res
			}
			col.Set(ar...)
		}

		// now add this column as new param
		params[col.name] = col.rows
	}

	return nil
}
