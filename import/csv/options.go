package csv

import "github.com/datasweet/datatable"

// Options are options to import a csv
type Options struct {
	HasHeaders            bool
	ColumnNames           []string               // if len == 0 => take headers else "col #i"
	ColumnTypes           []datatable.ColumnType // if len == 0 => detection
	IgnoreIfReadLineError bool
	Comma                 rune
	Comment               rune
	LazyQuotes            bool
	TrimLeadingSpace      bool
	DateFormats           []string
}

// Option is a setter
type Option func(*Options)

// HasHeader to retrieve column names on line #1
func HasHeader(v bool) Option {
	return func(opts *Options) {
		opts.HasHeaders = v
	}
}

// ColumnNames defines the column name
func ColumnNames(v ...string) Option {
	return func(opts *Options) {
		opts.ColumnNames = v
	}
}

// ColumnTypes defines the column type
func ColumnTypes(v ...datatable.ColumnType) Option {
	return func(opts *Options) {
		opts.ColumnTypes = v
	}
}

// IgnoreLineWithError to not stop the reading process if a line has an error
func IgnoreLineWithError(v bool) Option {
	return func(opts *Options) {
		opts.IgnoreIfReadLineError = v
	}
}

// Comma is the field delimiter
// Default to ','
func Comma(v rune) Option {
	return func(opts *Options) {
		opts.Comma = v
	}
}

// Comment if not 0, is the comment character. Lines beginning with the
// Comment character without preceding whitespace are ignored.
func Comment(v rune) Option {
	return func(opts *Options) {
		opts.Comment = v
	}
}

// LazyQuotes is true, a quote may appear in an unquoted field and a
// non-doubled quote may appear in a quoted field.
func LazyQuotes(v bool) Option {
	return func(opts *Options) {
		opts.LazyQuotes = v
	}
}

// TrimLeadingSpace is true, leading white space in a field is ignored.
// This is done even if the field delimiter, Comma, is white space.
func TrimLeadingSpace(v bool) Option {
	return func(opts *Options) {
		opts.TrimLeadingSpace = v
	}
}

// AcceptDate to accept a specific date format
func AcceptDate(v string) Option {
	return func(opts *Options) {
		opts.DateFormats = append(opts.DateFormats, v)
	}
}
