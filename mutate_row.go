package datatable

import (
	"github.com/pkg/errors"
)

type MutateOptions struct {
	UseZero    bool // UseZero to creates "zero" value instead
	KeepValues bool
}

type MutateOption func(opts *MutateOptions)

func (t *DataTable) NewRow() Row {
	r := make(Row)
	return r
}

// Append rows to the table
func (t *DataTable) Append(row ...Row) {
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
			//col.serie.Append(nil)
		}
		t.nrows++
	}
	t.dirty = true
}

// Append row to the table
func (t *DataTable) append(row Row, useZero bool) {
	for _, col := range t.cols {
		if col.IsComputed() {
			col.serie.Append(nil)
			continue
		}
		if row != nil {
			if cell, ok := row[col.name]; ok {
				col.serie.Append(cell)
				continue
			}
		}
		if useZero {
			col.serie.Grow(1)
			continue
		}
		col.serie.Append(nil)
	}
	t.nrows++
	t.dirty = true
}

func (t *DataTable) AppendRow(v ...interface{}) error {
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

// SwapRow in table
func (t *DataTable) SwapRow(i, j int) {
	for _, col := range t.cols {
		col.serie.Swap(i, j)
	}
}

func (t *DataTable) Grow(size int) {
	for _, col := range t.cols {
		col.serie.Grow(size)
	}
}

func (t *DataTable) Update(at int, row Row, opt ...MutateOption) error {
	options := MutateOptions{}
	for _, o := range opt {
		o(&options)
	}

	if row == nil {
		row = make(Row, 0)
	}

	for _, col := range t.cols {
		if col.IsComputed() {
			continue
		}
		cell, ok := row[col.name]
		if ok {
			col.serie.Update(at, cell)
			continue
		}
		if options.KeepValues {
			continue
		}
		if options.UseZero {
			// col.serie.Update(at, ) ZERO !!
			continue
		}
		col.serie.Update(at, nil)
		// NIL
	}

	return nil

}
