package datatable

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// Row contains a row relative to columns
type Row map[string]interface{}

// Set cell
func (r Row) Set(k string, v interface{}) Row {
	// Check colName exists
	if _, ok := r[k]; ok {
		r[k] = v
	}
	return r
}

// Get cell
func (r Row) Get(k string) interface{} {
	// Check colName exists
	if v, ok := r[k]; ok {
		return v
	}
	return nil
}

// Hash computes the hash code from this datarow
// can be used to filter the table (distinct rows)
func (r Row) Hash() (uint64, bool) {
	var hash uint64

	// we xor-ing all keys / values to determinate
	// the same hash of a datarow despite the fact the map has not
	// been initalized in the same way
	for k, v := range r {
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
