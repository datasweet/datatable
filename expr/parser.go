package expr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type associativity int

const (
	left associativity = iota + 1
	right
)

type info struct {
	precedence    int
	associativity associativity
}

var unaryOperators = map[string]info{
	"not": {50, left},
	"!":   {50, left},
	"-":   {500, left},
	"+":   {500, left},
}

var binaryOperators = map[string]info{
	"or":      {10, left},
	"||":      {10, left},
	"and":     {15, left},
	"&&":      {15, left},
	"|":       {16, left},
	"^":       {17, left},
	"&":       {18, left},
	"==":      {20, left},
	"!=":      {20, left},
	"<":       {20, left},
	">":       {20, left},
	">=":      {20, left},
	"<=":      {20, left},
	"not in":  {20, left},
	"in":      {20, left},
	"matches": {20, left},
	"..":      {25, left},
	"+":       {30, left},
	"-":       {30, left},
	"~":       {40, left},
	"*":       {60, left},
	"/":       {60, left},
	"%":       {60, left},
	"**":      {200, right},
}

var builtins = map[string]bool{
	// AGG
	"avg":            true,
	"count":          true,
	"count_distinct": true,
	"cusum":          true,
	"max":            true,
	"median":         true,
	"min":            true,
	"percentile":     true,
	"stddev":         true,
	"sum":            true,
	"variance":       true,

	// DATE
	"date_diff": true,
	"day":       true,
	"hour":      true,
	"minute":    true,
	"month":     true,
	"quarter":   true,
	"second":    true,
	"week":      true,
	"weekday":   true,
	"year":      true,

	// MATH
	"abs":   true,
	"acos":  true,
	"asin":  true,
	"atan":  true,
	"ceil":  true,
	"cos":   true,
	"floor": true,
	"log":   true,
	"log10": true,
	"pow":   true,
	"round": true,
	"sin":   true,
	"tan":   true,

	// TEXT
	"concat": true,
	"length": true,
	"lower":  true,
	"trim":   true,
	"upper":  true,
}

type parser struct {
	input    string
	tokens   []token
	position int
	current  token
	strict   bool
}

// OptionFn for configuring parser.
type OptionFn func(p *parser)

// Parse parses input into ast.
func Parse(input string, ops ...OptionFn) (Node, error) {
	tokens, err := lex(input)
	if err != nil {
		return nil, err
	}

	p := &parser{
		input:   input,
		tokens:  tokens,
		current: tokens[0],
	}

	for _, op := range ops {
		op(p)
	}

	node, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}

	if !p.isEOF() {
		return nil, p.errorf("unexpected token %v", p.current)
	}

	return node, nil
}

func (p *parser) errorf(format string, args ...interface{}) *syntaxError {
	return &syntaxError{
		message: fmt.Sprintf(format, args...),
		input:   p.input,
		pos:     p.current.pos,
	}
}

func (p *parser) next() error {
	p.position++
	if p.position >= len(p.tokens) {
		return p.errorf("unexpected end of expression")
	}
	p.current = p.tokens[p.position]
	return nil
}

func (p *parser) expect(kind tokenKind, values ...string) error {
	if p.current.is(kind, values...) {
		return p.next()
	}
	return p.errorf("unexpected token %v", p.current)
}

func (p *parser) isEOF() bool {
	return p.current.is(eof)
}

// parse functions

