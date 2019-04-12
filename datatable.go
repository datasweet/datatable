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
	AddRow(dr ...DataRow) bool
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
	rows   []DataRow
	cindex map[string]int
	dirty  bool
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
	return len(t.rows)
}

// NumCols returns the number of cols in datatable
func (t *table) NumCols() int {
	return len(t.cols)
}

// Columns returns the columns in datatable
func (t *table) Columns() []DataColumn {
	columns := make([]DataColumn, len(t.cols))
	for i, c := range t.cols {
		columns[i] = c
	}
	return columns
}

// AddColumn adds a new column
func (t *table) AddColumn(name string, ctyp ColumnType, values ...interface{}) (DataColumn, error) {
	if _, c := t.GetColumn(name); c != nil {
		return nil, errors.Errorf("column '%s' already exists", name)
	}

	col := newColumn(name, ctyp)
	t.cols = append(t.cols, col)
	t.cindex[name] = len(t.cols) - 1

	l := len(values)
	nrows := len(t.rows)

	if l < nrows {
		for i := 0; i < l; i++ {
			t.rows[i][name] = col.AsValue(values[i])
		}
		for i := l; i < nrows; i++ {
			t.rows[i][name] = col.ZeroValue()
		}
	} else {
		for i := 0; i < nrows; i++ {
			t.rows[i][name] = col.AsValue(values[i])
		}
		for i := nrows; i < l; i++ {
			dr := t.NewRow()
			dr[name] = col.AsValue(values[i])
			t.rows = append(t.rows, dr)
		}
	}

	t.dirty = true

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

	t.dirty = true

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

// DeleteColumnn to delete a column
func (t *table) DeleteColumn(name string) error {
	index, _ := t.GetColumn(name)
	if index < 0 {
		return errors.Errorf("column '%s' does not exists", name)
	}

	// delete columns in row
	for _, dr := range t.rows {
		delete(dr, name)
	}

	// delete column
	t.cols = append(t.cols[:index], t.cols[index+1:]...)

	t.dirty = true

	return nil
}

// Rows returns the rows in datatable
// Computes all expressions.
func (t *table) Rows() []DataRow {
	if t.dirty {
		t.evaluateExpressions()
	}
	return t.rows
}

// NewRow to create a new row based on column schema
func (t *table) NewRow() DataRow {
	dr := make(map[string]interface{}, len(t.cols))
	for _, c := range t.cols {
		dr[c.Name()] = c.ZeroValue()
	}
	return dr
}

// AddRow to add
func (t *table) AddRow(dr ...DataRow) bool {
	added := 0
	for _, r := range dr {
		if r != nil {
			t.rows = append(t.rows, r)
			added++
		}
	}
	t.dirty = (added > 0)
	return added == len(dr)
}

// AppendRow add a new row to our table
// Must faster than dt.AddRow(dt.NewRow()) when you know the structure of datatable
// <!> a expr col will ignore the passed value
func (t *table) AppendRow(v ...interface{}) bool {
	lv := len(v)
	if lv == 0 {
		return false
	}

	dr := make(DataRow, len(t.cols))

	for i, col := range t.cols {
		if !col.Computed() && i < lv {
			dr[col.Name()] = col.AsValue(v[i])
		} else {
			dr[col.Name()] = col.ZeroValue()
		}
	}

	t.rows = append(t.rows, dr)
	t.dirty = true

	return true
}

// GetRow returns the datarow at index
func (t *table) GetRow(at int) DataRow {
	if at < 0 || at >= t.NumRows() {
		return nil
	}

	if t.dirty {
		t.evaluateExpressions()
	}

	return t.rows[at]
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

// evaluateExpressions to evaluate all columns with a binded expression
func (t *table) evaluateExpressions() error {
	var cols []string
	var exprCols []int
	for i, c := range t.cols {
		if c.Computed() {
			exprCols = append(exprCols, i)
		} else {
			cols = append(cols, c.Name())
		}
	}

	l := len(exprCols)
	if l == 0 {
		t.dirty = false
		return nil
	}

	// Initialize params
	params := make(map[string][]interface{}, len(t.cols))
	for _, r := range t.rows {
		for _, name := range cols {
			params[name] = append(params[name], r.Get(name))
		}
	}

	// Evaluate
	for _, idx := range exprCols {
		col := t.cols[idx]
		res, err := col.expr.Eval(params)
		if err != nil {
			return err
		}

		name := col.Name()

		if arr, ok := res.([]interface{}); ok {
			// Is array
			l := len(arr)
			nrows := len(t.rows)

			// Resyn size and add the res as new param
			if l < nrows {
				for i := 0; i < l; i++ {
					t.rows[i][name] = arr[i]
					params[name] = append(params[name], arr[i])
				}
				for i := l; i < nrows; i++ {
					t.rows[i][name] = nil
					params[name] = nil
				}
			} else {
				for i := 0; i < nrows; i++ {
					t.rows[i][name] = arr[i]
					params[name] = append(params[name], arr[i])
				}
				for i := nrows; i < l; i++ {
					dr := make(DataRow, len(t.cols))
					for _, cname := range cols {
						val := col.ZeroValue()
						if cname == name {
							val = arr[i]
						}
						dr[cname] = val
						params[cname] = append(params[cname], val)
					}
					t.rows = append(t.rows, dr)
				}
			}

		} else {
			// Is scalar
			for _, r := range t.rows {
				r[name] = res
				params[name] = append(params[name], res)
			}
		}
	}

	t.dirty = false

	return nil
}
