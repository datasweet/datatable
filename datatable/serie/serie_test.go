package serie

import (
	"fmt"
	"testing"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func AsInt8(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	if n, ok := cast.AsInt(v); ok {
		// check overflow
		casted := int8(n)
		if int64(casted) == n {
			return casted
		}
	}
	return nil
}

type Value interface {
	Set(v interface{})
	Val() interface{}
	IsValid() bool
}

type Int32 struct {
	val int32
}

func TestNewSerie(t *testing.T) {
	s := newSerie(int8(0), AsInt8)
	assert.NotNil(t, s)
	assert.NoError(t, s.Error())
	fmt.Println(s.values)
	s.Append(1, 2, nil, "2", false, -1)
	fmt.Println(s.values)
}
