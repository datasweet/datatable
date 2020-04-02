package datatable

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

type PrintOptions struct {
	ColumnName bool
	ColumnType bool
	RowNumber  bool
	MaxRows    int
}

type PrintOption func(opts *PrintOptions)

func PrintColumnName(v bool) PrintOption {
	return func(opts *PrintOptions) {
		opts.ColumnName = v
	}
}

func PrintColumnType(v bool) PrintOption {
	return func(opts *PrintOptions) {
		opts.ColumnType = v
	}
}

func PrintRowNumber(v bool) PrintOption {
	return func(opts *PrintOptions) {
		opts.RowNumber = v
	}
}

func PrintMaxRows(v int) PrintOption {
	return func(opts *PrintOptions) {
		opts.MaxRows = v
	}
}

func (t *table) Print(opt ...PrintOption) string {
	options := PrintOptions{
		ColumnName: true,
		ColumnType: true,
		RowNumber:  true,
		MaxRows:    100,
	}

	for _, o := range opt {
		o(&options)
	}

	var sb strings.Builder
	tw := tablewriter.NewWriter(&sb)
	tw.SetAutoWrapText(false)
	tw.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)
	tw.SetCenterSeparator("")
	tw.SetColumnSeparator("")
	tw.SetRowSeparator("")
	tw.SetHeaderLine(false)
	tw.SetBorder(false)
	tw.SetTablePadding("\t")
	tw.SetNoWhiteSpace(true)

	// if options.ColumnName || options.ColumnType {

	// table.SetHeader([]string{"Name", "Status", "Role", "Version"})
	// }

	//tw.SetHeader([]string{"Name", "Sign", "Rating"})

	tw.Render()
	return sb.String()
}
