package column

import "github.com/datasweet/datatable/value"

// Column defines a column in our datatable
type Column interface {
	Name() string
	Type() value.Type
	IsVisible() bool
}

// ColumnValue defines a 'value' column
type ColumnValue interface {
	Column
	NewValue() value.Value
}

// ColumnExpr defines a 'computed' column
type ColumnExpr interface {
	Column
	Expr() string
}

type column struct {
	name   string
	typ    value.Type
	hidden bool
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Type() value.Type {
	return c.typ
}

func (c *column) IsVisible() bool {
	return !c.hidden
}
