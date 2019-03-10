package expr_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/datasweet/datatable/expr"
)

type evalTest struct {
	input    string
	env      interface{}
	expected interface{}
}

type evalParams map[string]interface{}

func (p evalParams) Max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func (p evalParams) Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

var evalTests = []evalTest{
	{
		"foo",
		map[string]int{"foo": 33},
		33,
	},
	{
		"foo == bar",
		map[string]interface{}{"foo": 1, "bar": 1},
		true,
	},
	{
		"foo || (bar && !false && true)",
		map[string]interface{}{"foo": false, "bar": true},
		true,
	},
	{
		"foo && bar",
		map[string]interface{}{"foo": false, "bar": true},
		false,
	},
	{
		"!foo && bar",
		map[string]interface{}{"foo": false, "bar": true},
		true,
	},
	{
		"true || false",
		nil,
		true,
	},
	{
		"false && true",
		nil,
		false,
	},
	{
		"2+2==4",
		nil,
		true,
	},
	{
		"2+3",
		nil,
		float64(5),
	},
	{
		"5-2",
		nil,
		float64(3),
	},
	{
		"2*3",
		nil,
		float64(6),
	},
	{
		"6/2",
		nil,
		float64(3),
	},
	{
		"8%3",
		nil,
		float64(2),
	},
	{
		"2**4",
		nil,
		float64(16),
	},
	{
		"2**4",
		nil,
		float64(16),
	},
	{
		"-(2-5)**3-2/(+4-3)+-2",
		nil,
		float64(23),
	},
	{
		`"hello" ~ ' ' ~ "world"`,
		nil,
		"hello world",
	},
	{
		"+0 == -0",
		nil,
		true,
	},
	{
		"1 < 2 and 3 > 2",
		nil,
		true,
	},
	{
		"!(1 != 1) && 2 >= 2 && 3 <= 3",
		nil,
		true,
	},
	{
		"[1, 02, 1e3, 1.2e-4]",
		nil,
		[]interface{}{float64(1), float64(2), float64(1000), float64(0.00012)},
	},
	{
		`"1" in ["1", "2"]`,
		nil,
		true,
	},
	{
		`"0" not in ["1", "2"]`,
		nil,
		true,
	},
	{
		`0 in nil`,
		nil,
		false,
	},
	{
		`60 & 13`,
		nil,
		int64(12),
	},
	{
		`60 ^ 13`,
		nil,
		int64(49),
	},
	{
		`60 | 13`,
		nil,
		int64(61),
	},
	{
		`"seafood" matches "foo.*"`,
		nil,
		true,
	},
	{
		`"seafood" matches "sea" ~ "food"`,
		nil,
		true,
	},
	{
		`not ("seafood" matches "[0-9]+") ? "a" : "b"`,
		nil,
		"a",
	},
	{
		`false ? "a" : "b"`,
		nil,
		"b",
	},
	{
		`foo("world")`,
		map[string]interface{}{"foo": func(in string) string { return "hello " + in }},
		"hello world",
	},
	{
		"Max(a, b)",
		evalParams{"a": 1.23, "b": 3.21},
		3.21,
	},
	{
		"Min(a, b)",
		evalParams{"a": 1.23, "b": 3.21},
		1.23,
	},
	{
		"`sessions` + `bounces`",
		map[string]interface{}{"sessions": []interface{}{10, 20, 30}, "bounces": []interface{}{5, 10, 15}},
		[]interface{}{float64(15), float64(30), float64(45)},
	},
	{
		`cond > 0 ? "yes" : "no"`,
		map[string]interface{}{"cond": 10, "country": []interface{}{"fr", "de", "es", "gb"}},
		"yes",
	},
	{
		"`cond` > 0 ? `country` : 'no'",
		map[string]interface{}{"cond": []interface{}{10, -20, 30}, "country": []interface{}{"fr", "de", "es", "gb"}},
		[]interface{}{"fr", "no", "es"},
	},
	{
		"`cond` > 0 ? ['fr', 'de', 'es', 'gb'] : 'no'",
		map[string]interface{}{"cond": []interface{}{10, -20, 30}},
		[]interface{}{"fr", "no", "es"},
	},
	{
		"`food` matches \"sea\" ~ \"|\" ~ \"food\"",
		map[string]interface{}{"food": []interface{}{"sea", "is", "more", "food"}},
		[]interface{}{true, false, false, true},
	},
}

func TestEval(t *testing.T) {
	for _, test := range evalTests {
		actual, err := expr.Eval(test.input, test.env)
		if err != nil {
			t.Errorf("%s:\n%v", test.input, err)
			continue
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s:\ngot\n\t%#v\nexpected\n\t%#v", test.input, actual, test.expected)
		}
	}
}

func TestEval_panic(t *testing.T) {
	node, err := expr.Parse("foo()")
	if err != nil {
		t.Fatal(err)
	}

	_, err = expr.Run(node, map[string]interface{}{"foo": nil})
	if err == nil {
		err = fmt.Errorf("<nil>")
	}

	expected := "reflect: call of reflect.Value.Call on zero Value"
	if err.Error() != expected {
		t.Errorf("\ngot\n\t%+v\nexpected\n\t%v", err.Error(), expected)
	}
}

func TestEval_func(t *testing.T) {
	type testEnv struct {
		Func func() string
	}

	env := &testEnv{
		Func: func() string {
			return "func"
		},
	}

	input := `Func()`

	node, err := expr.Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := expr.Run(node, env)
	if err != nil {
		t.Fatal(err)
	}

	expected := "func"
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestEval_method:\ngot\n\t%#v\nexpected:\n\t%#v", actual, expected)
	}
}
