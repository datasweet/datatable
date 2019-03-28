package datatable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	col := newColumn("test", String)

	// Extend
	assert.True(t, col.Size(3))
	rowsEq(t, col, nil, nil, nil)

	assert.True(t, col.SetAt(1, "test"))
	rowsEq(t, col, nil, "test", nil)

	// Shrink
	assert.True(t, col.Size(2))
	rowsEq(t, col, nil, "test")

	// Invalid
	assert.False(t, col.Size(-1))
	rowsEq(t, col, nil, "test")
}

func TestSet(t *testing.T) {
	col := newColumn("test", String)

	// Invalid
	assert.False(t, col.Set())

	// Auto extend
	assert.True(t, col.Set("hello", "world"))
	rowsEq(t, col, "hello", "world")

	// Rewrite
	assert.True(t, col.Set("bonjour", "monde"))
	rowsEq(t, col, "bonjour", "monde")
}

func TestSetAt(t *testing.T) {
	col := newColumn("test", String)

	// Invalid
	assert.False(t, col.SetAt(0))
	assert.False(t, col.SetAt(-1, "hello", "world"))

	// Auto extend
	assert.True(t, col.SetAt(3, "hello", "world"))
	rowsEq(t, col, nil, nil, nil, "hello", "world")

	// Rewrite
	assert.True(t, col.SetAt(1, "bonjour", "monde"))
	rowsEq(t, col, nil, "bonjour", "monde", "hello", "world")
}

func TestSetBool(t *testing.T) {
	col := newColumn("test", Bool)

	col.SetAt(0, true, false, "true", "false", 1, 0, "wrong")
	rowsEq(t, col, true, false, true, false, true, false, nil)

	col.Size(2)
	rowsEq(t, col, true, false)

	col.InsertAt(2, true, false, "true", "false", 1, 0, "wrong")
	rowsEq(t, col, true, false, true, false, true, false, true, false, nil)
}

func TestSetString(t *testing.T) {
	col := newColumn("test", String)

	col.SetAt(0, true, false, 12345, 3.14, "hello", nil)
	rowsEq(t, col, "true", "false", "12345", "3.14", "hello", nil)

	col.Size(0)
	rowsEq(t, col)

	col.InsertAt(0, true, false, 12345, 3.14, "hello", nil)
	rowsEq(t, col, "true", "false", "12345", "3.14", "hello", nil)
}

func TestSetNumber(t *testing.T) {
	col := newColumn("test", Number)

	col.SetAt(0, true, false, 12345, 3.14, "12345", "3.14", "hello", nil)
	rowsEq(t, col, float64(1), float64(0), float64(12345), float64(3.14), float64(12345), float64(3.14), nil, nil)

	col.Size(0)
	rowsEq(t, col)

	col.InsertAt(0, true, false, 12345, 3.14, "12345", "3.14", "hello", nil)
	rowsEq(t, col, float64(1), float64(0), float64(12345), float64(3.14), float64(12345), float64(3.14), nil, nil)
}

func TestSetDatetime(t *testing.T) {
	jstimestamp := int64(1551435220270)
	date := time.Unix(0, jstimestamp*int64(time.Millisecond)).UTC()

	col := newColumn("test", Datetime)

	col.SetAt(0, true, false, 1551435220270, "1551435220270", "2019-03-01T10:13:40.27Z", "hello", nil)
	rowsEq(t, col, nil, nil, date, date, date, nil, nil)

	col.Size(0)
	rowsEq(t, col)

	col.InsertAt(0, true, false, 1551435220270, "1551435220270", "2019-03-01T10:13:40.27Z", "hello", nil)
	rowsEq(t, col, nil, nil, date, date, date, nil, nil)
}

func TestInsertAt(t *testing.T) {
	col := newColumn("test", String)

	col.Set("hello", "world")
	rowsEq(t, col, "hello", "world")

	// Invalid
	assert.False(t, col.InsertAt(1))
	assert.False(t, col.InsertAt(-1, "bonjour"))
	assert.False(t, col.InsertAt(3, "bonjour"))

	assert.True(t, col.InsertAt(1, "go", "land"))
	rowsEq(t, col, "hello", "go", "land", "world")

	assert.True(t, col.InsertAt(0, "prepend"))
	rowsEq(t, col, "prepend", "hello", "go", "land", "world")

	assert.True(t, col.InsertAt(col.Len(), "append"))
	rowsEq(t, col, "prepend", "hello", "go", "land", "world", "append")
}

func TestAppend(t *testing.T) {
	col := newColumn("test", String)

	// Invalid
	assert.False(t, col.Append())

	assert.True(t, col.Append("hello", "world", "!"))
	rowsEq(t, col, "hello", "world", "!")
}

func TestDeleteAt(t *testing.T) {
	col := newColumn("test", String)

	col.Set("hello", "happy", "go", "land", "world", "!")
	rowsEq(t, col, "hello", "happy", "go", "land", "world", "!")

	// Invalid
	assert.False(t, col.DeleteAt(1, 0))
	assert.False(t, col.DeleteAt(-1, 2))
	assert.False(t, col.DeleteAt(1, 10))

	assert.True(t, col.DeleteAt(2, 2))
	rowsEq(t, col, "hello", "happy", "world", "!")

	assert.True(t, col.DeleteAt(0, 1))
	rowsEq(t, col, "happy", "world", "!")

	assert.True(t, col.DeleteAt(col.Len()-1, 1))
	rowsEq(t, col, "happy", "world")
}

func TestGetAt(t *testing.T) {
	col := newColumn("test", String)

	col.Set("hello", "happy", "world")
	rowsEq(t, col, "hello", "happy", "world")

	assert.Nil(t, col.GetAt(-1))
	assert.Equal(t, "hello", col.GetAt(0))
	assert.Equal(t, "happy", col.GetAt(1))
	assert.Equal(t, "world", col.GetAt(2))
	assert.Nil(t, col.GetAt(3))
}

func TestInsertEmpty(t *testing.T) {
	col := newColumn("test", String)

	col.Set("hello", "world")
	rowsEq(t, col, "hello", "world")

	// Invalid
	assert.False(t, col.InsertEmpty(1, 0))
	assert.False(t, col.InsertEmpty(-1, 2))
	assert.False(t, col.InsertEmpty(3, 2))

	assert.True(t, col.InsertEmpty(1, 2))
	rowsEq(t, col, "hello", nil, nil, "world")

	assert.True(t, col.InsertEmpty(0, 1))
	rowsEq(t, col, nil, "hello", nil, nil, "world")

	assert.True(t, col.InsertEmpty(col.Len(), 1))
	rowsEq(t, col, nil, "hello", nil, nil, "world", nil)
}
