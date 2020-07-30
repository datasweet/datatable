package read

type CsvOptions struct {
	HasHeaders            bool
	DetectType            bool
	IgnoreIfReadLineError bool
	Comma                 rune
	Comment               rune
	LazyQuotes            bool
	TrimLeadingSpace      bool
}

type CsvOption func(*CsvOptions)

// func CSV(path string, opt ...CsvOption) (*datatable.DataTable, error) {
// 	options := CsvOptions{
// 		DetectType:            true,
// 		IgnoreIfReadLineError: true,
// 		Comma:                 ';',
// 		TrimLeadingSpace:      true,
// 	}
// 	for _, o := range opt {
// 		o(&options)
// 	}

// 	// Open the file
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "read file")
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.Comma = options.Comma
// 	reader.Comment = options.Comment
// 	reader.LazyQuotes = options.LazyQuotes
// 	reader.TrimLeadingSpace = options.TrimLeadingSpace
// 	var cols []string

// 	if options.HasHeaders {
// 		rec, err := reader.Read()
// 		if err != nil {
// 			return nil, errors.Wrap(err, "can't read headers")
// 		}
// 		cols = append(cols, rec...)
// 	}

// 	line := 0
// 	for {

// 	}

// 	return nil, nil
// }
