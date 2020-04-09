package datatable

import (
	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/datatable/value"
	"github.com/datasweet/expr"
)

type ColumnOptions struct {
	Hidden bool
	Format string
}

type Column interface {
	Name() string
	Type() value.Type
	IsVisible() bool
	IsComputed() bool
}

type column struct {
	name     string
	hidden   bool
	formulae string
	expr     expr.Node
	serie    serie.Serie
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Type() value.Type {
	return c.serie.Type()
}

func (c *column) IsVisible() bool {
	return !c.hidden
}

func (c *column) IsComputed() bool {
	return len(c.formulae) > 0
}