func (p *parser) parseExpression(precedence int) (Node, error) {
	node, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}
	token := p.current
	for token.is(operator) {
		if op, ok := binaryOperators[token.value]; ok {
			if op.precedence >= precedence {
				if err = p.next(); err != nil {
					return nil, err
				}

				var expr Node
				if op.associativity == left {
					expr, err = p.parseExpression(op.precedence + 1)
					if err != nil {
						return nil, err
					}
				} else {
					expr, err = p.parseExpression(op.precedence)
					if err != nil {
						return nil, err
					}
				}

				if token.is(operator, "matches") {
					var r *regexp.Regexp
					if s, ok := expr.(textNode); ok {
						r, err = regexp.Compile(s.value)
						if err != nil {
							return nil, p.errorf("%v", err)
						}
					}
					node = matchesNode{r: r, left: node, right: expr}
				} else {
					node = binaryNode{operator: token.value, left: node, right: expr}
				}
				token = p.current
				continue
			}
		}
		break
	}

	if precedence == 0 {
		node, err = p.parseConditionalExpression(node)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

func (p *parser) parsePrimary() (Node, error) {
	token := p.current

	if token.is(operator) {
		if op, ok := unaryOperators[token.value]; ok {
			if err := p.next(); err != nil {
				return nil, err
			}
			expr, err := p.parseExpression(op.precedence)
			if err != nil {
				return nil, err
			}

			return unaryNode{operator: token.value, node: expr}, nil
		}
	}

	if token.is(punctuation, "(") {
		if err := p.next(); err != nil {
			return nil, err
		}

		expr, err := p.parseExpression(0)
		if err != nil {
			return nil, err
		}

		err = p.expect(punctuation, ")")
		if err != nil {
			return nil, p.errorf("an opened parenthesis is not properly closed")
		}

		return expr, nil
	}

	return p.parsePrimaryExpression()
}

func (p *parser) parseConditionalExpression(node Node) (Node, error) {
	var err error
	var expr1, expr2 Node
	for p.current.is(punctuation, "?") {
		if err := p.next(); err != nil {
			return nil, err
		}
		if !p.current.is(punctuation, ":") {
			expr1, err = p.parseExpression(0)
			if err != nil {
				return nil, err
			}
			if err := p.expect(punctuation, ":"); err != nil {
				return nil, err
			}
			expr2, err = p.parseExpression(0)
			if err != nil {
				return nil, err
			}
		} else {
			if err := p.next(); err != nil {
				return nil, err
			}
			expr1 = node
			expr2, err = p.parseExpression(0)
			if err != nil {
				return nil, err
			}
		}

		node = conditionalNode{node, expr1, expr2}
	}
	return node, nil
}

func (p *parser) parsePrimaryExpression() (Node, error) {
	var err error
	var node Node
	token := p.current
	switch token.kind {
	case name:
		if err := p.next(); err != nil {
			return nil, err
		}
		switch token.value {
		case "true":
			return boolNode{value: true}, nil
		case "false":
			return boolNode{value: false}, nil
		case "nil":
			return nilNode{}, nil
		default:
			node, err = p.parseNameExpression(token)
			if err != nil {
				return nil, err
			}
		}

	case number:
		if err := p.next(); err != nil {
			return nil, err
		}
		number, err := strconv.ParseFloat(token.value, 64)
		if err != nil {
			return nil, p.errorf("%v", err)
		}
		return numberNode{value: number}, nil

	case text:
		if err := p.next(); err != nil {
			return nil, err
		}
		return textNode{value: token.value}, nil

	default:
		if token.is(punctuation, "[") {
			node, err = p.parseArrayExpression()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, p.errorf("unexpected token %v", token).at(token)
		}
	}

	return node, nil
}

func (p *parser) parseNameExpression(token token) (Node, error) {
	var node Node
	if p.current.is(punctuation, "(") {
		arguments, err := p.parseArguments()
		if err != nil {
			return nil, err
		}
		tv := strings.ToLower(token.value)
		if _, ok := builtins[tv]; ok {
			node = builtinNode{name: tv, arguments: arguments}
		} else {
			node = functionNode{name: token.value, arguments: arguments}
		}
	} else {
		node = nameNode{name: token.value}
	}
	return node, nil
}

func (p *parser) parseArrayExpression() (Node, error) {
	nodes, err := p.parseList("array items", "[", "]")
	if err != nil {
		return nil, err
	}
	return arrayNode{nodes}, nil
}

func isValidIdentifier(str string) bool {
	if len(str) == 0 {
		return false
	}
	h, w := utf8.DecodeRuneInString(str)
	if !isAlphabetic(h) {
		return false
	}
	for _, r := range str[w:] {
		if !isAlphaNumeric(r) {
			return false
		}
	}
	return true
}

func (p *parser) parseArguments() ([]Node, error) {
	return p.parseList("arguments", "(", ")")
}

func (p *parser) parseList(what, start, end string) ([]Node, error) {
	err := p.expect(punctuation, start)
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for !p.current.is(punctuation, end) {
		if len(nodes) > 0 {
			err = p.expect(punctuation, ",")
			if err != nil {
				return nil, p.errorf("%v must be separated by a comma", what)
			}
		}
		node, err := p.parseExpression(0)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	err = p.expect(punctuation, end)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}
