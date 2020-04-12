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
	Clone(includeValues bool) Column
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

func (c *column) Clone(includeValues bool) Column {
	clone := &column{
		name:     c.name,
		hidden:   c.hidden,
		formulae: c.formulae,
		serie:    c.serie.Clone(includeValues),
	}
	if len(clone.formulae) > 0 {
		if parsed, err := expr.Parse(clone.formulae); err != nil {
			clone.expr = parsed
		}
	}
	return clone
}
