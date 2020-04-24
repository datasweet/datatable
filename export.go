package datatable

type ExportOptions struct {
	HiddenColumns  []string
	IncludeHeaders bool
}

type ExportOption func(opts *ExportOptions)

func newExportOptions(opt ...ExportOption) ExportOptions {
	options := ExportOptions{}
	for _, o := range opt {
		o(&options)
	}
	return options
}
