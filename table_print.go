package datatable

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// PrintOptions to control the printer
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

// Print the tables with options
func (t *DataTable) Print(writer io.Writer, opt ...PrintOption) {
	options := PrintOptions{
		ColumnName: true,
		ColumnType: true,
		RowNumber:  true,
		MaxRows:    100,
	}

	for _, o := range opt {
		o(&options)
	}

	if writer == nil {
		writer = os.Stdout
	}

	tw := tablewriter.NewWriter(writer)
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

	if options.ColumnName || options.ColumnType {
		headers := make([]string, 0, len(t.cols))

		for _, col := range t.cols {
			var h []string
			if options.ColumnName {
				h = append(h, col.Name())
			}
			if options.ColumnType {
				h = append(h, fmt.Sprintf("<%s>", col.serie.Type().Name()))
			}
			headers = append(headers, strings.Join(h, " "))
		}
		tw.SetHeader(headers)
	}

	if options.MaxRows > 1 && options.MaxRows <= t.NumRows() {
		mr := options.MaxRows / 2
		tw.AppendBulk(t.Head(mr).Records())
		seps := make([]string, 0, len(t.cols))
		for range t.cols {
			seps = append(seps, "...")
		}
		tw.Append(seps)
		tw.AppendBulk(t.Tail(mr).Records())
	} else {
		tw.AppendBulk(t.Records())
	}

	tw.Render()
}
