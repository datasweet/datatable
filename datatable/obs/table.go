package datatable

import (
	"github.com/datasweet/datatable/serie"
	"github.com/pkg/errors"
)

type Table interface {
	Name() string
	NCols() int
	NRows() int
	AddColumn(name string, s serie.Serie) error
	Columns() []string
	Rows() []Row
}

type table struct {
	name     string
	cols     []*column
	colnames map[string]int
	nrows    int
	cuid     int
	err      error
	dirty    bool
}

func New(name string) Table {
	return &table{
		name:     name,
		cuid:     1,
		colnames: make(map[string]int),
	}
}

func (t *table) Name() string {
	return t.name
}

// NCols returns the number of columns (visible or not) in the table
func (t *table) NCols() int {
	return len(t.cols)
}

// NRows return the number of rows in datatable
func (t *table) NRows() int {
	return t.nrows
}

func (t *table) AddColumn(name string, s serie.Serie) error {
	if _, ok := t.colnames[name]; ok {
		return errors.Errorf("column '%s' already exists", name)
	}
	if s == nil {
		return errors.New("serie provided id nil")
	}
	if s.Error() != nil {
		return errors.Wrap(s.Error(), "serie provided has error")
	}

	// options : don't clone ?
	cpy := s.Clone()

	l := cpy.Len()

	// adjust size
	if l < t.nrows {
		for i := l; i < t.nrows; i++ {
			cpy.Append(nil)
		}
	} else {
		for _, col := range t.cols {
			// create nils values: options zero values ?
			for i := t.nrows; i < l; i++ {
				col.serie.Append(nil)
			}
		}
	}

	t.cols = append(t.cols, newColumn(name, cpy))
	t.colnames[name] = len(t.cols) - 1
	t.nrows = cpy.Len()
	t.dirty = true
	return nil
}

func (t *table) Columns() []string {
	var cols []string
	for _, c := range t.cols {
		if c.IsVisible() {
			cols = append(cols, c)
		}
	}
	return cols
}

func (t *table) Rows() []Row {
	// Row: map[string]interface{}
	rows := make([]Row, t.nrows)
	cols := t.Columns()

	for i := 0; i < t.nrows; i++ {
		r := make(Row)
		for _, col := range t.cols {
			r[name] = t.cols[pos].serie.Value(i)
		}
		rows[i] = r
	}
	return rows
}
