package datatable

import (
	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
)

type ColumnType uint8

const (
	Raw ColumnType = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Datetime
	Duration
)

// ColumnOptions are options when a columns is added to a datatable
type ColumnOptions struct {
	Name string
	Type ColumnType
	Len  int
}

type ColumnOption func(opts *ColumnOptions)

type Column interface {
	Name() string
	Type() string
	IsVisible() bool
	IsComputed() bool
}

// column is a column in our datatable
// A column contains all rows
type column struct {
	name     string
	formulae string
	expr     expr.Node
	hidden   bool
	serie    serie.Serie
}

// newColumn to create a column
func newColumn(name string, serie serie.Serie) *column {
	col := &column{
		name:  name,
		serie: serie,
	}

	return col
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Type() string {
	return string(c.serie.Type())
}

func (c *column) IsVisible() bool {
	return !c.hidden
}

func (c *column) IsComputed() bool {
	return c.expr != nil
}

// switch typ {
// 	// case Bool:
// 	// 	col.serie = serie.NewBool()
// 	case Int:
// 		col.serie = serie.NewInt()
// 	case Int8:
// 		col.serie = serie.NewInt8()
// 	case Int16:
// 		col.serie = serie.NewInt16()
// 	case Int32:
// 		col.serie = serie.NewInt32()
// 	case Int64:
// 		col.serie = serie.NewInt64()
// 	case Uint:
// 		col.serie = serie.NewUint()
// 	case Uint8:
// 		col.serie = serie.NewUint8()
// 	case Uint16:
// 		col.serie = serie.NewUint16()
// 	case Uint32:
// 		col.serie = serie.NewUint32()
// 	case Uint64:
// 		col.serie = serie.NewUint64()
// 	case Float32:
// 		col.serie = serie.NewFloat32()
// 	case Float64:
// 		col.serie = serie.NewFloat64()
// 	case String:
// 		col.serie = serie.NewString()
// 	// case Datetime:
// 	// 	col.serie = serie.NewDatetime()
// 	// case Duration:
// 	// 	col.serie = serie.NewDuration()
// 	default:
// 		col.serie = serie.NewRaw()
// 	}
