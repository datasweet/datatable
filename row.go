package datatable

import (
	"bytes"
	"encoding/gob"

	"github.com/cespare/xxhash"
)

// Row contains a row relative to columns
type Row map[string]interface{}

// Set cell
func (r Row) Set(k string, v interface{}) Row {
	r[k] = v
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
func (r Row) Hash() uint64 {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	for _, v := range r {
		enc.Encode(v)
	}
	return xxhash.Sum64(buff.Bytes())
}
