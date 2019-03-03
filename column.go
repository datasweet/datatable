package datatable

// Column is a column in our datatable
// A column contains all rows
type Column struct {
	name  string
	ctype ColumnType
	rows  []interface{}
}

// ColumnType defines the Column type (used for formatter / check the Rows)
type ColumnType string

const (
	Bool     ColumnType = "bool"
	String   ColumnType = "string"
	Number   ColumnType = "number"
	Datetime ColumnType = "datetime"
)

// NewColumn creates a column
func NewColumn(name string, ctyp ColumnType) *Column {
	col := &Column{
		name:  name,
		ctype: ctyp,
	}
	return col
}

// Name returns name of string
func (c *Column) Name() string {
	return c.name
}

// Columns returns the binded column type
func (c *Column) ColumnType() ColumnType {
	return c.ctype
}

// Rows returns rows in Column
func (c *Column) Rows() []interface{} {
	return c.rows
}

// Len returns the number of rows
func (c *Column) Len() int {
	return len(c.rows)
}

// Size set the Column size, ie the number of rows
// Extend or shrink the rows
func (c *Column) Size(size int) bool {
	if size < 0 {
		return false
	}

	lv := len(c.rows)
	if lv < size {
		// extend
		c.rows = append(c.rows, make([]interface{}, size-lv)...)
	} else if lv > size {
		// shrink
		c.rows = c.rows[:size]
	}
	return true
}

// Set sets the rows
// If more values are provided than the number of rows in current Column,
// the Column is extended
func (c *Column) Set(values ...interface{}) bool {
	return c.SetAt(0, values...)
}

// SetAt sets the rows at index {at}
// If more values are provided than the number of rows in current Column,
// the Column is extended
func (c *Column) SetAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 {
		return false
	}

	// Extends if needed
	if max := at + nrows; max > len(c.rows) {
		c.Size(max)
	}

	for i := 0; i < nrows; i++ {
		c.rows[i+at] = c.cast(values[i])
	}
	return true
}

// Append add a new row at the end
func (c *Column) Append(values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 {
		return false
	}

	for _, v := range values {
		c.rows = append(c.rows, c.cast(v))
	}
	return true
}

// InsertAt insert rows at index {at}
func (c *Column) InsertAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 || at > len(c.rows) {
		return false
	}

	casted := make([]interface{}, nrows)
	for i, v := range values {
		casted[i] = c.cast(v)
	}

	c.rows = append(c.rows[:at], append(casted, c.rows[at:]...)...)
	return true
}

// InsertEmpty insert {nrows} empty rows at index {at}
func (c *Column) InsertEmpty(at int, nrows int) bool {
	if nrows <= 0 || at < 0 || at > len(c.rows) {
		return false
	}
	rows := make([]interface{}, nrows)
	c.rows = append(c.rows[:at], append(rows, c.rows[at:]...)...)
	return true
}

// DeleteAt deletes the {n} rows at idx {from}
func (c *Column) DeleteAt(from, n int) bool {
	if n <= 0 || from < 0 || from+n > len(c.rows) {
		return false
	}
	c.rows = append(c.rows[:from], c.rows[from+n:]...)
	return true
}

// GetAt retrieves a row value in Column at index {at}
func (c *Column) GetAt(at int) interface{} {
	if at < 0 || at >= len(c.rows) {
		return nil
	}
	return c.rows[at]
}

// Cast the value to the Column type
// return nil if cast is wrong
func (c *Column) cast(v interface{}) interface{} {
	switch c.ctype {
	case Bool:
		if casted, ok := AsBool(v); ok {
			return casted
		}
	case Number:
		if casted, ok := AsNumber(v); ok {
			return casted
		}
	case String:
		if casted, ok := AsString(v); ok {
			return casted
		}
	case Datetime:
		if casted, ok := AsDatetime(v); ok {
			return casted
		}
	}
	return nil
}
