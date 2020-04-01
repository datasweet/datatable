package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestBoolValue(t *testing.T) {
	val := serie.NewBoolValue(1)
	assert.Equal(t, serie.Bool, val.Type())
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

func TestCloneBoolValue(t *testing.T) {
	val := serie.NewBoolValue(1)
	assert.NotNil(t, val)
	assert.Equal(t, serie.Bool, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, true, val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, serie.Bool, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, true, cpy.Val())
}
