package datatable

import (
	"fmt"

	"github.com/pkg/errors"
)

// Union concats 2 datatables
// Union removes duplicate rows
// This method is must slower than unionall
func Union(left, right DataTable) (DataTable, error) {
	if left == nil {
		return nil, errors.New("left is nil datatable")
	}
	if right == nil {
		return nil, errors.New("right is nil datatable")
	}

	lcols := left.Columns()
	rcols := right.Columns()

	ll := len(lcols)
	lr := len(rcols)

	if ll != lr {
		return nil, errors.New("wrong number of columns between left and right")
	}

	dt := New(fmt.Sprintf("union-%s-%s", left.Name(), right.Name()))

	// Create columns
	for i := 0; i < ll; i++ {
		lc := lcols[i]
		rc := rcols[i]

		if lc.Name() != rc.Name() {
			return nil, errors.Errorf("wrong column name between left '%s' and right '%s' at index %d", lc.Name(), rc.Name(), i)
		}

		if lc.Type() != rc.Type() {
			return nil, errors.Errorf("wrong column type between left '%v' and right '%v' at index %d", lc.Type(), rc.Type(), i)
		}

		if lc.Expr() != rc.Expr() {
			return nil, errors.Errorf("wrong column expr between left '%s' and right '%s' at index %d", lc.Expr(), rc.Expr(), i)
		}

		ctyp := lc.Type()
		if lc.Computed() {
			ctyp = Raw
		}
		if _, err := dt.AddColumn(lc.Name(), ctyp); err != nil {
			return nil, err
		}
	}

	// Copy rows
	cache := make(map[uint64]bool, 0)
	rows := append(left.Rows(), right.Rows()...)
	for i, r := range rows {
		// Create hash
		h, ok := r.Hash()
		if !ok {
			return nil, errors.Errorf("invalid row at index %d", i)
		}

		// Already exists
		if _, exists := cache[h]; exists {
			continue
		}

		if !dt.AddRow(r) {
			return nil, errors.Errorf("can't add row at index %d", i)
		}

		// Add to cache
		cache[h] = true
	}

	return dt, nil
}

// UnionAll concats 2 datatables
// UnionAll does not remove duplicate rows.
func UnionAll(left, right DataTable) (DataTable, error) {
	if left == nil {
		return nil, errors.New("left is nil datatable")
	}
	if right == nil {
		return nil, errors.New("right is nil datatable")
	}

	lcols := left.Columns()
	rcols := right.Columns()

	ll := len(lcols)
	lr := len(rcols)

	if ll != lr {
		return nil, errors.New("wrong number of columns between left and right")
	}

	dt := New(fmt.Sprintf("unionall-%s-%s", left.Name(), right.Name()))

	for i := 0; i < ll; i++ {
		lc := lcols[i]
		rc := rcols[i]

		if lc.Name() != rc.Name() {
			return nil, errors.Errorf("wrong column name between left '%s' and right '%s' at index %d", lc.Name(), rc.Name(), i)
		}

		if lc.Type() != rc.Type() {
			return nil, errors.Errorf("wrong column type between left '%v' and right '%v' at index %d", lc.Type(), rc.Type(), i)
		}

		if lc.Expr() != rc.Expr() {
			return nil, errors.Errorf("wrong column expr between left '%s' and right '%s' at index %d", lc.Expr(), rc.Expr(), i)
		}

		ctyp := lc.Type()
		if lc.Computed() {
			ctyp = Raw
		}
		if _, err := dt.AddColumn(lc.Name(), ctyp); err != nil {
			return nil, err
		}
	}

	// copy rows
	dt.AddRow(left.Rows()...)
	dt.AddRow(right.Rows()...)

	return dt, nil
}
