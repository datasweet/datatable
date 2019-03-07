package datatable

import (
	"github.com/datasweet/datatable/expr"
)

func (t *DataTable) AddFormulaColumn(name string, ctyp ColumnType, formulae string) error {
	parsed, err := expr.Parse(formulae)
	if err != nil {
		return err
	}

	params := make(map[string]interface{}, len(t.cols)-1)

	for _, c := range t.cols {
		if c.name == name {
			continue
		}
		params[c.name] = c.rows
	}

	res, err := parsed.Eval(params)

	if arr, ok := res.([]interface{}); ok {
		t.AddColumn(name, ctyp, arr...)
	} else {
		t.AddColumn(name, ctyp, res)
	}

	return nil
}
