package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/datasweet/cast"
	"github.com/datasweet/datatable"
	"github.com/pkg/errors"
)

// Import a csv
func Import(name, path string, opt ...Option) (*datatable.DataTable, error) {
	options := Options{
		IgnoreIfReadLineError: true,
		Comma:                 ',',
		TrimLeadingSpace:      true,
	}
	for _, o := range opt {
		o(&options)
	}

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = options.Comma
	reader.Comment = options.Comment
	reader.LazyQuotes = options.LazyQuotes
	reader.TrimLeadingSpace = options.TrimLeadingSpace

	dt := datatable.New(name)

	line := 1

	// Get columns names with headers
	if options.HasHeaders {
		rec, err := reader.Read()
		if err != nil {
			return nil, errors.Wrap(err, "can't read headers")
		}
		if len(options.ColumnNames) == 0 {
			options.ColumnNames = append(options.ColumnNames, rec...)
		}
		line++
	}

	for {
		rec, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				if line == 1 {
					return nil, errors.New("nil datas")
				}
				break
			}
			if options.IgnoreIfReadLineError {
				continue
			}
			return nil, errors.Wrapf(err, "error line %d", line)
		}

		// Do we have columns names ?
		if len(options.ColumnNames) == 0 {
			options.ColumnNames = make([]string, 0, len(rec))
			for i := range rec {
				options.ColumnNames = append(options.ColumnNames, fmt.Sprintf("col %d", i+1))
			}
		}

		// Do we have columns in datatable
		if dt.NumCols() == 0 {
			// Detect type if needed
			if len(options.ColumnTypes) == 0 {
				options.ColumnTypes = detectTypes(rec, options.DateFormats)
			}
			if len(options.ColumnNames) != len(options.ColumnTypes) {
				return nil, errors.Errorf("expected %d types, got %d", len(options.ColumnNames), len(options.ColumnTypes))
			}
			for i := range options.ColumnNames {
				if err := dt.AddColumn(options.ColumnNames[i], options.ColumnTypes[i], datatable.ColumnTimeFormats(options.DateFormats...)); err != nil {
					return nil, errors.Wrapf(err, "add column '%s' with type '%s'", options.ColumnNames[i], options.ColumnTypes[i])
				}
			}
		}

		// conv => []interface{}
		cells := make([]interface{}, 0, len(rec))
		for _, r := range rec {
			cells = append(cells, r)
		}
		dt.AppendRow(cells...)
		line++
	}

	return dt, nil
}

func detectTypes(rec, dateformat []string) []datatable.ColumnType {
	ctypes := make([]datatable.ColumnType, 0, len(rec))
	for _, r := range rec {
		if _, ok := cast.AsFloat64(r); ok {
			ctypes = append(ctypes, datatable.Float64)
			continue
		}
		if _, ok := cast.AsBool(r); ok {
			ctypes = append(ctypes, datatable.Bool)
			continue
		}
		if _, ok := cast.AsTime(r, dateformat...); ok {
			ctypes = append(ctypes, datatable.Time)
			continue
		}
		ctypes = append(ctypes, datatable.String)
	}
	return ctypes
}
