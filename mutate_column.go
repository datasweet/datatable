package datatable

import (
	"strings"

	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

func (t *DataTable) addColumn(col *column) error {
	if col == nil {
		return errors.New("nil column")
	}

	// Check name
	if len(col.name) == 0 {
		return errors.New("nil column name")
	}
	if c := t.Column(col.name); c != nil {
		return errors.Errorf("column '%s' already exists", col.name)
	}

	// Check typ
	if len(col.typ) == 0 {
		return errors.New("nil column type")
	}

	// Check formula
	if len(col.formulae) > 0 {
		parsed, err := expr.Parse(col.formulae)
		if err != nil {
			return errors.Wrapf(err, "formulae syntax")
		}
		col.expr = parsed
		t.hasExpr = true
	}

	// Check serie
	if col.serie == nil {
		return errors.New("nil serie")
	}
	ln := col.serie.Len()

	if ln < t.nrows {
		col.serie.Grow(t.nrows - ln)
	} else if ln > t.nrows {
		size := ln - t.nrows
		for _, col := range t.cols {
			col.serie.Grow(size)
		}
		t.nrows = ln
	}

	t.cols = append(t.cols, col)
	t.dirty = true
	return nil
}

// AddColumn to datatable with a serie of T
func (t *DataTable) AddColumn(name string, ctyp ColumnType, opt ...ColumnOption) error {
	options := ColumnOptions{
		NullAvailable: true,
	}
	for _, o := range opt {
		o(&options)
	}

	// create serie based on ctyp
	sr, err := newColumnSerie(ctyp, options)
	if err != nil {
		return errors.Wrap(err, "create serie")
	}

	return t.addColumn(&column{
		name:     strings.TrimSpace(name),
		typ:      ctyp,
		serie:    sr,
		formulae: strings.TrimSpace(options.Expr),
	})
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
