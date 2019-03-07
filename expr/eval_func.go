package expr

import (
	"math"
	"reflect"

	"github.com/datasweet/datatable/cast"
)

type unaryOperatorFunc func(x interface{}) interface{}

func (fn unaryOperatorFunc) Call(a interface{}) interface{} {
	if arr, ok := asArray(a); ok {
		cnt := len(arr)
		out := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			out[i] = fn(arr[i])
		}
		return out
	}

	return fn(a)
}

type binaryOperatorFunc func(x, y interface{}) interface{}

func (fn binaryOperatorFunc) Call(a, b interface{}) interface{} {
	arrA, okA := asArray(a)
	arrB, okB := asArray(b)

	if okA && okB {
		lenA, lenB := len(arrA), len(arrB)
		cnt := lenA
		if lenB > lenA {
			cnt = lenB
		}
		arrC := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			arrC[i] = fn(getAt(arrA, i), getAt(arrB, i))
		}
		return arrC
	}

	if okA {
		cnt := len(arrA)
		arrC := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			arrC[i] = fn(getAt(arrA, i), b)
		}
		return arrC
	}

	if okB {
		cnt := len(arrB)
		arrC := make([]interface{}, cnt)
		for i := 0; i < cnt; i++ {
			arrC[i] = fn(a, getAt(arrB, i))
		}
		return arrC
	}

	return fn(a, b)
}

func asArray(v interface{}) ([]interface{}, bool) {
	if casted, ok := v.([]interface{}); ok {
		return casted, true
	}
	return nil, false
}

func getAt(arr []interface{}, at int) interface{} {
	if at < 0 || at >= len(arr) {
		return nil
	}
	return arr[at]
}

// operator not
// x is bool
var not unaryOperatorFunc = func(x interface{}) interface{} {
	nx, _ := cast.AsBool(x)
	return !nx
}

// operator plus
// x is float
var plus unaryOperatorFunc = func(x interface{}) interface{} {
	if nx, okx := cast.AsFloat(x); okx {
		return +nx
	}
	return nil
}

// operator minus
// x is float
var minus unaryOperatorFunc = func(x interface{}) interface{} {
	if nx, okx := cast.AsFloat(x); okx {
		return -nx
	}
	return nil
}

// operator ||
// x,y are bool
var logicalOR binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsBool(x)
	cy, oky := cast.AsBool(y)
	if !okx || !oky {
		return nil
	}
	return cx || cy
}

// operator &&
// x,y are bool
var logicalAND binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsBool(x)
	cy, oky := cast.AsBool(y)
	if !okx || !oky {
		return nil
	}
	return cx && cy
}

// operator ==
// x,y are interface
// Based on testify/assert.Equals
var equals binaryOperatorFunc = func(x, y interface{}) interface{} {
	if x == nil || y == nil {
		return x == y
	}
	return reflect.DeepEqual(x, y)
}

// operator ~
// x,y are string
var concat binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsString(x)
	cy, oky := cast.AsString(y)
	if !okx || !oky {
		return nil
	}
	return cx + cy
}

// operator |
// x,y are int
var bitwiseOR binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsInt(x)
	cy, oky := cast.AsInt(y)
	if !okx || !oky {
		return nil
	}
	return cx | cy
}

// operator ^
// x,y are int
var bitwiseXOR binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsInt(x)
	cy, oky := cast.AsInt(y)
	if !okx || !oky {
		return nil
	}
	return cx ^ cy
}

// operator &
// x,y are int
var bitwiseAND binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsInt(x)
	cy, oky := cast.AsInt(y)
	if !okx || !oky {
		return nil
	}
	return cx & cy
}

// operator <
// x,y are float
var lt binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx < cy
}

// operator <=
// x,y are float
var lte binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx <= cy
}

// operator >
// x,y are float
var gt binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx > cy
}

// operator >=
// x,y are float
var gte binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx >= cy
}

// operator +
// x,y are float
var add binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx + cy
}

// operator -
// x,y are float
var substract binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx - cy
}

// operator *
// x,y are float
var multiply binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx * cy
}

// operator /
// x,y are float
var divide binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return cx / cy
}

// operator %
// x,y are int
var remainder binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsInt(x)
	cy, oky := cast.AsInt(y)
	if !okx || !oky {
		return nil
	}
	return cx % cy
}

// operator **
// x, y are float
var pow binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return math.Pow(cx, cy)
}

// math
