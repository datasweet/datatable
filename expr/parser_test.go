package expr

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type parseTest struct {
	input    string
	expected Node
}

type parseErrorTest struct {
	input string
	err   string
}

var parseTests = []parseTest{
	{
		"a",
		nameNode{"a"},
	},
	{
		`"a"`,
		textNode{"a"},
	},
	{
		"3",
		numberNode{3},
	},
	{
		"true",
		boolNode{true},
	},
	{
		"false",
		boolNode{false},
	},
	{
		"nil",
		nilNode{},
	},
	{
		"-3",
		unaryNode{"-", numberNode{3}},
	},
	{
		"1 - 2",
		binaryNode{"-", numberNode{1}, numberNode{2}},
	},
	{
		"(1 - 2) * 3",
		binaryNode{"*", binaryNode{"-", numberNode{1}, numberNode{2}}, numberNode{3}},
	},
	{
		"a or b or c",
		binaryNode{"or", binaryNode{"or", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
	},
	{
		"a or b and c",
		binaryNode{"or", nameNode{"a"}, binaryNode{"and", nameNode{"b"}, nameNode{"c"}}},
	},
	{
		"(a or b) and c",
		binaryNode{"and", binaryNode{"or", nameNode{"a"}, nameNode{"b"}}, nameNode{"c"}},
	},
	{
		"2**4-1",
		binaryNode{"-", binaryNode{"**", numberNode{2}, numberNode{4}}, numberNode{1}},
	},
	{
		"foo(bar())",
		functionNode{"foo", []Node{functionNode{"bar", []Node{}}}},
	},
	{
		"true ? true : false",
		conditionalNode{boolNode{true}, boolNode{true}, boolNode{false}},
	},
	{
		"a ?: b",
		conditionalNode{nameNode{"a"}, nameNode{"a"}, nameNode{"b"}},
	},
	{
		"+0 != -0",
		binaryNode{"!=", unaryNode{"+", numberNode{0}}, unaryNode{"-", numberNode{0}}},
	},
	{
		"[a, b, c]",
		arrayNode{[]Node{nameNode{"a"}, nameNode{"b"}, nameNode{"c"}}},
	},
	{
		"length(\"foo\")",
		builtinNode{"length", []Node{textNode{"foo"}}},
	},
	{
		`foo matches "foo"`,
		matchesNode{left: nameNode{"foo"}, right: textNode{"foo"}},
	},
	{
		`foo matches regex`,
		matchesNode{left: nameNode{"foo"}, right: nameNode{"regex"}},
	},
	{
		"UPPER(`foo`)",
		builtinNode{"upper", []Node{nameNode{"foo"}}},
	},
}

var parseErrorTests = []parseErrorTest{
	{
		"foo(",
		`unclosed "("`,
	},
	{
		"a+",
		"unexpected token EOF",
	},
	{
		"a ? (1+2) c",
		"unexpected token name(c)",
	},
	{
		"[a b]",
		"array items must be separated by a comma",
	},
	{
		"bar(a b)",
		"arguments must be separated by a comma",
	},
	{
		"a matches 'a)(b'",
		"error parsing regexp: unexpected )",
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		actual, err := Parse(test.input)
		if err != nil {
			t.Errorf("%s:\n%v", test.input, err)
			continue
		}
		if m, ok := actual.(matchesNode); ok {
			m.r = nil
			actual = m
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s:\ngot\n\t%#v\nexpected\n\t%#v", test.input, actual, test.expected)
		}
	}
}

func TestParse_error(t *testing.T) {
	for _, test := range parseErrorTests {
		_, err := Parse(test.input)
		if err == nil {
			err = fmt.Errorf("<nil>")
		}
		if !strings.HasPrefix(err.Error(), test.err) || test.err == "" {
			t.Errorf("%s:\ngot\n\t%+v\nexpected\n\t%v", test.input, err.Error(), test.err)
		}
	}
}
