package value

import (
	"time"

	"github.com/datasweet/cast"
)

const TimeType = Type("time")

type timeValue struct {
	val  time.Time
	null bool
}

// Time to create a new time value
// if v is a string => formats accepted are :
// - "2006-01-02",
// - "2006-01-02 15:04:05",
// - time.RFC3339,
// - "02/01/2006",
// - "02/01/2006 15:04:05",
// - "20060102",
// - "20060102150405",
func Time(v ...interface{}) Value {
	value := &timeValue{}
	if len(v) == 1 {
		value.Set(v[0])
	}
	return value
}

func (value *timeValue) Type() Type {
	return TimeType
}

func (value *timeValue) Val() interface{} {
	if value.null {
		return nil
	}
	return value.val
}

func (value *timeValue) Set(v interface{}) Value {
	value.val = time.Time{}
	value.null = true

	if casted, ok := cast.AsDatetime(v); ok {
		value.val = casted
		value.null = false
	}

	return value
}

func (value *timeValue) IsValid() bool {
	return !value.null
}

// Compare the current 'value' 'to' an other value
// returns -1 if value < to, 0 if value == to, 1 if value > to
// nil | not a Time < value
func (value *timeValue) Compare(to Value) int {
	if to == nil {
		return Gt
	}

	tv, ok := to.(*timeValue)
	if !ok {
		// try to convert
		tv = Time(to.Val()).(*timeValue)
	}

	if tv.null {
		if value.null {
			return Eq
		}
		return Gt
	}

	if value.val.Equal(tv.val) {
		return Eq
	}

	if value.val.After(tv.val) {
		return Gt
	}

	return Lt
}

func (value *timeValue) Clone() Value {
	var val timeValue
	val = *value
	return &val
}

func (value *timeValue) String() string {
	if value.null {
		return nullValueStr
	}
	return value.val.String()

}
