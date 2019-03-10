package expr

import (
	"fmt"
	"reflect"
	"testing"
)

type printTest struct {
	input    Node
	expected string
}

var printTests = []printTest{
	{
		builtinNode{"sum", []Node{nameNode{"top 5 - speaker"}}},
		"sum(`top 5 - speaker`)",
	},
	{
		builtinNode{"length", []Node{textNode{"foo"}}},
		"length(\"foo\")",
	},
	{
		functionNode{"call", []Node{arrayNode{[]Node{numberNode{1}, unaryNode{"not", boolNode{true}}}}}},
		"call([1, not true])",
	},
	{
		binaryNode{"and", binaryNode{"or", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"(a or b) and c",
	},
	{
		conditionalNode{nameNode{"a"}, nameNode{"a"}, nameNode{"b"}},
		"a ? a : b",
	},
	{
		matchesNode{left: nameNode{"foo"}, right: textNode{"foobar"}},
		"(foo matches \"foobar\")",
	},
	{
		binaryNode{"or", binaryNode{"or", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"a or b or c",
	},
	{
		binaryNode{"and", binaryNode{"or", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"(a or b) and c",
	},
	{
		binaryNode{"or", binaryNode{"and", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"a and b or c",
	},
	{
		binaryNode{"and", nameNode{"a"}, binaryNode{"or", nameNode{"b"}, nameNode{"c"}}},
		"a and (b or c)",
	},
	{
		binaryNode{"*", nameNode{"a"}, binaryNode{"+", nameNode{"b"}, nameNode{"c"}}},
		"a * (b + c)",
	},
	{
		binaryNode{"*", binaryNode{"+", nameNode{"a"}, nameNode{"b"}}, binaryNode{"+", nameNode{"c"}, nameNode{"d"}}},
		"(a + b) * (c + d)",
	},
	{
		binaryNode{"+", binaryNode{"+", binaryNode{"+", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}}, nameNode{"d"}},
		"a + b + c + d",
	},
	{
		binaryNode{"**", binaryNode{"**", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"(a ** b) ** c",
	},
	{
		unaryNode{"-", unaryNode{"+", unaryNode{"-", nameNode{"b"}}}},
		"(-(+(-b)))",
	},
	{
		binaryNode{"or", binaryNode{"and", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
		"a and b or c",
	},
	{
		binaryNode{"or", nameNode{"a"}, binaryNode{"and", nameNode{"b"}, nameNode{"c"}}},
		"a or b and c",
	},
}

func TestPrint(t *testing.T) {
	for _, test := range printTests {
		actual := fmt.Sprintf("%v", test.input)
		if actual != test.expected {
			t.Errorf("%s:\ngot\n\t%#v\nexpected\n\t%#v", test.expected, actual, test.expected)
		}
		// Parse again and check if ast same as before.
		ast, err := Parse(actual)
		if err != nil {
			t.Errorf("%s: can't parse printed expression", actual)
		}
		if m, ok := ast.(matchesNode); ok {
			m.r = nil
			ast = m
		}
		if !reflect.DeepEqual(ast, test.input) {
			t.Errorf("%s:\ngot\n\t%#v\nexpected\n\t%#v", test.expected, ast, test.input)
		}
	}
}
