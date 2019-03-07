package expr_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/datasweet/datatable/expr"
)

type evalTest struct {
	input    string
	env      interface{}
	expected interface{}
}

type evalErrorTest struct {
	input string
	env   interface{}
	err   string
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
		"len(foo) == 3",
		map[string]interface{}{"foo": []rune{'a', 'b', 'c'}},
		true,
	},
	{
		`len(foo) == 6`,
		map[string]string{"foo": "foobar"},
		true,
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
		12,
	},
	{
		`60 ^ 13`,
		nil,
		49,
	},
	{
		`60 | 13`,
		nil,
		61,
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
		"`food` matches \"sea\" ~ \"food\"",
		map[string]interface{}{"food": []interface{}{"sea", "is", "more", "food"}},
		[]interface{}{true, false, false, true},
	},
}

var evalErrorTests = []evalErrorTest{
	{
		"bar",
		map[string]int{"foo": 1},
		`undefined: bar`,
	},
	{
		`"foo" ~ foo`,
		map[string]*int{"foo": nil},
		`interface conversion: interface {} is *int, not string`,
	},
	{
		"1 or 0",
		nil,
		"interface conversion: interface {} is float64, not bool",
	},
	{
		"not nil",
		nil,
		"interface conversion: interface {} is nil, not bool",
	},
	{
		"nil matches 'nil'",
		nil,
		"interface conversion: interface {} is nil, not string",
	},
	{
		"foo['bar'].baz",
		map[string]interface{}{"foo": nil},
		`cannot get "bar" from <nil>: foo["bar"]`,
	},
	{
		"foo.bar(abc)",
		map[string]interface{}{"foo": nil},
		`cannot get method bar from <nil>: foo.bar(abc)`,
	},
	{
		`"seafood" matches "a(b"`,
		nil,
		"error parsing regexp: missing closing ): `a(b`",
	},
	{
		`"seafood" matches "a" ~ ")b"`,
		nil,
		"error parsing regexp: unexpected ): `a)b`",
	},
	{
		`1 matches "1" ~ "2"`,
		nil,
		"interface conversion: interface {} is float64, not string",
	},
	{
		`1 matches "1"`,
		nil,
		"interface conversion: interface {} is float64, not string",
	},
	{
		`"1" matches 1`,
		nil,
		"interface conversion: interface {} is float64, not string",
	},
	{
		`foo ? 1 : 2`,
		map[string]interface{}{"foo": 0},
		`interface conversion: interface {} is int, not bool`,
	},
	{
		`foo()`,
		map[string]interface{}{"foo": func() (int, int) { return 0, 1 }},
		`func "foo" must return only one value`,
	},
	{
		`foo()`,
		map[string]interface{}{"foo": nil},
		`reflect: call of reflect.Value.Call on zero Value`,
	},
	{
		"1..1e6+1",
		nil,
		"range 1..1000001 exceeded max size of 1e6",
	},
	{
		"1/0",
		nil,
		"division by zero",
	},
	{
		"1%0",
		nil,
		"division by zero",
	},
	{
		"1 + 'a'",
		nil,
		`cannot convert "a" (type string) to type float64`,
	},
	{
		"'a' + 1",
		nil,
		`cannot convert "a" (type string) to type float64`,
	},
	{
		"[1, 2]['a']",
		nil,
		`cannot get "a" from []interface {}: [1, 2]["a"]`,
	},
	{
		`1 in "a"`,
		nil,
		`operator "in" not defined on string`,
	},
	{
		`nil in map`,
		map[string]interface{}{"map": map[string]interface{}{"true": "yes"}},
		`cannot use <nil> as index to map[string]interface {}`,
	},
	{
		`nil in foo`,
		map[string]interface{}{"foo": struct{ Bar bool }{true}},
		`cannot use <nil> as field name of struct { Bar bool }`,
	},
	{
		`true in foo`,
		map[string]interface{}{"foo": struct{ Bar bool }{true}},
		`cannot use bool as field name of struct { Bar bool }`,
	},
	{
		"len()",
		nil,
		"missing argument: len()",
	},
	{
		"len(1)",
		nil,
		"invalid argument len(1) (type float64)",
	},
	{
		"len(a, b)",
		nil,
		"too many arguments: len(a, b)",
	},
	{
		"Foo.Panic",
		struct{ Foo interface{} }{Foo: nil},
		"unrecognized character: U+002E '.'",
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

func TestEval_error(t *testing.T) {
	for _, test := range evalErrorTests {
		result, err := expr.Eval(test.input, test.env)
		if err == nil {
			err = fmt.Errorf("%v, <nil>", result)
		}
		if !strings.HasPrefix(err.Error(), test.err) || test.err == "" {
			t.Errorf("%s:\ngot\n\t%+v\nexpected\n\t%v", test.input, err.Error(), test.err)
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

func TestEval_method(t *testing.T) {
	env := &testEnv{
		Hello: "hello",
		World: testWorld{
			name: []string{"w", "o", "r", "l", "d"},
		},
		testVersion: &testVersion{
			version: 2,
		},
	}

	input := `Title(Hello) ~ Space() ~ (CompareVersion(1, 3) ? World.String() : '')`

	node, err := expr.Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := expr.Run(node, env)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello world"
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestEval_method:\ngot\n\t%#v\nexpected:\n\t%#v", actual, expected)
	}
}

type testVersion struct {
	version float64
}

func (c *testVersion) CompareVersion(min, max float64) bool {
	return min < c.version && c.version < max
}

type testWorld struct {
	name []string
}

func (w testWorld) String() string {
	return strings.Join(w.name, "")
}

type testEnv struct {
	*testVersion
	Hello string
	World testWorld
}

func (e *testEnv) Title(s string) string {
	return strings.Title(s)
}

func (e *testEnv) Space() string {
	return " "
}
