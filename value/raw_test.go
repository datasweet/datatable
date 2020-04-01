package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/value"
)

type champion struct {
	Name string
	Win  int
}

func TestRawValue(t *testing.T) {
	val := value.NewRaw(champion{"Teemo", 100}, champion{})
	assert.NotNil(t, val)
	assert.Equal(t, value.Type("champion"), val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, champion{"Teemo", 100}, val.Val())

	val.Set(1)
	assert.False(t, val.IsValid())
	assert.Nil(t, val.Val())
}

func TestCloneRawValue(t *testing.T) {
	val := value.NewRaw(champion{"Teemo", 100}, champion{})
	assert.NotNil(t, val)
	assert.Equal(t, value.Type("champion"), val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, champion{"Teemo", 100}, val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.Type("champion"), cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.NotSame(t, val.Val(), cpy.Val())
	assert.Equal(t, champion{"Teemo", 100}, cpy.Val())
}
