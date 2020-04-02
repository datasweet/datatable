package value_test

import (
	"math"
	"testing"

	"github.com/datasweet/datatable/value"

	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	val := value.Float64(1)
	assert.Equal(t, value.Float64Type, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, 1.0, val.Val())

	val.Set(3.14)
	assert.True(t, val.IsValid())
	assert.Equal(t, 3.14, val.Val())
}

func TestFloat32(t *testing.T) {
	val := value.Float32(1)
	assert.Equal(t, value.Float32Type, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, float32(1), val.Val())

	val.Set(3.14)
	assert.True(t, val.IsValid())
	assert.Equal(t, float32(3.14), val.Val())

	val.Set(math.MaxFloat64) // overflow :)
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneFloat(t *testing.T) {
	val := value.Float64(3.14)
	assert.NotNil(t, val)
	assert.Equal(t, value.Float64Type, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, float64(3.14), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.Float64Type, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, float64(3.14), cpy.Val())
}

func TestCompareFloat(t *testing.T) {
	a := value.Float64(1.23)
	assert.Equal(t, value.Eq, a.Compare(value.Float64(1.23)))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.Float64(0)))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.Float32(1.23)))
	assert.Equal(t, value.Eq, a.Compare(value.String("1.23")))
	assert.Equal(t, value.Gt, a.Compare(value.Bool("teemo")))
	assert.Equal(t, value.Gt, a.Compare(value.Bool(false)))

	a.Set(-1.23)
	assert.Equal(t, value.Lt, a.Compare(value.Float64(1.23)))
	assert.Equal(t, value.Eq, a.Compare(value.Float32(-1.23)))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
