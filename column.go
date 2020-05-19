package datatable

import (
	"reflect"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
)

type ColumnOptions struct {
	Hidden bool
	Format string
}

type Column interface {
	Name() string
	Type() reflect.Type
	IsVisible() bool
	IsComputed() bool
	//Clone(includeValues bool) Column
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

func (c *column) Type() reflect.Type {
	return c.serie.Type()
}

func (c *column) IsVisible() bool {
	return !c.hidden
}

func (c *column) IsComputed() bool {
	return len(c.formulae) > 0
}

func (c *column) emptyCopy() *column {
	cpy := &column{
		name:     c.name,
		hidden:   c.hidden,
		formulae: c.formulae,
		serie:    c.serie.EmptyCopy(),
	}
	if len(cpy.formulae) > 0 {
		if parsed, err := expr.Parse(cpy.formulae); err == nil {
			cpy.expr = parsed
		}
	}
	return cpy
}

func (c *column) copy() *column {
	cpy := &column{
		name:     c.name,
		hidden:   c.hidden,
		formulae: c.formulae,
		serie:    c.serie.Copy(),
	}
	if len(cpy.formulae) > 0 {
		if parsed, err := expr.Parse(cpy.formulae); err == nil {
			cpy.expr = parsed
		}
	}
	return cpy
}
