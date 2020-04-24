package datatable

// import (
// 	"fmt"
// 	"sort"

// 	"github.com/datasweet/datatable/value"
// )

// // Serie describes a serie
// // a serie can contains scalar or table
// type Serie interface {
// 	Type() value.Type
// 	Value(at int) value.Value
// 	Values() []value.Value

// 	// Mutate
// 	Append(v ...interface{})
// 	Prepend(v ...interface{})
// 	Insert(at int, v ...interface{}) error
// 	Update(at int, v interface{}) error
// 	Delete(at int) error
// 	Grow(size int) error
// 	Shrink(size int) error

// 	// Select
// 	Head(size int) Serie
// 	Tail(size int) Serie
// 	Subset(at, size int) Serie
// 	Pick(at ...int) Serie
// 	FindRows(where Predicate) []int
// 	Filter(where Predicate) Serie
// 	Distinct() Serie

// 	// Copy
// 	Copy(mode CopyMode) Serie

// 	// Sort
// 	sort.Interface
// 	SortAsc()
// 	SortDesc()

// 	// Print
// 	Print(opts ...PrintOption) string
// 	fmt.Stringer
// }
