package serie

import (
	"reflect"

	"github.com/datasweet/cast"
)

// AsFloat64 to converts a serie to a serie of float64
// Used for some statistics
func AsFloat64(s Serie, missing *float64) Serie {
	if s == nil || s.Len() == 0 {
		return Float64() // empty list of floats
	}

	switch kind := s.Type().Kind(); kind {
	case reflect.Float64:
		return s
	default:
		ln := s.Len()
		arr := make([]float64, 0, ln)
		for i := 0; i < ln; i++ {
			if f, ok := cast.AsFloat(s.Get(i)); ok {
				arr = append(arr, f)
				continue
			}

			if missing != nil {
				arr = append(arr, *missing)
			}
		}

		sf := Float64()
		(sf.(*serie)).slice = reflect.ValueOf(arr)
		return sf
	}
}
