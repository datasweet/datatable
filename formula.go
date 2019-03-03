package datatable

import (
	"github.com/Knetic/govaluate"
	"github.com/datasweet/datatable/formula"
)

func (t *DataTable) AddFormulaColumn(name string, ctyp ColumnType, formulae string) error {
	expr, err := govaluate.NewEvaluableExpressionWithFunctions(formulae, formula.All)
	if err != nil {
		return err
	}

	calculated := make([]interface{}, t.nrows)

	for r := 0; r < t.nrows; r++ {
		params := make(map[string]interface{}, len(t.cols)-1)

		for _, c := range t.cols {
			if c.Name() == name {
				continue
			}
			params[c.Name()] = c.GetAt(r)
		}

		if res, err := expr.Evaluate(params); err == nil {
			calculated[r] = res
		}
	}

	t.AddColumn(name, ctyp, calculated...)

	return nil
}
