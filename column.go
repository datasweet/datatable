package datatable

import (
	"strings"

	"github.com/datasweet/cast"
	"github.com/datasweet/expr"
)

// DataColumn defines a column in our datatable
type DataColumn interface {
	Name() string
	Type() ColumnType
	Label(v ...string) string // getter / setter
	Hidden(v ...bool) bool    // getter / setter
	Computed() bool
	Expr() string
	AsValue(v interface{}) interface{}
	ZeroValue() interface{}
}

// column is a column in our datatable
// A column contains all rows
type column struct {
	name     string
	ctype    ColumnType
	formulae string
	expr     expr.Node
	hidden   bool
	label    string
}

// ColumnType defines the column type (used for formatter / check the Rows)
type ColumnType string

const (
	Bool     ColumnType = "bool"
	String   ColumnType = "string"
	Int      ColumnType = "int"
	Float    ColumnType = "float"
	Datetime ColumnType = "datetime"
	Raw      ColumnType = "raw"
)

// newColumn to create a column
func newColumn(name string, ctyp ColumnType) *column {
	col := &column{
		name:  name,
		ctype: ctyp,
	}
	return col
}

// newExprColumn to create a column with a binded expression
func newExprColumn(name, formulae string) (*column, error) {
	parsed, err := expr.Parse(formulae)
	if err != nil {
		return nil, err
	}

	return &column{
		name:     name,
		ctype:    Raw,
		formulae: formulae,
		expr:     parsed,
	}, nil
}

// Name returns name of string
func (c *column) Name() string {
	return c.name
}

// Columns returns the binded column type
func (c *column) Type() ColumnType {
	return c.ctype
}

// Label sets / gets the label of our column
// If no label, will return the Name()
func (c *column) Label(v ...string) string {
	if l := len(v); l == 1 {
		c.label = strings.TrimSpace(v[0])
	}
	if len(c.label) > 0 {
		return c.label
	}
	return c.name
}

// Hidden sets / gets if the col will be exported
func (c *column) Hidden(v ...bool) bool {
	if l := len(v); l == 1 {
		c.hidden = v[0]
	}
	return c.hidden
}

// Computed to know if the column is a computed column
func (c *column) Computed() bool {
	return c.expr != nil
}

// Expr returns the expression formulae used by column
func (c *column) Expr() string {
	return c.formulae
}

// AsValue cast the value to the column type
// return nil value if cast is wrong
func (c *column) AsValue(v interface{}) interface{} {
	switch c.ctype {
	case Bool:
		if casted, ok := cast.AsBool(v); ok {
			return casted
		}
	case Int:
		if casted, ok := cast.AsInt(v); ok {
			return casted
		}
	case Float:
		if casted, ok := cast.AsFloat(v); ok {
			return casted
		}
	case String:
		if casted, ok := cast.AsString(v); ok {
			return casted
		}
	case Datetime:
		if casted, ok := cast.AsDatetime(v); ok {
			return casted
		}
	case Raw:
		return v
	}
	return nil
}

// ZeroValue is a default value for a column type
func (c *column) ZeroValue() interface{} {
	switch c.ctype {
	case Bool:
		return false
	case Int:
		return int64(0)
	case Float:
		return float64(0)
	case String:
		return ""
	default:
		return nil
	}
}
