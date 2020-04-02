package column

import (
	"github.com/datasweet/datatable/value"
	"github.com/datasweet/expr"
)

type exprColumn struct {
	*column
	formulae string
	expr     expr.Node
}

func newExprColumn(name, formulae string) (*exprColumn, error) {
	parsed, err := expr.Parse(formulae)
	if err != nil {
		return nil, err
	}
	return &exprColumn{
		column:   &column{name: name, typ: value.Raw},
		formulae: formulae,
		expr:     parsed,
	}, nil
}

func (c *exprColumn) Expr() string {
	return c.formulae
}
