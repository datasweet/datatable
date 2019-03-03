package formula

import "github.com/Knetic/govaluate"

var All map[string]govaluate.ExpressionFunction

func init() {
	All = make(map[string]govaluate.ExpressionFunction)
}
