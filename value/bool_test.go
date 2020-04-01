package value_test

import (
	"testing"

	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	val := value.NewBool(1)
	assert.Equal(t, value.Bool, val.Type())
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
	val := value.NewBool(1)
	assert.NotNil(t, val)
	assert.Equal(t, value.Bool, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, true, val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.Bool, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, true, cpy.Val())
}

func TestCompare(t *testing.T) {
	a := value.NewBool(true)
	assert.Equal(t, value.Eq, a.Compare(value.NewBool(true)))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool(false)))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.NewInt(1)))
	assert.Equal(t, value.Eq, a.Compare(value.NewString("true")))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool("teemo")))

	a.Set(false)
	assert.Equal(t, value.Lt, a.Compare(value.NewBool(true)))
	assert.Equal(t, value.Eq, a.Compare(value.NewBool(false)))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
