package serie

import (
	"fmt"
	"sort"

	"github.com/datasweet/datatable/value"
	"github.com/pkg/errors"
)

// Serie describes a serie
type Serie interface {
	Type() value.Type
	Value(at int) value.Value
	Values() []value.Value

	// Mutate
	Append(v ...interface{})
	Prepend(v ...interface{})
	Insert(at int, v ...interface{}) error
	Update(at int, v interface{}) error
	Delete(at int) error
	Grow(size int) error
	Shrink(size int) error

	// Select
	Head(size int) Serie
	Tail(size int) Serie
	Subset(at, size int) Serie
	Pick(at ...int) Serie
	FindRows(where Predicate) []int
	Filter(where Predicate) Serie
	Distinct() Serie

	// Clone
	Clone(includeValues bool) Serie

	// Sort
	sort.Interface
	SortAsc()
	SortDesc()

	// Print
	Print(opts ...PrintOption) string
	fmt.Stringer
}

// Predicate describes a predicate to filter a serie
type Predicate func(val value.Value) bool

// New to create a new serie
func New(builder value.Builder, v ...interface{}) (Serie, error) {
	if builder == nil {
		return nil, errors.New("value builder is nil")
	}
	s := &serie{
		builder: builder,
		typ:     builder().Type(),
	}
	if len(v) > 0 {
		s.Append(v...)
	}
	return s, nil
}

// Serie implementation
type serie struct {
	typ     value.Type
	builder value.Builder
	values  []value.Value
}

func (s *serie) Type() value.Type {
	return s.typ
}

func (s *serie) Value(at int) value.Value {
	if at < 0 || at >= len(s.values) {
		return s.builder(nil)
	}
	return s.values[at]
}

func (s *serie) Values() []value.Value {
	return s.values
}
