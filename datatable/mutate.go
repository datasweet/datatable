package datatable

import (
	"strings"

	"github.com/datasweet/expr"

	"github.com/datasweet/datatable/serie"
	"github.com/pkg/errors"
)

type ExprType uint8

const (
	ExprRaw ExprType = iota
	ExprBool
	ExprNumber
	ExprString
)

func (t *table) addColumn(name string, serie serie.Serie, formulae string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("you must provided a column name")
	}
	if c := t.Column(name); c != nil {
		return errors.Errorf("column '%s' already exists", name)
	}
	if serie == nil {
		return errors.New("nil serie provided")
	}
	if err := serie.Error(); err != nil {
		return errors.Wrap(err, "serie provided has error")
	}

	var ex expr.Node
	if formulae = strings.TrimSpace(formulae); len(formulae) > 0 {
		parsed, err := expr.Parse(formulae)
		if err != nil {
			return errors.Wrapf(err, "formulae syntax")
		}
		ex = parsed
	}

	serie = serie.Clone()
	l := serie.Len()

	if l < t.nrows {
		serie.Grow(t.nrows - l)
	} else if l > t.nrows {
		size := l - t.nrows
		for _, col := range t.cols {
			col.serie.Grow(size)
		}
		t.nrows = l
	}

	t.cols = append(t.cols, &column{
		name:     name,
		serie:    serie,
		formulae: formulae,
		expr:     ex,
	})
	t.dirty = true
	return nil

}

func (t *table) AddColumn(name string, serie serie.Serie) error {
	return t.addColumn(name, serie, "")
}

func (t *table) AddExprColumn(name string, typ ExprType, formulae string) error {
	var s serie.Serie
	switch typ {
	case ExprBool:
		s = serie.Bool()
	case ExprNumber:
		s = serie.Float64()
	case ExprString:
		s = serie.String()
	case ExprRaw:
		s = serie.Raw()
	default:
		return errors.Errorf("unknown expression type %v", typ)
	}

	return t.addColumn(name, s, formulae)
}

func (t *table) NewRow() Row {
	r := make(Row)
	return r
}

// Append rows to the table
func (t *table) Append(row ...Row) {
	for _, r := range row {
		if r == nil {
			continue
		}
		for _, col := range t.cols {
			if col.IsComputed() {
				col.serie.Append(nil)
				continue
			}
			if cell, ok := r[col.Name()]; ok {
				col.serie.Append(cell)
				continue
			}
			col.serie.Grow(1)
		}
		t.nrows++
	}
	t.dirty = true
}

func (t *table) AppendRow(v ...interface{}) error {
	if len(v) != len(t.cols) {
		return errors.Errorf("length mismatch: expected %d elements, values have %d elements", len(t.cols), len(v))
	}

	for i, col := range t.cols {
		if col.IsComputed() {
			col.serie.Append(nil)
		} else {
			col.serie.Append(v[i])
		}
	}

	t.nrows++
	t.dirty = true

	return nil
}
