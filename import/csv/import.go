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
		return nil, errors.Wrap(err, datatable.ErrOpenFile.Error())
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
			return nil, errors.Wrap(err, datatable.ErrCantReadHeaders.Error())
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
					return nil, datatable.ErrNilDatas
				}
				break
			}
			if options.IgnoreIfReadLineError {
				continue
			}
			err := errors.Wrapf(err, "error line %d", line)
			return nil, errors.Wrap(err, datatable.ErrReadingLine.Error())
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
				err := errors.Errorf("expected %d types, got %d", len(options.ColumnNames), len(options.ColumnTypes))
				return nil, errors.Wrap(err, datatable.ErrWrongNumberOfTypes.Error())
			}
			for i := range options.ColumnNames {
				if err := dt.AddColumn(options.ColumnNames[i], options.ColumnTypes[i], datatable.TimeFormats(options.DateFormats...)); err != nil {
					err = errors.Wrapf(err, "add column '%s' with type '%s'", options.ColumnNames[i], options.ColumnTypes[i])
					return nil, errors.Wrap(err, datatable.ErrAddingColumn.Error())
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
