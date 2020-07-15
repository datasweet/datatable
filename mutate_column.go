package datatable

import (
	"strings"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

func (t *DataTable) addColumn(name string, sr serie.Serie, formulae string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("you must provided a column name")
	}
	if c := t.Column(name); c != nil {
		return errors.Errorf("column '%s' already exists", name)
	}
	if sr == nil {
		return errors.New("nil serie provided")
	}

	var ex expr.Node
	if formulae = strings.TrimSpace(formulae); len(formulae) > 0 {
		parsed, err := expr.Parse(formulae)
		if err != nil {
			return errors.Wrapf(err, "formulae syntax")
		}
		ex = parsed
		t.hasExpr = true
	}

	sr = sr.Copy()
	l := sr.Len()

	if l < t.nrows {
		sr.Grow(t.nrows - l)
	} else if l > t.nrows {
		size := l - t.nrows
		for _, col := range t.cols {
			col.serie.Grow(size)
		}
		t.nrows = l
	}

	t.cols = append(t.cols, &column{
		name:     name,
		serie:    sr,
		formulae: formulae,
		expr:     ex,
	})
	t.dirty = true
	return nil

}

// AddColumn to datatable with a serie of T
func (t *DataTable) AddColumn(name string, sr serie.Serie) error {
	return t.addColumn(name, sr, "")
}

// AddExprColumn to add a calculated column with a serie of T
func (t *DataTable) AddExprColumn(name string, sr serie.Serie, formulae string) error {
	return t.addColumn(name, sr.EmptyCopy(), formulae)
}

// AddIntColumn to add a column of nullable int
func (t *DataTable) AddIntColumn(name string, v ...interface{}) error {
	return t.addColumn(name, serie.IntN(v...), "")
}

// AddIntExprColumn to add a calculated column of nullable int
func (t *DataTable) AddIntExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.IntN(), expr)
}

// AddInt32Column to add a column of nullable int32
func (t *DataTable) AddInt32Column(name string, v ...interface{}) error {
	return t.addColumn(name, serie.Int32N(v...), "")
}

// AddInt32ExprColumn to add a calculated column of nullable int32
func (t *DataTable) AddInt32ExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.Int32N(), expr)
}

// AddInt64Column to add a column of nullable int64
func (t *DataTable) AddInt64Column(name string, v ...interface{}) error {
	return t.addColumn(name, serie.Int64N(v...), "")
}

// AddInt64ExprColumn to add a calculated column of nullable int32
func (t *DataTable) AddInt64ExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.Int64N(), expr)
}

// AddBoolColumn to add a column of nullable bool
func (t *DataTable) AddBoolColumn(name string, v ...interface{}) error {
	return t.addColumn(name, serie.BoolN(v...), "")
}

// AddBoolExprColumn to add a calculated column of nullable bool
func (t *DataTable) AddBoolExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.BoolN(), expr)
}

// AddFloat32Column to add a column of nullable float32
func (t *DataTable) AddFloat32Column(name string, v ...interface{}) error {
	return t.addColumn(name, serie.Float32N(v...), "")
}

// AddFloat32ExprColumn to add a calculated column of nullable bool
func (t *DataTable) AddFloat32ExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.Float32N(), expr)
}

// AddFloat64Column to add a column of nullable float64
func (t *DataTable) AddFloat64Column(name string, v ...interface{}) error {
	return t.addColumn(name, serie.Float64N(v...), "")
}

// AddFloat64ExprColumn to add a calculated column of nullable bool
func (t *DataTable) AddFloat64ExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.Float64N(), expr)
}

// AddStringColumn to add a column of nullable string
func (t *DataTable) AddStringColumn(name string, v ...interface{}) error {
	return t.addColumn(name, serie.StringN(v...), "")
}

// AddStringExprColumn to add a calculated column of nullable bool
func (t *DataTable) AddStringExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.StringN(), expr)
}

// AddTimeColumn to add a column of nullable time
func (t *DataTable) AddTimeColumn(name string, v ...interface{}) error {
	return t.addColumn(name, serie.TimeN(v...), "")
}

// AddTimeExprColumn to add a calculated column of nullable bool
func (t *DataTable) AddTimeExprColumn(name string, expr string) error {
	return t.addColumn(name, serie.TimeN(), expr)
}

// RenameColumn to rename a column
func (t *DataTable) RenameColumn(old, name string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("you must provided a column name")
	}
	if c := t.Column(name); c != nil {
		return errors.Errorf("column '%s' already exists", name)
	}
	if col := t.Column(old); col != nil {
		col.(*column).name = name
		return nil
	}
	return errors.Errorf("column '%s' does not exist", name)
}

// HideAll to hides all column
// a hidden column will not be exported
func (t *DataTable) HideAll() {
	for _, col := range t.cols {
		col.hidden = true
	}
}

// HideColumn hides a column
// a hidden column will not be exported
func (t *DataTable) HideColumn(name string) {
	if c := t.Column(name); c != nil {
		(c.(*column)).hidden = true
	}
}

// ShowAll to show all column
// a shown column will be exported
func (t *DataTable) ShowAll() {
	for _, col := range t.cols {
		col.hidden = false
	}
}

// ShowColumn shows a column
// a shown column will be exported
func (t *DataTable) ShowColumn(name string) {
	if c := t.Column(name); c != nil {
		(c.(*column)).hidden = false
	}
}

// SwapColumn to swap 2 columns
func (t *DataTable) SwapColumn(a, b string) error {
	i := t.ColumnIndex(a)
	if i < 0 {
		return errors.Errorf("column '%s' not found", a)
	}
	j := t.ColumnIndex(b)
	if j < 0 {
		return errors.Errorf("column '%s' not found", b)
	}
	t.cols[i], t.cols[j] = t.cols[j], t.cols[i]
	return nil
}
