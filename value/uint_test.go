package value_test

import (
	"testing"

	"github.com/datasweet/datatable/value"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	val := value.NewUint(1)
	assert.Equal(t, value.Uint, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint(1), val.Val())
}

func TestUint64(t *testing.T) {
	val := value.NewUint64(1)
	assert.Equal(t, value.Uint64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(1), val.Val())

	val.Set(4294967296)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(4294967296), val.Val())
}

func TestUint32(t *testing.T) {
	val := value.NewUint32(1)
	assert.Equal(t, value.Uint32, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint32(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint32(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestUint16(t *testing.T) {
	val := value.NewUint16(1)
	assert.Equal(t, value.Uint16, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint16(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint16(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestUint8(t *testing.T) {
	val := value.NewUint8(1)
	assert.Equal(t, value.Uint8, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint8(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, uint8(123), val.Val())

	val.Set(4294967296) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneUint(t *testing.T) {
	val := value.NewUint64(4294967296)
	assert.NotNil(t, val)
	assert.Equal(t, value.Uint64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, uint64(4294967296), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.Uint64, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, uint64(4294967296), cpy.Val())
}

func TestCompareUint(t *testing.T) {
	a := value.NewUint64(123)
	assert.Equal(t, value.Eq, a.Compare(value.NewUint64(123)))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.NewUint64(0)))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.NewUint16(123)))
	assert.Equal(t, value.Eq, a.Compare(value.NewString("123")))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool("teemo")))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool(false)))

	a.Set(3)
	assert.Equal(t, value.Lt, a.Compare(value.NewUint64(123)))
	assert.Equal(t, value.Eq, a.Compare(value.NewUint16(3)))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
