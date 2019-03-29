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
	Len() int
	Rows() []interface{}
	GetAt(at int) interface{}
	IsExpr() bool
}

// column is a column in our datatable
// A column contains all rows
type column struct {
	name   string
	ctype  ColumnType
	rows   []interface{}
	expr   expr.Node
	hidden bool
	label  string
}

// ColumnType defines the column type (used for formatter / check the Rows)
type ColumnType string

const (
	Bool     ColumnType = "bool"
	String   ColumnType = "string"
	Number   ColumnType = "number"
	Datetime ColumnType = "datetime"
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
func newExprColumn(name string, ctyp ColumnType, expr expr.Node) *column {
	return &column{
		name:  name,
		ctype: ctyp,
		expr:  expr,
	}
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

// Len returns the number of rows
func (c *column) Len() int {
	return len(c.rows)
}

// Rows returns rows in column
func (c *column) Rows() []interface{} {
	return c.rows
}

// IsExpr to know if the column is an expression column,
// ie a calculated column
func (c *column) IsExpr() bool {
	return c.expr != nil
}

// Size set the column size, ie the number of rows
// Extend (fill with zero values) or shrink the rows
func (c *column) Size(size int) bool {
	if size < 0 {
		return false
	}

	lv := len(c.rows)
	if lv < size {
		// extend
		c.rows = append(c.rows, c.zeroValues(size-lv)...)
	} else if lv > size {
		// shrink
		c.rows = c.rows[:size]
	}
	return true
}

// Set sets the rows
// If more values are provided than the number of rows in current column,
// the column is extended
func (c *column) Set(values ...interface{}) bool {
	return c.SetAt(0, values...)
}

// SetAt sets the rows at index {at}
// If more values are provided than the number of rows in current column,
// the column is extended
func (c *column) SetAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 {
		return false
	}

	// Extends or shrink if needed
	if max := at + nrows; max > len(c.rows) {
		c.Size(max)
	}

	for i := 0; i < nrows; i++ {
		c.rows[i+at] = c.asValue(values[i])
	}
	return true
}

// Append add a new row at the end
func (c *column) Append(values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 {
		return false
	}

	for _, v := range values {
		c.rows = append(c.rows, c.asValue(v))
	}
	return true
}

// InsertAt insert rows at index {at}
func (c *column) InsertAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 || at > len(c.rows) {
		return false
	}

	casted := make([]interface{}, nrows)
	for i, v := range values {
		casted[i] = c.asValue(v)
	}

	c.rows = append(c.rows[:at], append(casted, c.rows[at:]...)...)
	return true
}

// InsertEmpty insert {nrows} empty rows at index {at}
func (c *column) InsertEmpty(at int, nrows int) bool {
	if nrows <= 0 || at < 0 || at > len(c.rows) {
		return false
	}
	c.rows = append(c.rows[:at], append(c.zeroValues(nrows), c.rows[at:]...)...)
	return true
}

// DeleteAt deletes the {n} rows at idx {from}
func (c *column) DeleteAt(from, n int) bool {
	if n <= 0 || from < 0 || from+n > len(c.rows) {
		return false
	}
	c.rows = append(c.rows[:from], c.rows[from+n:]...)
	return true
}

// GetAt retrieves a row value in column at index {at}
func (c *column) GetAt(at int) interface{} {
	if at < 0 || at >= len(c.rows) {
		return nil
	}
	return c.rows[at]
}

// asValue cast the value to the column type
// return nil if cast is wrong
func (c *column) asValue(v interface{}) interface{} {
	switch c.ctype {
	case Bool:
		if casted, ok := cast.AsBool(v); ok {
			return casted
		}
	case Number:
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
	}
	return nil
}

func (c *column) ZeroValue() interface{} {
	switch c.ctype {
	case Bool:
		return false
	case Number:
		return float64(0)
	case String:
		return ""
	default:
		return nil
	}
}

func (c *column) zeroValues(n int) []interface{} {
	zero := c.ZeroValue()
	zv := make([]interface{}, n)
	for i := range zv {
		zv[i] = zero
	}
	return zv
}
