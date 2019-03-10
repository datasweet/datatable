package expr

import (
	"math"

	"github.com/datasweet/datatable/cast"
)

// ABS(x float)
var abs unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Abs(cx)
	}
	return nil
}

// ACOS(x float)
var acos unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Acos(cx)
	}
	return nil
}

// ASIN(x float)
var asin unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Asin(cx)
	}
	return nil
}

// ATAN(x float)
var atan unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Atan(cx)
	}
	return nil
}

// CEIL(x float)
var ceil unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Ceil(cx)
	}
	return nil
}

// COS(x float)
var cos unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Cos(cx)
	}
	return nil
}

// FLOOR(x float)
var floor unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Floor(cx)
	}
	return nil
}

// LOG(x float)
var log unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Log(cx)
	}
	return nil
}

// LOG10(x float)
var log10 unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Log10(cx)
	}
	return nil
}

// POW(x float, y float)
// Used with operator **
var pow binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, okx := cast.AsFloat(x)
	cy, oky := cast.AsFloat(y)
	if !okx || !oky {
		return nil
	}
	return math.Pow(cx, cy)
}

// ROUND(x float)
var round unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Round(cx)
	}
	return nil
}

// SIN(x float)
var sin unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Sin(cx)
	}
	return nil
}

// TAN(x float)
var tan unaryOperatorFunc = func(x interface{}) interface{} {
	if cx, okx := cast.AsFloat(x); okx {
		return math.Tan(cx)
	}
	return nil
}
