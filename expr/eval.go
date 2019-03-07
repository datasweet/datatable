package expr

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/datasweet/datatable/cast"
)

// Eval parses and evaluates given input.
func Eval(input string, env interface{}) (interface{}, error) {
	node, err := Parse(input)
	if err != nil {
		return nil, err
	}
	return Run(node, env)
}

// Run evaluates given ast.
func Run(node Node, env interface{}) (out interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	return node.Eval(env)
}

// eval functions

func (n nilNode) Eval(env interface{}) (interface{}, error) {
	return nil, nil
}

func (n identifierNode) Eval(env interface{}) (interface{}, error) {
	return n.value, nil
}

func (n numberNode) Eval(env interface{}) (interface{}, error) {
	return n.value, nil
}

func (n boolNode) Eval(env interface{}) (interface{}, error) {
	return n.value, nil
}

func (n textNode) Eval(env interface{}) (interface{}, error) {
	return n.value, nil
}

func (n nameNode) Eval(env interface{}) (interface{}, error) {
	v, ok := extract(env, n.name)
	if !ok {
		return nil, fmt.Errorf("undefined: %v", n)
	}
	return v, nil
}

func (n unaryNode) Eval(env interface{}) (interface{}, error) {
	val, err := n.node.Eval(env)
	if err != nil {
		return nil, err
	}

	switch n.operator {
	case "not", "!":
		return not.Call(val), nil
	case "-":
		return minus.Call(val), nil
	case "+":
		return plus.Call(val), nil
	}

	return nil, fmt.Errorf("implement unary %q operator", n.operator)
}

func (n binaryNode) Eval(env interface{}) (interface{}, error) {
	left, err := n.left.Eval(env)
	if err != nil {
		return nil, err
	}

	right, err := n.right.Eval(env)
	if err != nil {
		return nil, err
	}

	switch n.operator {
	case "or", "||":
		return logicalOR.Call(left, right), nil

	case "and", "&&":
		return logicalAND.Call(left, right), nil

	case "==":
		return equals.Call(left, right), nil

	case "!=":
		eq := equals.Call(left, right).(bool)
		return !eq, nil

	case "in":
		ok, err := contains(left, right)
		if err != nil {
			return nil, err
		}
		return ok, nil

	case "not in":
		ok, err := contains(left, right)
		if err != nil {
			return nil, err
		}
		return !ok, nil

	case "~":
		return concat.Call(left, right), nil

	case "|":
		return bitwiseOR.Call(left, right), nil

	case "^":
		return bitwiseXOR.Call(left, right), nil

	case "&":
		return bitwiseAND.Call(left, right), nil

	case "<":
		return lt.Call(left, right), nil

	case ">":
		return gt.Call(left, right), nil

	case ">=":
		return gte.Call(left, right), nil

	case "<=":
		return lte.Call(left, right), nil

	case "+":
		return add.Call(left, right), nil

	case "-":
		return substract.Call(left, right), nil

	case "*":
		return multiply.Call(left, right), nil

	case "/":
		if div, ok := cast.AsFloat(right); ok && div == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return divide.Call(left, right), nil

	case "%":
		if div, ok := cast.AsInt(right); ok && div == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return remainder.Call(left, right), nil

	case "**":
		return pow.Call(left, right), nil
	}

	return nil, fmt.Errorf("implement %q operator", n.operator)
}

func (n matchesNode) Eval(env interface{}) (interface{}, error) {
	left, err := n.left.Eval(env)
	if err != nil {
		return nil, err
	}

	if n.r != nil {
		return n.r.MatchString(left.(string)), nil
	}

	right, err := n.right.Eval(env)
	if err != nil {
		return nil, err
	}

	matched, err := regexp.MatchString(right.(string), left.(string))
	if err != nil {
		return nil, err
	}
	return matched, nil
}

func (n builtinNode) Eval(env interface{}) (interface{}, error) {
	switch n.name {
	case "len":
		if len(n.arguments) == 0 {
			return nil, fmt.Errorf("missing argument: %v", n)
		}
		if len(n.arguments) > 1 {
			return nil, fmt.Errorf("too many arguments: %v", n)
		}

		i, err := n.arguments[0].Eval(env)
		if err != nil {
			return nil, err
		}

		switch reflect.TypeOf(i).Kind() {
		case reflect.Array, reflect.Slice, reflect.String:
			return float64(reflect.ValueOf(i).Len()), nil
		}
		return nil, fmt.Errorf("invalid argument %v (type %T)", n, i)
	}

	return nil, fmt.Errorf("unknown %q builtin", n.name)
}

func (n functionNode) Eval(env interface{}) (interface{}, error) {
	fn, ok := getFunc(env, n.name)
	if !ok {
		return nil, fmt.Errorf("undefined: %v", n.name)
	}

	in := make([]reflect.Value, 0)

	for _, a := range n.arguments {
		i, err := a.Eval(env)
		if err != nil {
			return nil, err
		}
		in = append(in, reflect.ValueOf(i))
	}

	out := reflect.ValueOf(fn).Call(in)

	if len(out) == 0 {
		return nil, nil
	} else if len(out) > 1 {
		return nil, fmt.Errorf("func %q must return only one value", n.name)
	}

	if out[0].IsValid() && out[0].CanInterface() {
		return out[0].Interface(), nil
	}

	return nil, nil
}

func (n conditionalNode) Eval(env interface{}) (interface{}, error) {
	cond, err := n.cond.Eval(env)
	if err != nil {
		return nil, err
	}

	// Not optimized we evaluate both then and else
	yes, err := n.exp1.Eval(env)
	if err != nil {
		return nil, err
	}

	no, err := n.exp2.Eval(env)
	if err != nil {
		return nil, err
	}

	if arr, ok := asArray(cond); ok {
		arrYes, okYes := asArray(yes)
		arrNo, okNo := asArray(no)

		arrOut := make([]interface{}, len(arr))

		for i, c := range arr {
			if c.(bool) {
				if okYes {
					arrOut[i] = getAt(arrYes, i)
				} else {
					arrOut[i] = yes
				}
			} else {
				if okNo {
					arrOut[i] = getAt(arrNo, i)
				} else {
					arrOut[i] = no
				}
			}
		}

		return arrOut, nil
	} else if cond.(bool) {
		return yes, nil
	} else {
		return no, nil
	}
}

func (n arrayNode) Eval(env interface{}) (interface{}, error) {
	array := make([]interface{}, 0)
	for _, node := range n.nodes {
		val, err := node.Eval(env)
		if err != nil {
			return nil, err
		}
		array = append(array, val)
	}
	return array, nil
}
