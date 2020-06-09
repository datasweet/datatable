package datatable

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

var hasher = &hasherImpl{}

type hasherImpl struct{}

func (h *hasherImpl) Row(row Row, cols []string) uint64 {
	if row == nil {
		return 0
	}
	hash := fnv.New64()
	buff := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buff)

	for _, name := range cols {
		enc.Encode(row.Get(name))
		hash.Write(buff.Bytes())
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
