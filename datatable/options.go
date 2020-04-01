package datatable

type Options struct {
	Indexes []string
}

// ColumnOptions are options when a columns is added to a datatable
type ColumnOptions struct {
	Name string
	//Type serie.ValueType
}

type ColumnOption func(opts *ColumnOptions)
