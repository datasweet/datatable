package export

// type SliceOptions struct {
// 	Header bool
// }

// type SliceOption func(opts *SliceOptions)

// func Slice(t datatable.Table, opt ...SliceOption) [][]interface{} {
// 	if t == nil {
// 		return nil
// 	}

// 	options := SliceOptions{
// 		Header: true
// 	}

// 	for _, o := range opt {
// 		o(&options)
// 	}

// 	var hidden []string
// 	for _, col := range t.Columns() {
// 		if !col.IsVisible() {
// 			hidden = append(hidden, col.Name())
// 		}
// 	}

// 	rows := t.Rows()
// 	out := make([][]interface{}, 0, len(rows))

// 	for _, r := range rows {
// 		for _, h := range hidden {
// 			delete(r, h)
// 		}
// 		out = append(out, r)
// 	}

// 	return out

// 	in := t.Rows()
// 	var out [][]interface{}
// 	var hr []interface{}
// 	if headers {
// 		for _, col := range cols {
// 			if !col.Hidden() {
// 				label := col.Label()
// 				if len(label) == 0 {
// 					label = col.Name()
// 				}
// 				hr = append(hr, label)
// 			}
// 		}
// 	}

// }
