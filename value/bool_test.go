package value_test

import (
	"testing"

	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	val := value.Bool()
	assert.Equal(t, value.BoolType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, false, val.Val())

	val = value.Bool(1)
	assert.Equal(t, value.BoolType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, true, val.Val())

	val.Set(0)
	assert.True(t, val.IsValid())
	assert.Equal(t, false, val.Val())

	val.Set("true")
	assert.True(t, val.IsValid())
	assert.Equal(t, true, val.Val())

	val.Set("teemo")
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneBool(t *testing.T) {
	val := value.Bool(1)
	assert.NotNil(t, val)
	assert.Equal(t, value.BoolType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, true, val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.BoolType, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, true, cpy.Val())
}

func TestCompare(t *testing.T) {
	a := value.Bool(true)
	assert.Equal(t, value.Eq, a.Compare(value.Bool(true)))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.Bool(false)))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.Int(1)))
	assert.Equal(t, value.Eq, a.Compare(value.String("true")))
	assert.Equal(t, value.Gt, a.Compare(value.Bool("teemo")))

	a.Set(false)
	assert.Equal(t, value.Lt, a.Compare(value.Bool(true)))
	assert.Equal(t, value.Eq, a.Compare(value.Bool(false)))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
