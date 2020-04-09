package serie

import (
	"strconv"
	"strings"
)

type PrintOptions struct {
	WithType       bool
	WithRowNumber  bool
	ValueSeparator string
	MaxRows        int
}

type PrintOption func(opts *PrintOptions)

func PrintType(v bool) PrintOption {
	return func(opts *PrintOptions) {
		opts.WithType = v
	}
}

func PrintRowNumber(v bool) PrintOption {
	return func(opts *PrintOptions) {
		opts.WithRowNumber = v
	}
}

func PrintValueSeparator(v string) PrintOption {
	return func(opts *PrintOptions) {
		opts.ValueSeparator = v
	}
}

func PrintMaxRows(v int) PrintOption {
	return func(opts *PrintOptions) {
		opts.MaxRows = v
	}
}

// Print to print the serie
func (s *serie) Print(opt ...PrintOption) string {
	options := PrintOptions{
		WithType:       true,
		WithRowNumber:  true,
		MaxRows:        100,
		ValueSeparator: "\n",
	}
	for _, o := range opt {
		o(&options)
	}

	var sb strings.Builder
	if options.WithType {
		if options.WithRowNumber {
			sb.WriteString(strings.Repeat(" ", len(strconv.Itoa(len(s.values)))+2))
		}
		sb.WriteString("<")
		sb.WriteString(string(s.typ))
		sb.WriteString(">")
		sb.WriteString(options.ValueSeparator)
	}

	if len(s.values) == 0 {
		sb.WriteString("nil")
		return sb.String()
	}

	for i, val := range s.values {
		if options.WithRowNumber {
			sb.WriteString(strconv.Itoa(i + 1))
			sb.WriteString(": ")
		}
		sb.WriteString(val.String())
		sb.WriteString(options.ValueSeparator)
	}
	return sb.String()
}

func (s *serie) String() string {
	return s.Print()
}
