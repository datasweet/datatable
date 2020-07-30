package datatable

import (
	"fmt"
	"strings"

	"github.com/datasweet/datatable/serie"
	"github.com/pkg/errors"
)

var stypes map[string]func() serie.Serie

func init() {
	stypes = make(map[string]func() serie.Serie)
	RegisterType(Boolean, func() serie.Serie { return serie.BoolN() })
	RegisterType(String, func() serie.Serie { return serie.StringN() })
	RegisterType(Int, func() serie.Serie { return serie.IntN() })
	//RegisterType(Int8, func() serie.Serie { return serie.Int8N() })
	//RegisterType(Int16, func() serie.Serie { return serie.Int16N() })
	RegisterType(Int32, func() serie.Serie { return serie.Int32N() })
	RegisterType(Int64, func() serie.Serie { return serie.Int64N() })
	// RegisterType(UInt, func() serie.Serie { return serie.UIntN() })
	// RegisterType(UInt8, func() serie.Serie { return serie.UInt8N() })
	// RegisterType(UInt16, func() serie.Serie { return serie.UInt16N() })
	// RegisterType(UInt32, func() serie.Serie { return serie.UInt32N() })
	// RegisterType(UInt64, func() serie.Serie { return serie.UInt64N() })
	RegisterType(Float32, func() serie.Serie { return serie.Float32N() })
	RegisterType(Float64, func() serie.Serie { return serie.Float64N() })
	RegisterType(Time, func() serie.Serie { return serie.TimeN() })
	RegisterType(Raw, func() serie.Serie { return serie.Raw() })
}

const (
	Boolean = "boolean"
	String  = "string"
	Int     = "int"
	// Int8      = "int8"
	// Int16     = "int16"
	Int32 = "int32"
	Int64 = "int64"
	// Uint  = "uint"
	// Uint8     = "uint8"
	// Uint16    = "uint16"
	// Uint32    = "uint32"
	// Uint64    = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	Time    = "time"
	Raw     = "raw"
)

// RegisterType to extends the known type
func RegisterType(name string, serier func() serie.Serie) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("empty name")
	}
	if serier == nil {
		return errors.New("nil factory")
	}
	if _, ok := stypes[name]; ok {
		return errors.Errorf("type '%s' already exists", name)
	}
	stypes[name] = serier
	return nil
}

// newSerie to create a serie from a known type
func newSerie(typ string) serie.Serie {
	if s, ok := stypes[typ]; ok {
		return s()
	}
	panic(fmt.Sprintf("unknown serie type '%s'", typ))
}
