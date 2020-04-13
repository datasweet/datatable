package value_test

import (
	"testing"
	"time"

	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	val := value.Time()
	assert.Equal(t, value.TimeType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, time.Time{}, val.Val())

	val = value.Time("2018-06-04")
	assert.Equal(t, value.TimeType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, time.Date(2018, time.June, 04, 0, 0, 0, 0, time.UTC), val.Val())

	val.Set("teemo")
	assert.False(t, val.IsValid())
	assert.Equal(t, nil, val.Val())
}

func TestCloneTime(t *testing.T) {
	val := value.Time("2018-06-04")
	assert.NotNil(t, val)
	assert.Equal(t, value.TimeType, val.Type())
	assert.True(t, val.IsValid())
	assert.Equal(t, time.Date(2018, time.June, 04, 0, 0, 0, 0, time.UTC), val.Val())

	cpy := val.Clone()
	assert.NotNil(t, cpy)
	assert.NotSame(t, val, cpy)
	assert.Equal(t, value.TimeType, cpy.Type())
	assert.True(t, cpy.IsValid())
	assert.Equal(t, time.Date(2018, time.June, 04, 0, 0, 0, 0, time.UTC), cpy.Val())
}

func TestCompareTime(t *testing.T) {
	a := value.Time("2018-06-04")
	assert.Equal(t, value.Eq, a.Compare(value.Time(time.Date(2018, time.June, 04, 0, 0, 0, 0, time.UTC))))
	assert.Equal(t, value.Gt, a.Compare(nil))
	assert.Equal(t, value.Gt, a.Compare(value.Time("2018-06-01")))

	// convert type
	assert.Equal(t, value.Eq, a.Compare(value.Int(1528070400000)))
	assert.Equal(t, value.Eq, a.Compare(value.String("2018-06-04")))

	a.Set("2017-06-04")
	assert.Equal(t, value.Lt, a.Compare(value.Time("2018-06-04")))
	assert.Equal(t, value.Eq, a.Compare(value.Time("2017-06-04")))
	assert.Equal(t, value.Gt, a.Compare(nil))
}
