package datatable

import (
	"bytes"
	"encoding/gob"

	"github.com/cespare/xxhash"
)

var hasher = &hasherImpl{}

type hasherImpl struct{}

func (h *hasherImpl) Row(row Row, cols []string) uint64 {
	if row == nil {
		return 0
	}
	hash := xxhash.New()
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)

	for _, name := range cols {
		enc.Encode(row[name])
	}

	return hash.Sum64()
}

func (h *hasherImpl) Table(dt *DataTable, cols []string) map[uint64][]int {
	if dt == nil {
		return nil
	}
	mh := make(map[uint64][]int, 0)
	for i, row := range dt.Rows() {
		hash := h.Row(row, cols)
		mh[hash] = append(mh[hash], i)
	}
	return mh
}
