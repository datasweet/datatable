package datatable

// Swap swap 2 columns
func (t *DataTable) Swap(colA, colB string) bool {
	a := t.findColIndex(colA)
	if a < 0 {
		return false
	}

	b := t.findColIndex(colB)
	if b < 0 {
		return false
	}

	tmp := t.cols[a]
	t.cols[a] = t.cols[b]
	t.cols[b] = tmp
	return true
}
