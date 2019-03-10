package expr

import (
	"regexp"
	"strings"

	"github.com/datasweet/datatable/cast"
)

// CONCAT(x, y)
// Used with operator ~
var concat binaryOperatorFunc = func(x, y interface{}) interface{} {
	cx, _ := cast.AsString(x)
	cy, _ := cast.AsString(y)
	return cx + cy
}

// LENGTH(x string)
// Returns the number of characters in X
var length unaryOperatorFunc = func(x interface{}) interface{} {
	cx, _ := cast.AsString(x)
	return len(cx)
}

// LOWER(x string)
// Converts X to lowercase
var lower unaryOperatorFunc = func(x interface{}) interface{} {
	cx, _ := cast.AsString(x)
	return strings.ToLower(cx)
}

// MATCHES(x string, y string or regexp)
// Returns true if X matches Y, false otherwise
var matches binaryOperatorFunc = func(x interface{}, y interface{}) interface{} {
	cx, _ := cast.AsString(x)

	if rg, ok := y.(*regexp.Regexp); ok && rg != nil {
		return rg.MatchString(cx)
	}

	cy, _ := cast.AsString(y)
	if matched, err := regexp.MatchString(cy, cx); err == nil {
		return matched
	}

	return false
}

// TRIM(x string)
// Returns X with leading and trailing whitespace removed
var trim unaryOperatorFunc = func(x interface{}) interface{} {
	cx, _ := cast.AsString(x)
	return strings.TrimSpace(cx)
}

// UPPER(x string)
// Converts X to uppercase
var upper unaryOperatorFunc = func(x interface{}) interface{} {
	cx, _ := cast.AsString(x)
	return strings.ToUpper(cx)
}
