package datatable

import (
	"strings"

	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

func (t *DataTable) addColumn(col *column) error {
	if col == nil {
		return ErrNilColumn
	}

	// Check name
	if len(col.name) == 0 {
		return ErrNilColumnName
	}
	if c := t.Column(col.name); c != nil {
		err := errors.Errorf("column '%s' already exists", col.name)
		return errors.Wrap(err, ErrColumnAlreadyExists.Error())
	}

	// Check typ
	if len(col.typ) == 0 {
		return ErrNilColumnType
	}

	// Check formula
	if len(col.formulae) > 0 {
		parsed, err := expr.Parse(col.formulae)
		if err != nil {
			return errors.Wrapf(err, ErrFormulaeSyntax.Error())
		}
		col.expr = parsed
		t.hasExpr = true
	}

	// Check serie
	if col.serie == nil {
		return ErrNilSerie
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
	var options ColumnOptions
	for _, o := range opt {
		o(&options)
	}

	// create serie based on ctyp
	sr, err := newColumnSerie(ctyp, options)
	if err != nil {
		return errors.Wrap(err, ErrCreateSerie.Error())
	}

	return t.addColumn(&column{
		name:     strings.TrimSpace(name),
		typ:      ctyp,
		serie:    sr,
		hidden:   options.Hidden,
		formulae: strings.TrimSpace(options.Expr),
	})
}

// RenameColumn to rename a column
func (t *DataTable) RenameColumn(old, name string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		err := errors.New("you must provided a column name")
		return errors.Wrap(err, ErrNilColumnName.Error())
	}
	if c := t.Column(name); c != nil {
		err := errors.Errorf("column '%s' already exists", name)
		return errors.Wrap(err, ErrColumnAlreadyExists.Error())
	}
	if col := t.Column(old); col != nil {
		col.(*column).name = name
		return nil
	}
	err := errors.Errorf("column '%s' does not exist", name)
	return errors.Wrap(err, ErrColumnNotFound.Error())
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
		err := errors.Errorf("column '%s' not found", a)
		return errors.Wrap(err, ErrColumnNotFound.Error())
	}
	j := t.ColumnIndex(b)
	if j < 0 {
		err := errors.Errorf("column '%s' not found", b)
		return errors.Wrap(err, ErrColumnNotFound.Error())
	}
	t.cols[i], t.cols[j] = t.cols[j], t.cols[i]
	return nil
}
