package datatable

import (
	"reflect"
	"strings"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/expr"
	"github.com/pkg/errors"
)

// ColumnType defines the valid column type in datatable
type ColumnType string

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

// ColumnOptions describes options to be apply on a column
type ColumnOptions struct {
	Hidden      bool
	Expr        string
	Values      []interface{}
	TimeFormats []string
}

// ColumnOption sets column options
type ColumnOption func(opts *ColumnOptions)

// ColumnHidden sets the visibility
func ColumnHidden(v bool) ColumnOption {
	return func(opts *ColumnOptions) {
		opts.Hidden = v
	}
}

// Expr sets the expr for the column
// <!> Incompatible with ColumnValues
func Expr(v string) ColumnOption {
	return func(opts *ColumnOptions) {
		opts.Expr = v
	}
}

// Values fills the column with the values
// <!> Incompatible with ColumnExpr
func Values(v ...interface{}) ColumnOption {
	return func(opts *ColumnOptions) {
		opts.Values = v
	}
}

// TimeFormats sets the valid time formats.
// <!> Only for Time Column
func TimeFormats(v ...string) ColumnOption {
	return func(opts *ColumnOptions) {
		opts.TimeFormats = append(opts.TimeFormats, v...)
	}
}

// ColumnSerier to create a serie from column options
type ColumnSerier func(ColumnOptions) serie.Serie

// ctypes is our column type registry
var ctypes map[ColumnType]ColumnSerier

func init() {
	ctypes = make(map[ColumnType]ColumnSerier)
	RegisterColumnType(Bool, func(opts ColumnOptions) serie.Serie {
		return serie.BoolN(opts.Values...)
	})
	RegisterColumnType(String, func(opts ColumnOptions) serie.Serie {
		return serie.StringN(opts.Values...)
	})
	RegisterColumnType(Int, func(opts ColumnOptions) serie.Serie {
		return serie.IntN(opts.Values...)
	})
	RegisterColumnType(Int32, func(opts ColumnOptions) serie.Serie {
		return serie.Int32N(opts.Values...)
	})
	RegisterColumnType(Int64, func(opts ColumnOptions) serie.Serie {
		return serie.Int64N(opts.Values...)
	})
	RegisterColumnType(Float32, func(opts ColumnOptions) serie.Serie {
		return serie.Float32N(opts.Values...)
	})
	RegisterColumnType(Float64, func(opts ColumnOptions) serie.Serie {
		return serie.Float64N(opts.Values...)
	})
	RegisterColumnType(Time, func(opts ColumnOptions) serie.Serie {
		sr := serie.TimeN(opts.TimeFormats...)
		if len(opts.Values) > 0 {
			sr.Append(opts.Values...)
		}
		return sr
	})
	RegisterColumnType(Raw, func(opts ColumnOptions) serie.Serie {
		return serie.Raw(opts.Values...)
	})
}

// RegisterColumnType to extends the known type
func RegisterColumnType(name ColumnType, serier ColumnSerier) error {
	name = ColumnType(strings.TrimSpace(string(name)))
	if len(name) == 0 {
		return ErrEmptyName
	}
	if serier == nil {
		return ErrNilFactory
	}
	if _, ok := ctypes[name]; ok {
		err := errors.Errorf("type '%s' already exists", name)
		return errors.Wrap(err, ErrTypeAlreadyExists.Error())
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

// newColumnSerie to create a serie from a known type
func newColumnSerie(ctyp ColumnType, options ColumnOptions) (serie.Serie, error) {
	if s, ok := ctypes[ctyp]; ok {
		return s(options), nil
	}
	err := errors.Errorf("unknown column type '%s'", ctyp)
	return nil, errors.Wrap(err, ErrUnknownColumnType.Error())
}

// Column describes a column in our datatable
type Column interface {
	Name() string
	Type() ColumnType
	UnderlyingType() reflect.Type
	IsVisible() bool
	IsComputed() bool
	//Clone(includeValues bool) Column
}

type column struct {
	name     string
	typ      ColumnType
	hidden   bool
	formulae string
	expr     expr.Node
	serie    serie.Serie
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Type() ColumnType {
	return c.typ
}

func (c *column) UnderlyingType() reflect.Type {
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
		typ:      c.typ,
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
		typ:      c.typ,
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
