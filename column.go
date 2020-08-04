package datatable

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

// ColumnType defines the valid column type in datatable
type ColumnType string

var ctypes map[ColumnType]func() serie.Serie

func init() {
	ctypes = make(map[ColumnType]func() serie.Serie)
	RegisterColumnType(Bool, func() serie.Serie { return serie.BoolN() })
	RegisterColumnType(String, func() serie.Serie { return serie.StringN() })
	RegisterColumnType(Int, func() serie.Serie { return serie.IntN() })
	//RegisterType(Int8, func() serie.Serie { return serie.Int8N() })
	//RegisterType(Int16, func() serie.Serie { return serie.Int16N() })
	RegisterColumnType(Int32, func() serie.Serie { return serie.Int32N() })
	RegisterColumnType(Int64, func() serie.Serie { return serie.Int64N() })
	// RegisterType(UInt, func() serie.Serie { return serie.UIntN() })
	// RegisterType(UInt8, func() serie.Serie { return serie.UInt8N() })
	// RegisterType(UInt16, func() serie.Serie { return serie.UInt16N() })
	// RegisterType(UInt32, func() serie.Serie { return serie.UInt32N() })
	// RegisterType(UInt64, func() serie.Serie { return serie.UInt64N() })
	RegisterColumnType(Float32, func() serie.Serie { return serie.Float32N() })
	RegisterColumnType(Float64, func() serie.Serie { return serie.Float64N() })
	RegisterColumnType(Time, func() serie.Serie { return serie.TimeN() })
	RegisterColumnType(Raw, func() serie.Serie { return serie.Raw() })
}

const (
	Bool   ColumnType = "bool"
	String ColumnType = "string"
	Int    ColumnType = "int"
	// Int8     ColumnType = "int8"
	// Int16    ColumnType = "int16"
	Int32 ColumnType = "int32"
	Int64 ColumnType = "int64"
	// Uint  ColumnType = "uint"
	// Uint8     ColumnType = "uint8"
	// Uint16    ColumnType = "uint16"
	// Uint32    ColumnType = "uint32"
	// Uint64    ColumnType = "uint64"
	Float32 ColumnType = "float32"
	Float64 ColumnType = "float64"
	Time    ColumnType = "time"
	Raw     ColumnType = "raw"
)

// RegisterColumnType to extends the known type
func RegisterColumnType(name ColumnType, serier func() serie.Serie) error {
	name = ColumnType(strings.TrimSpace(string(name)))
	if len(name) == 0 {
		return errors.New("empty name")
	}
	if serier == nil {
		return errors.New("nil factory")
	}
	if _, ok := ctypes[name]; ok {
		return errors.Errorf("type '%s' already exists", name)
	}
	ctypes[name] = serier
	return nil
}

// ColumnTypes to list all column type
func ColumnTypes() []ColumnType {
	ctyp := make([]ColumnType, 0, len(ctypes))
	for k := range ctypes {
		ctyp = append(ctyp, k)
	}
	return ctyp
}

// newSerie to create a serie from a known type
func newSerie(ctyp ColumnType) serie.Serie {
	if s, ok := ctypes[ctyp]; ok {
		return s()
	}
	panic(fmt.Sprintf("unknown column type '%s'", ctyp))
}

// Column describes a column in our datatable
type Column interface {
	Name() string
	Type() reflect.Type
	IsVisible() bool
	IsComputed() bool
	//Clone(includeValues bool) Column
}

type column struct {
	name     string
	hidden   bool
	formulae string
	expr     expr.Node
	serie    serie.Serie
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Type() reflect.Type {
	return c.serie.Type()
}

func (c *column) IsVisible() bool {
	return !c.hidden
}

func (c *column) IsComputed() bool {
	return len(c.formulae) > 0
}

func (c *column) emptyCopy() *column {
	cpy := &column{
		name:     c.name,
		hidden:   c.hidden,
		formulae: c.formulae,
		serie:    c.serie.EmptyCopy(),
	}
	if len(cpy.formulae) > 0 {
		if parsed, err := expr.Parse(cpy.formulae); err == nil {
			cpy.expr = parsed
		}
	}
	return cpy
}

func (c *column) copy() *column {
	cpy := &column{
		name:     c.name,
		hidden:   c.hidden,
		formulae: c.formulae,
		serie:    c.serie.Copy(),
	}
	if len(cpy.formulae) > 0 {
		if parsed, err := expr.Parse(cpy.formulae); err == nil {
			cpy.expr = parsed
		}
	}
	return cpy
}
