package value_test

import (
	"testing"

	"github.com/datasweet/datatable/value"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	val := value.NewInt(1)
	assert.Equal(t, value.Int, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, 1, val.Val())
}

func TestInt64(t *testing.T) {
	val := value.NewInt64(1)
	assert.Equal(t, value.Int64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(1), val.Val())

	val.Set(2147483648)
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(2147483648), val.Val())
}

func TestInt32(t *testing.T) {
	val := value.NewInt32(1)
	assert.Equal(t, value.Int32, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int32(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int32(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestInt16(t *testing.T) {
	val := value.NewInt16(1)
	assert.Equal(t, value.Int16, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int16(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int16(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestInt8(t *testing.T) {
	val := value.NewInt8(1)
	assert.Equal(t, value.Int8, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int8(1), val.Val())

	val.Set(123)
	assert.True(t, val.IsValid())
	assert.Equal(t, int8(123), val.Val())

	val.Set(2147483648) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneInt(t *testing.T) {
	val := value.NewInt64(2147483648)
	assert.NotNil(t, val)
	assert.Equal(t, value.Int64, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, int64(2147483648), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.Int64, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, int64(2147483648), cpy.Val())
}

func TestCompareInt(t *testing.T) {
	a := value.NewInt64(123)
	assert.Equal(t, value.Eq, a.Compare(value.NewInt64(123)))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.NewInt64(0)))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.NewInt16(123)))
	assert.Equal(t, value.Eq, a.Compare(value.NewString("123")))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool("teemo")))
	assert.Equal(t, value.Gt, a.Compare(value.NewBool(false)))

	a.Set(-123)
	assert.Equal(t, value.Lt, a.Compare(value.NewInt64(123)))
	assert.Equal(t, value.Eq, a.Compare(value.NewInt16(-123)))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
