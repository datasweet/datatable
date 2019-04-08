package datatable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStringColumn(t *testing.T) {
	col := newColumn("test", String)
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, String, col.Type())
	assert.False(t, col.Computed())
	assert.Equal(t, "", col.Expr())
	assert.Equal(t, "", col.ZeroValue())
	assert.Equal(t, nil, col.AsValue(nil))
	assert.Equal(t, "-1", col.AsValue(-1))
	assert.Equal(t, "3.14", col.AsValue(3.14))
	assert.Equal(t, "hello", col.AsValue("hello"))
}

func TestNewIntColumn(t *testing.T) {
	col := newColumn("test", Int)
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, Int, col.Type())
	assert.False(t, col.Computed())
	assert.Equal(t, "", col.Expr())
	assert.Equal(t, int64(0), col.ZeroValue())
	assert.Equal(t, nil, col.AsValue(nil))
	assert.Equal(t, int64(-1), col.AsValue(-1))
	assert.Equal(t, int64(3), col.AsValue(3.14))
	assert.Equal(t, nil, col.AsValue("hello"))
}

func TestNewBoolColumn(t *testing.T) {
	col := newColumn("test", Bool)
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, Bool, col.Type())
	assert.False(t, col.Computed())
	assert.Equal(t, "", col.Expr())
	assert.Equal(t, false, col.ZeroValue())
	assert.Equal(t, nil, col.AsValue(nil))
	assert.Equal(t, false, col.AsValue(-1))
	assert.Equal(t, true, col.AsValue(3.14))
	assert.Equal(t, nil, col.AsValue("hello"))
}

func TestNewDatetimeColumn(t *testing.T) {
	jstimestamp := int64(1551435220270)
	date := time.Unix(0, jstimestamp*int64(time.Millisecond)).UTC()

	col := newColumn("test", Datetime)
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, Datetime, col.Type())
	assert.False(t, col.Computed())
	assert.Equal(t, "", col.Expr())
	assert.Equal(t, nil, col.ZeroValue())
	assert.Equal(t, nil, col.AsValue(nil))
	assert.Equal(t, date, col.AsValue(1551435220270))
	assert.Equal(t, date, col.AsValue("1551435220270"))
	assert.Equal(t, date, col.AsValue("2019-03-01T10:13:40.27Z"))
	assert.Equal(t, nil, col.AsValue("hello"))
}

func TestNewExprColumn(t *testing.T) {
	col, err := newExprColumn("test", "sum(win) * 100 / (sum(win) + sum(loose))")
	assert.NoError(t, err)
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, Raw, col.Type())
	assert.True(t, col.Computed())
	assert.Equal(t, "sum(win) * 100 / (sum(win) + sum(loose))", col.Expr())
}
