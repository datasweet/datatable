package export

import "github.com/datasweet/datatable/datatable"

type MapOptions struct {
}

type MapOption func(opts *MapOptions)

// Map to export the datatable to a map
func Map(t datatable.Table, opt ...MapOption) []map[string]interface{} {
	if t == nil {
		return nil
	}
	// if t.Error() != nil {
	// 	return nil
	// }

	var hidden []string
	for _, col := range t.Columns() {
		if !col.IsVisible() {
			hidden = append(hidden, col.Name())
		}
	}

	rows := t.Rows()
	out := make([]map[string]interface{}, 0, len(rows))

	for _, r := range rows {
		for _, h := range hidden {
			delete(r, h)
		}
		out = append(out, r)
	}

	return out
}
