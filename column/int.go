package column

import "github.com/datasweet/datatable/value"

// IntOptions defines some options to an int column
type IntOptions struct {
	NilValue  *int64
	Formatter string // pattern printf
}

// IntOption is an abstract setter
type IntOption func(opts *IntOptions)

type intColumn struct {
	*column
	options IntOptions
	builder value.Builder
}

func newInt(name string, builder value.Builder, opt ...IntOption) *intColumn {
	options := IntOptions{}
	for _, o := range opt {
		o(&options)
	}
	return &intColumn{
		column: &column{
			name: name,
			typ:  builder(nil).Type(),
		},
		options: options,
		builder: builder,
	}
}

func (c *intColumn) NewValue() value.Value {
	return c.builder(nil)
}

// Int to create an int column
func Int(name string, opt ...IntOption) ColumnValue {
	return newInt(name, value.NewInt, opt...)
}

// Int8 to create an int8 column
func Int8(name string, opt ...IntOption) ColumnValue {
	return newInt(name, value.NewInt8, opt...)
}

// Int16 to create an int16 column
func Int16(name string, opt ...IntOption) ColumnValue {
	return newInt(name, value.NewInt16, opt...)
}

// Int32 to create an int32 column
func Int32(name string, opt ...IntOption) ColumnValue {
	return newInt(name, value.NewInt32, opt...)
}

// Int64 to create an int64 column
func Int64(name string, opt ...IntOption) ColumnValue {
	return newInt(name, value.NewInt64, opt...)
}
