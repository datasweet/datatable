package datatable

import (
	"github.com/datasweet/datatable/serie"
	"github.com/pkg/errors"
)

type Table interface {
	Name() string
	NCols() int
	NRows() int
}

type table struct {
	name  string
	cols  []serie.Serie
	nrows int
	cuid  int
	dirty bool
}

func NewTable(name string) Table {
	return &table{
		name:     name,
		cuid:     1,
		colnames: make(map[string]int),
	}
}

func (t *table) Name() string {
	return t.name
}

func (t *table) NCols() int {
	return len(t.cols)
}

func (t *table) NRows() int {
	return t.nrows
}

func (t *table) AddColumn(name string, s serie.Serie) error {
	if _, ok := t.colnames[name]; ok {
		return errors.Errorf("column '%s' already exists", name)
	}

	if s == nil {
		return errors.New("serie prodived id nil")
	}

	if s.Error() != nil {
		return errors.New("serie provided has error")
	}

	cpy := s.Clone()

	l := cpy.Len()

	// adjust size
	if l < t.nrows {
		for i := l; i < t.nrows; i++ {
			cpy.Append(nil)
		}
	} else {
		for _, col := range t.cols {
			// create nils values
			for i := t.nrows; i < l; i++ {
				col.Append(nil)
			}
		}
	}

	t.cols = append(t.cols, cpy)
	t.colnames[name] = len(t.cols) - 1
	t.nrows = cpy.Len()
	t.dirty = true

	return nil
}

func (t *table) Rows() []Row {
	// Row: map[string]interface{}
	// => we will the orders of cols.
	rows := make([]Row, t.nrows)

	for i := 0; i < t.nrows; i++ {
		r := make(Row)
		for name, pos := range t.colnames {
			r[name] = t.cols[pos].Value(i)
		}
		rows[i] = r
	}
	return rows
}

func (t *table) Rows()
