package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestUintValue(t *testing.T) {
	const wordSize = 32 << (^uint(0) >> 32 & 1)
	val := serie.NewUintValue(1)

	switch wordSize {
	case 32:
		assert.Equal(t, serie.Uint32, val.Type())
		assert.True(t, val.IsValid())
		assert.Equal(t, uint32(1), val.Val())
	case 64:
		assert.Equal(t, serie.Uint64, val.Type())
		assert.True(t, val.IsValid())
		assert.Equal(t, uint64(1), val.Val())
	default:
		t.Fatalf("unknown word size %d", wordSize)
	}

	val.Set(4294967296)

	switch wordSize {
	case 32:
		// overflow
		assert.False(t, val.IsValid())
		assert.Equal(t, nil, val.Val())

	case 64:
		assert.True(t, val.IsValid())
		assert.Equal(t, uint64(4294967296), val.Val())
	}

}

func TestUint64Value(t *testing.T) {
	val := serie.NewUint64Value(1)
	assert.Equal(t, serie.Uint64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(1), val.Val())

	val.Set(4294967296)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(4294967296), val.Val())
}

func TestUint32Value(t *testing.T) {
	val := serie.NewUint32Value(1)
	assert.Equal(t, serie.Uint32, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint32(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint32(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestUint16Value(t *testing.T) {
	val := serie.NewUint16Value(1)
	assert.Equal(t, serie.Uint16, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint16(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint16(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestUint8Value(t *testing.T) {
	val := serie.NewUint8Value(1)
	assert.Equal(t, serie.Uint8, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint8(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint8(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneUintValue(t *testing.T) {
	val := serie.NewUint64Value(2147483648)
	assert.NotNil(t, val)
	assert.Equal(t, serie.Uint64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(2147483648), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, serie.Uint64, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, uint64(2147483648), cpy.Val())
}
