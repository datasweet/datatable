package datatable

import (
	"github.com/pkg/errors"
)

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
			if !col.IsComputed() {
				if cell, ok := r[col.Name()]; ok {
					col.serie.Append(cell)
					continue
				}
			}
			col.serie.Grow(1)
		}
		t.nrows++
	}
	t.dirty = true
}

func (t *DataTable) AppendRow(v ...interface{}) error {
	if len(v) != len(t.cols) {
		return errors.Errorf("length mismatch: expected %d elements, values have %d elements", len(t.cols), len(v))
	}

	for i, col := range t.cols {
		if col.IsComputed() {
			col.serie.Grow(1)
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

func (t *DataTable) Update(at int, row Row) error {
	if row == nil {
		row = make(Row, 0)
	}

	for _, col := range t.cols {
		if col.IsComputed() {
			continue
		}
		cell, ok := row[col.name]
		if ok {
			if err := col.serie.Set(at, cell); err != nil {
				return errors.Wrapf(err, "col %s", col.name)
			}
			continue
		}
		if err := col.serie.Set(at, nil); err != nil {
			return errors.Wrapf(err, "col %s", col.name)
		}
	}

	return nil
}
