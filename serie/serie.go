package serie

import (
	"fmt"
	"sort"

	"github.com/datasweet/datatable/value"
	"github.com/pkg/errors"
)

// Serie describe a serie
type Serie interface {
	Type() value.Type
	Error() error
	Value(at int) value.Value
	Values() []value.Value

	// Mutate
	Append(v ...interface{})
	Prepend(v ...interface{})
	Insert(at int, v ...interface{})
	Update(at int, v interface{})
	Delete(at int)
	Grow(size int)
	Shrink(size int)

	// Select
	Head(size int) Serie
	Tail(size int) Serie
	Subset(at, size int) Serie
	Pick(at ...int) Serie
	FindRows(where Predicate) []int
	Filter(where Predicate) Serie
	Distinct() Serie

	// Clone
	Clone() Serie

	// Sort
	sort.Interface
	SortAsc() Serie
	SortDesc() Serie

	// Print
	Print(opts ...PrintOption) string
	fmt.Stringer
}

type Predicate func(val value.Value) bool

// New to create a new serie
func New(builder value.Builder, v ...interface{}) Serie {
	s := &serie{
		builder: builder,
	}
	if builder == nil {
		s.typ = value.Type("?")
		s.err = errors.New("value builder is nil")
	} else {
		s.typ = builder().Type()
	}

	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

// Serie implementation
type serie struct {
	typ     value.Type
	builder value.Builder
	values  []value.Value
	err     error
}

func (s *serie) Type() value.Type {
	return s.typ
}

func (s *serie) Error() error {
	return s.err
}

func (s *serie) Value(at int) value.Value {
	if s.err != nil {
		return nil
	}
	if at < 0 || at >= len(s.values) {
		return nil
	}
	return s.values[at]
}

func (s *serie) Values() []value.Value {
	return s.values
}
