package datatable

import "github.com/pkg/errors"

// evaluateExpressions to evaluate all columns with a binded expression
func (t *table) evaluateExpressions() error {
	var cols []int
	var exprCols []int
	for i, c := range t.cols {
		if c.IsComputed() {
			exprCols = append(exprCols, i)
		} else {
			cols = append(cols, i)
		}
	}

	l := len(exprCols)
	if l == 0 {
		t.dirty = false
		return nil
	}

	// Initialize params
	params := make(map[string][]interface{}, len(t.cols))
	for _, pos := range cols {
		col := t.cols[pos]
		params[col.name] = col.serie.Values()
	}

	// Evaluate
	for _, idx := range exprCols {
		col := t.cols[idx]
		res, err := col.expr.Eval(params)
		if err != nil {
			return err
		}

		name := col.Name()

		if arr, ok := res.([]interface{}); ok {
			// Is array
			ls := col.serie.Len()
			la := len(arr)

			if t.nrows != ls || la != ls {
				return errors.Errorf("evaluate expr : size mismatch %d vs %d", la, ls)
			}

			for i := 0; i < t.nrows; i++ {
				col.serie.Update(i, arr[i])
			}

		} else {
			// Is scalar
			for i := 0; i < t.nrows; i++ {
				col.serie.Update(i, res)
			}
		}

		if err := col.serie.Error(); err != nil {
			return errors.Wrap(err, "evaluate expr failed")
		}

		// update dependency
		params[name] = col.serie.Values()

	}

	t.dirty = false

	return nil
}
