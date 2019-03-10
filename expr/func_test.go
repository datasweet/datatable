package expr

import (
	"testing"

	"github.com/datasweet/datatable/cast"
	"github.com/stretchr/testify/assert"
)

func TestBinaryOperatorFunc(t *testing.T) {
	var dummy binaryOperatorFunc = func(x, y interface{}) interface{} {
		cx, _ := cast.AsInt(x)
		cy, _ := cast.AsInt(y)
		return cx + cy
	}

	assert.Equal(t, int64(3), dummy(1, 2))

	// scalar
	assert.Equal(t, int64(3), dummy.Call(1, 2))

	// X is array
	assert.Equal(t,
		[]interface{}{int64(3), int64(4), int64(5)},
		dummy.Call(
			[]interface{}{1, 2, 3},
			2,
		),
	)

	// Y is array
	assert.Equal(t,
		[]interface{}{int64(3), int64(4), int64(5)},
		dummy.Call(
			2,
			[]interface{}{1, 2, 3},
		),
	)

	// X, Y are array, same length
	assert.Equal(t,
		[]interface{}{int64(4), int64(6), int64(8)},
		dummy.Call(
			[]interface{}{1, 2, 3},
			[]interface{}{3, 4, 5},
		),
	)

	// X, Y are array, X smaller
	assert.Equal(t,
		[]interface{}{int64(4), int64(6), int64(8), int64(6), int64(7)},
		dummy.Call(
			[]interface{}{1, 2, 3},
			[]interface{}{3, 4, 5, 6, 7},
		),
	)

	// X, Y are array, Y smaller
	assert.Equal(t,
		[]interface{}{int64(4), int64(6), int64(8), int64(6), int64(7)},
		dummy.Call(
			[]interface{}{3, 4, 5, 6, 7},
			[]interface{}{1, 2, 3},
		),
	)
}

func TestUnaryOperatorFunc(t *testing.T) {
	var dummy unaryOperatorFunc = func(x interface{}) interface{} {
		cx, _ := cast.AsInt(x)
		return cx + 10
	}

	assert.Equal(t, int64(11), dummy(1))

	// scalar
	assert.Equal(t, int64(11), dummy.Call(1))

	// X is array
	assert.Equal(t,
		[]interface{}{int64(11), int64(12), int64(13)},
		dummy.Call([]interface{}{1, 2, 3}),
	)
}
