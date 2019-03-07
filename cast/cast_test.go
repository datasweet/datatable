package cast_test

import (
	"testing"
	"time"

	"github.com/datasweet/datatable/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsBool(t *testing.T) {
	testBool(t, nil, false, false)
	testBool(t, false, false, true)
	testBool(t, true, true, true)
	testBool(t, 1, true, true)
	testBool(t, int8(1), true, true)
	testBool(t, int16(1), true, true)
	testBool(t, int32(1), true, true)
	testBool(t, int64(1), true, true)
	testBool(t, uint(1), true, true)
	testBool(t, uint8(1), true, true)
	testBool(t, uint16(1), true, true)
	testBool(t, uint32(1), true, true)
	testBool(t, uint64(1), true, true)
	testBool(t, float32(1), true, true)
	testBool(t, float64(1), true, true)
	testBool(t, "true", true, true)
	testBool(t, "wrong", false, false)
}

func TestAsString(t *testing.T) {
	testString(t, nil, "", false)
	testString(t, "hello", "hello", true)
	testString(t, 123, "123", true)
	testString(t, int8(123), "123", true)
	testString(t, int16(123), "123", true)
	testString(t, int32(123), "123", true)
	testString(t, int64(123), "123", true)
	testString(t, uint(123), "123", true)
	testString(t, uint8(123), "123", true)
	testString(t, uint16(123), "123", true)
	testString(t, uint32(123), "123", true)
	testString(t, uint64(123), "123", true)
	testString(t, float32(3.14), "3.14", true)
	testString(t, float64(3.14), "3.14", true)
	testString(t, true, "true", true)
}

func TestAsInt(t *testing.T) {
	testInt(t, "123", 123, true)
	testInt(t, "wrong", 0, false)
	testInt(t, true, 1, true)
	testInt(t, false, 0, true)
	testInt(t, 123, 123, true)
	testInt(t, int8(123), 123, true)
	testInt(t, int16(123), 123, true)
	testInt(t, int32(123), 123, true)
	testInt(t, int64(123), 123, true)
	testInt(t, uint(123), 123, true)
	testInt(t, uint8(123), 123, true)
	testInt(t, uint16(123), 123, true)
	testInt(t, uint32(123), 123, true)
	testInt(t, uint64(123), 123, true)
	testInt(t, float32(123), 123, true)
	testInt(t, float64(123), 123, true)
}

func TestAsFloat(t *testing.T) {
	testFloat(t, "123", 123, true)
	testFloat(t, "wrong", 0, false)
	testFloat(t, true, 1, true)
	testFloat(t, false, 0, true)
	testFloat(t, 123, 123, true)
	testFloat(t, int8(123), 123, true)
	testFloat(t, int16(123), 123, true)
	testFloat(t, int32(123), 123, true)
	testFloat(t, int64(123), 123, true)
	testFloat(t, uint(123), 123, true)
	testFloat(t, uint8(123), 123, true)
	testFloat(t, uint16(123), 123, true)
	testFloat(t, uint32(123), 123, true)
	testFloat(t, uint64(123), 123, true)
	testFloat(t, float32(123), 123, true)
	testFloat(t, float64(123), 123, true)
}

func TestAsDatetime(t *testing.T) {
	date := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC()
	testDatetime(t, date, date, true)
	testDatetime(t, 1551435220270, date, true)
	testDatetime(t, int64(1551435220270), date, true)
	testDatetime(t, "1551435220270", date, true)
	testDatetime(t, "2019-03-01T10:13:40.27Z", date, true)
	testDatetime(t, "wrong", time.Time{}, false)
}

func testBool(t *testing.T, value interface{}, expected bool, ok bool) {
	b, o := cast.AsBool(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testString(t *testing.T, value interface{}, expected string, ok bool) {
	b, o := cast.AsString(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt(t *testing.T, value interface{}, expected int64, ok bool) {
	b, o := cast.AsInt(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testFloat(t *testing.T, value interface{}, expected float64, ok bool) {
	b, o := cast.AsFloat(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testDatetime(t *testing.T, value interface{}, expected time.Time, ok bool) {
	b, o := cast.AsDatetime(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
