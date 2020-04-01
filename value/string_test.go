package value_test

import (
	"testing"

	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	val := value.NewString(1)
	assert.Equal(t, value.String, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, "1", val.Val())

	val.Set(true)
	assert.True(t, val.IsValid())
	assert.Equal(t, "true", val.Val())

	val.Set("teemo")
	assert.True(t, val.IsValid())
	assert.Equal(t, "teemo", val.Val())
}

func TestCloneString(t *testing.T) {
	val := value.NewString(1)
	assert.NotNil(t, val)
	assert.Equal(t, value.String, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, "1", val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.String, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, "1", cpy.Val())
}
