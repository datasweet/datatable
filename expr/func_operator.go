package expr

import (
	"reflect"

	"github.com/datasweet/datatable/cast"
)

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
	return float64(cx % cy)
}
