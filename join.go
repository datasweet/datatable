package datatable

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/pkg/errors"
)

// ColumnSelector to define a column selector between a left and right datatable
type ColumnSelector struct {
	Left  string
	Right string
}

func On(left, right string) ColumnSelector {
	return ColumnSelector{Left: left, Right: right}
}

func Using(colName string) ColumnSelector {
	return ColumnSelector{Left: colName, Right: colName}
}

// InnerJoin selects records that have matching values in both tables.
// left datatable is used as reference datatable.
// <!> InnerJoin transforms an expr column to a raw column
func InnerJoin(left, right DataTable, on ...ColumnSelector) (DataTable, error) {
	return join(left, right, innerJoin, on...)
}

// LeftJoin returns all records from the left table (table1), and the matched records from the right table (table2).
// The result is NULL from the right side, if there is no match.
// <!> LeftJoin transforms an expr column to a raw column
func LeftJoin(left, right DataTable, on ...ColumnSelector) (DataTable, error) {
	return join(left, right, leftJoin, on...)
}

// RightJoin returns all records from the right table (table2), and the matched records from the left table (table1).
// The result is NULL from the left side, when there is no match.
// <!> RightJoin transforms an expr column to a raw column
func RightJoin(left, right DataTable, on ...ColumnSelector) (DataTable, error) {
	return join(left, right, rightJoin, on...)
}

func computeHash(row DataRow, cols ...string) uint64 {
	hash := fnv.New64()
	buff := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buff)

	for _, name := range cols {
		enc.Encode(row.Get(name))
		hash.Write(buff.Bytes())
	}

	return hash.Sum64()
}

func computeHashTable(dt DataTable, cols ...string) map[uint64][]int {
	mh := make(map[uint64][]int, 0)

	// 2- Create hash for each row
	for i, row := range dt.Rows() {
		hash := computeHash(row, cols...)
		mh[hash] = append(mh[hash], i)
	}

	return mh
}

type joinType uint8

const (
	innerJoin joinType = iota
	leftJoin
	rightJoin
)

type joinMeta struct {
	Table   DataTable
	Clauses []string
	Cols    []string
}

func colname(dt DataTable, col string) string {
	var sb strings.Builder
	sb.WriteString(dt.Name())
	sb.WriteString(".")
	sb.WriteString(col)
	return sb.String()
}

func join(left, right DataTable, mode joinType, on ...ColumnSelector) (DataTable, error) {
	if left == nil {
		return nil, errors.New("left is nil datatable")
	}
	if right == nil {
		return nil, errors.New("right is nil datatable")
	}

	if len(on) == 0 {
		return nil, errors.New("no on clause")
	}

	// destructurate on clause
	var lclauses []string
	var rclauses []string
	for _, clause := range on {
		lclauses = append(lclauses, clause.Left)
		rclauses = append(rclauses, clause.Right)
	}

	dt := New(fmt.Sprintf("join-%s-%s", left.Name(), right.Name()))

	// Keep in memory the left / right columns names
	// to copy more easilier values
	var lcols []string
	var rcols []string

	// Create columns
	// Copy left columns : reference table
	for _, col := range left.Columns() {
		name := col.Name()
		ctyp := col.Type()
		if col.Computed() {
			ctyp = Raw
		}
		if _, err := dt.AddColumn(colname(left, name), ctyp); err != nil {
			return nil, err
		}
		lcols = append(lcols, name)
	}

	// Copy rights columns
	// <!> expr column can return "nil",
	// cause we can have an expr on a "id" right column
	// example:
	// InnerJoin(l, r, []string{"id"}, []string{"user_id"})
	// if we have on right datatable, an expr with LOWER("user_id") => bug
	for _, col := range right.Columns() {
		name := col.Name()
		found := false
		for _, clause := range on {
			if clause.Right == name {
				found = true
				break
			}
		}

		if found {
			continue
		}

		ctyp := col.Type()
		if col.Computed() {
			ctyp = Raw
		}
		if _, err := dt.AddColumn(colname(right, name), ctyp); err != nil {
			return nil, err
		}

		rcols = append(rcols, name)
	}

	var ref joinMeta
	var join joinMeta
	var inner bool

	switch mode {
	case innerJoin:
		ref = joinMeta{Table: left, Clauses: lclauses, Cols: lcols}
		join = joinMeta{Table: right, Clauses: rclauses, Cols: rcols}
		inner = true
	case leftJoin:
		ref = joinMeta{Table: left, Clauses: lclauses, Cols: lcols}
		join = joinMeta{Table: right, Clauses: rclauses, Cols: rcols}
		inner = false
	case rightJoin:
		ref = joinMeta{Table: right, Clauses: rclauses, Cols: rcols}
		join = joinMeta{Table: left, Clauses: lclauses, Cols: lcols}
		inner = false

	default:
		return nil, errors.Errorf("unknown mode '%v'", mode)
	}

	hashtable := computeHashTable(join.Table, join.Clauses...)

	// Copy rows
	for _, refrow := range ref.Table.Rows() {
		// Create hash
		hash := computeHash(refrow, ref.Clauses...)

		// Have we same hash in jointable ?
		if indexes, ok := hashtable[hash]; ok {
			for _, idx := range indexes {
				joinrow := join.Table.GetRow(idx)
				row := dt.NewRow()

				for _, name := range ref.Cols {
					row[colname(ref.Table, name)] = refrow.Get(name)
				}
				for _, name := range join.Cols {
					row[colname(join.Table, name)] = joinrow.Get(name)
				}

				dt.AddRow(row)
			}
		} else if !inner {
			row := make(DataRow, len(refrow))
			for k, v := range refrow {
				row[colname(ref.Table, k)] = v
				delete(row, k)
			}
			dt.AddRow(row)
		}
	}

	return dt, nil
}
