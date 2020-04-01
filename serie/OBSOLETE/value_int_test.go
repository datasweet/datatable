package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestIntValue(t *testing.T) {
	const wordSize = 32 << (^uint(0) >> 32 & 1)
	val := serie.NewIntValue(1)

	switch wordSize {
	case 32:
		assert.Equal(t, serie.Int32, val.Type())
		assert.True(t, val.IsValid())
		assert.Equal(t, int32(1), val.Val())
	case 64:
		assert.Equal(t, serie.Int64, val.Type())
		assert.True(t, val.IsValid())
		assert.Equal(t, int64(1), val.Val())
	default:
		t.Fatalf("unknown word size %d", wordSize)
	}

	val.Set(2147483648)

	switch wordSize {
	case 32:
		// overflow
		assert.False(t, val.IsValid())
		assert.Equal(t, nil, val.Val())

	case 64:
		assert.True(t, val.IsValid())
		assert.Equal(t, int64(2147483648), val.Val())
	}

}

func TestInt64Value(t *testing.T) {
	val := serie.NewInt64Value(1)
	assert.Equal(t, serie.Int64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(1), val.Val())

	val.Set(2147483648)
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(2147483648), val.Val())
}

func TestInt32Value(t *testing.T) {
	val := serie.NewInt32Value(1)
	assert.Equal(t, serie.Int32, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int32(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int32(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestInt16Value(t *testing.T) {
	val := serie.NewInt16Value(1)
	assert.Equal(t, serie.Int16, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int16(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int16(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestInt8Value(t *testing.T) {
	val := serie.NewInt8Value(1)
	assert.Equal(t, serie.Int8, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int8(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int8(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneIntValue(t *testing.T) {
	val := serie.NewInt64Value(2147483648)
	assert.NotNil(t, val)
	assert.Equal(t, serie.Int64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(2147483648), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, serie.Int64, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, int64(2147483648), cpy.Val())
}
