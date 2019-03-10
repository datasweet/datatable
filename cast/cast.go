package cast

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var (
	DateFormats = []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
)

// AsBool to convert as bool
func AsBool(v interface{}) (bool, bool) {
	switch d := v.(type) {
	case bool:
		return d, true
	case int:
		return d > 0, true
	case int8:
		return d > 0, true
	case int16:
		return d > 0, true
	case int32:
		return d > 0, true
	case int64:
		return d > 0, true
	case uint:
		return d > 0, true
	case uint8:
		return d > 0, true
	case uint16:
		return d > 0, true
	case uint32:
		return d > 0, true
	case uint64:
		return d > 0, true
	case float32:
		return d > 0, true
	case float64:
		return d > 0, true
	case string:
		if b, err := strconv.ParseBool(d); err == nil {
			return b, true
		}
		return false, false

	default:
		return false, false
	}

}

// AsString to convert as string
func AsString(v interface{}) (string, bool) {
	switch d := v.(type) {
	case string:
		return d, true
	case int:
		return strconv.FormatInt(int64(d), 10), true
	case int8:
		return strconv.FormatInt(int64(d), 10), true
	case int16:
		return strconv.FormatInt(int64(d), 10), true
	case int32:
		return strconv.FormatInt(int64(d), 10), true
	case int64:
		return strconv.FormatInt(d, 10), true
	case uint:
		return strconv.FormatUint(uint64(d), 10), true
	case uint8:
		return strconv.FormatUint(uint64(d), 10), true
	case uint16:
		return strconv.FormatUint(uint64(d), 10), true
	case uint32:
		return strconv.FormatUint(uint64(d), 10), true
	case uint64:
		return strconv.FormatUint(uint64(d), 10), true
	case float32:
		return strconv.FormatFloat(float64(d), 'f', -1, 32), true
	case float64:
		return strconv.FormatFloat(d, 'f', -1, 64), true
	case json.Number:
		return d.String(), true
	case bool:
		return strconv.FormatBool(d), true
	default:
		if v == nil {
			return "", false
		}

		return fmt.Sprintf("%v", v), false
	}
}

// AsInt to convert as a int
func AsInt(v interface{}) (int64, bool) {
	switch d := v.(type) {
	case int:
		return int64(d), true
	case int8:
		return int64(d), true
	case int16:
		return int64(d), true
	case int32:
		return int64(d), true
	case int64:
		return d, true
	case float32:
		return int64(d), true
	case float64:
		return int64(d), true
	case uint:
		return int64(d), true
	case uint8:
		return int64(d), true
	case uint16:
		return int64(d), true
	case uint32:
		return int64(d), true
	case uint64:
		return int64(d), true
	case json.Number:
		if f, err := d.Int64(); err == nil {
			return f, true
		}
		return 0, false
	case string:
		if i, err := strconv.ParseInt(d, 10, 64); err == nil {
			return i, true
		}
		return 0, false
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// AsIntArray to convert as an array of int64
func AsIntArray(values ...interface{}) ([]int64, bool) {
	arr := make([]int64, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsInt(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsFloat to convert as a float64
func AsFloat(v interface{}) (float64, bool) {
	switch d := v.(type) {
	case float64:
		return d, true
	case float32:
		return float64(d), true
	case int:
		return float64(d), true
	case int8:
		return float64(d), true
	case int16:
		return float64(d), true
	case int32:
		return float64(d), true
	case int64:
		return float64(d), true
	case uint:
		return float64(d), true
	case uint8:
		return float64(d), true
	case uint16:
		return float64(d), true
	case uint32:
		return float64(d), true
	case uint64:
		return float64(d), true
	case json.Number:
		if f, err := d.Float64(); err == nil {
			return f, true
		}
		return 0, false
	case string:
		if f, err := strconv.ParseFloat(d, 64); err == nil {
			return f, true
		}
		return 0, false
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// AsFloatArray to convert as an array of float64
func AsFloatArray(values ...interface{}) ([]float64, bool) {
	arr := make([]float64, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsFloat(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsDatetime to convert as datetime (time.Time)
func AsDatetime(v interface{}) (time.Time, bool) {
	switch d := v.(type) {
	case time.Time:
		return d.UTC(), true
	case int: // timestamp compliant with javascript
		return time.Unix(0, int64(d)*int64(time.Millisecond)).UTC(), true
	case int64:
		return time.Unix(0, d*int64(time.Millisecond)).UTC(), true
	case string:
		// Try to convert to int64
		if ts, err := strconv.ParseInt(d, 10, 64); err == nil {
			return time.Unix(0, ts*int64(time.Millisecond)).UTC(), true
		}

		// Try formats
		for _, format := range DateFormats {
			if dt, err := time.Parse(format, d); err == nil {
				return dt.UTC(), true
			}
		}
		return time.Time{}, false
	default:
		return time.Time{}, false
	}
}
