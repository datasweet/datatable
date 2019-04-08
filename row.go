package datatable

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// DataRow contains a row relative to columns
type DataRow map[string]interface{}

// Set cell
func (dr DataRow) Set(k string, v interface{}) DataRow {
	// Check colName exists
	if _, ok := dr[k]; ok {
		dr[k] = v
	}
	return dr
}

// Get cell
func (dr DataRow) Get(k string) interface{} {
	// Check colName exists
	if v, ok := dr[k]; ok {
		return v
	}
	return nil
}

// Hash computes the hash code from this datarow
// can be used to filter the table (distinct rows)
func (dr DataRow) Hash() (uint64, bool) {
	var hash uint64

	// we xor-ing all keys / values to determinate
	// the same hash of a datarow despite the fact the map has not
	// been initalized in the same way
	for k, v := range dr {
		h := fnv.New64()
		buf := bytes.NewBufferString(k)
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(v); err != nil {
			return 0, false
		}
		if _, err := h.Write(buf.Bytes()); err != nil {
			return 0, false
		}

		hash ^= h.Sum64()
	}

	return hash, true
}
