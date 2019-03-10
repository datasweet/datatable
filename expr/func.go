package expr

type unaryOperatorFunc func(x interface{}) interface{}

type binaryOperatorFunc func(x, y interface{}) interface{}

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

func flatten(args ...interface{}) []interface{} {
	var values []interface{}
	for _, a := range args {
		if arr, ok := asArray(a); ok {
			values = append(values, flatten(arr...)...)
		} else {
			values = append(values, a)
		}
	}
	return values
}
