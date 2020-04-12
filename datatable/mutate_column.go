package datatable

import (
	"strings"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

func (t *DataTable) addColumn(name string, serie serie.Serie, formulae string) error {
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

	serie = serie.Clone(true)
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

func (t *DataTable) AddColumn(name string, serie serie.Serie) error {
	return t.addColumn(name, serie, "")
}

func (t *DataTable) AddExprColumn(name string, typ serie.Serie, formulae string) error {
	// var s serie.Serie
	// switch typ {
	// case ExprBool:
	// 	s = serie.Bool()
	// case ExprNumber:
	// 	s = serie.Float64()
	// case ExprString:
	// 	s = serie.String()
	// case ExprRaw:
	// 	s = serie.Raw()
	// default:
	// 	return errors.Errorf("unknown expression type %v", typ)
	// }

	return t.addColumn(name, typ.Clone(false), formulae)
}

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
