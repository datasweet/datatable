package datatable

type Type uint8

const (
	Raw Type = iota
	Bool
	Int8
	Int16
	Int32
	Int64
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

// func (typ Type) NewSerie(v ...interface{}) serie.Serie {
// 	switch typ {
// 	case Raw:
// 		return serie.NewRaw(v...)
// 	case Bool:
// 		return serie.NewBool(v...)
// 	case Int8:
// 		return serie.NewInt8(v...)
// 	case Int16:
// 		return serie.NewInt16(v...)
// 	case Int32:
// 		return serie.NewInt32(v...)
// 	case Int64:
// 		return serie.NewInt64(v...)
// 	case Uint8:
// 		return serie.NewUint8(v...)
// 	case Uint16:
// 		return serie.NewUint16(v...)
// 	case Uint32:
// 		return serie.NewUint32(v...)
// 	case UInt64:
// 		return serie.NewUint64(v...)
// 	case String:
// 		return serie.NewString(v...)
// 	default:
// 		return nil
// 	}
// }
