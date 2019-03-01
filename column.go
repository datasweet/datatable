package datatable

// Column defines a column in table
type Column struct {
	Name       string
	ColumnType ColumnType
	Rows       []interface{}
}

// ColumnType defines the column type (used for formatter / check the Rows)
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
		Name:       name,
		ColumnType: ctyp,
	}
	return col
}

// Len returns the number of rows
func (c *Column) Len() int {
	return len(c.Rows)
}

// Size set the column size, ie the number of rows
// Extend or shrink the rows
func (c *Column) Size(size int) bool {
	if size < 0 {
		return false
	}

	lv := len(c.Rows)
	if lv < size {
		// extend
		c.Rows = append(c.Rows, make([]interface{}, size-lv)...)
	} else if lv > size {
		// shrink
		c.Rows = c.Rows[:size]
	}
	return true
}

// Set sets the rows
// If more values are provided than the number of rows in current column,
// the column is extended
func (c *Column) Set(values ...interface{}) bool {
	return c.SetAt(0, values...)
}

// SetAt sets the rows at index {at}
// If more values are provided than the number of rows in current column,
// the column is extended
func (c *Column) SetAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 {
		return false
	}

	// Extends if needed
	if max := at + nrows; max > len(c.Rows) {
		c.Size(max)
	}

	for i := 0; i < nrows; i++ {
		c.Rows[i+at] = c.cast(values[i])
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
		c.Rows = append(c.Rows, c.cast(v))
	}
	return true
}

// InsertAt insert rows at index {at}
func (c *Column) InsertAt(at int, values ...interface{}) bool {
	nrows := len(values)
	if nrows <= 0 || at < 0 || at > len(c.Rows) {
		return false
	}

	casted := make([]interface{}, nrows)
	for i, v := range values {
		casted[i] = c.cast(v)
	}

	c.Rows = append(c.Rows[:at], append(casted, c.Rows[at:]...)...)
	return true
}

// InsertEmpty insert {nrows} empty rows at index {at}
func (c *Column) InsertEmpty(at int, nrows int) bool {
	if nrows <= 0 || at < 0 || at > len(c.Rows) {
		return false
	}
	rows := make([]interface{}, nrows)
	c.Rows = append(c.Rows[:at], append(rows, c.Rows[at:]...)...)
	return true
}

// DeleteAt deletes the {n} rows at idx {from}
func (c *Column) DeleteAt(from, n int) bool {
	if n <= 0 || from < 0 || from+n > len(c.Rows) {
		return false
	}
	c.Rows = append(c.Rows[:from], c.Rows[from+n:]...)
	return true
}

// GetAt retrieves a row value in column at index {at}
func (c *Column) GetAt(at int) interface{} {
	if at < 0 || at >= len(c.Rows) {
		return nil
	}
	return c.Rows[at]
}

// Cast the value to the column type
// return nil if cast is wrong
func (c *Column) cast(v interface{}) interface{} {
	switch c.ColumnType {
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
