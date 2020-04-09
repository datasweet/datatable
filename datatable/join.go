package datatable

// // InnerJoin selects records that have matching values in both tables.
// // left datatable is used as reference datatable.
// // <!> InnerJoin transforms an expr column to a raw column
// func InnerJoin(tables []DataTable, on []*JoinOn) (DataTable, error) {
// 	return newJoinImpl(innerJoin, tables, on).Compute()
// }

// // LeftJoin returns all records from the left table (table1), and the matched records from the right table (table2).
// // The result is NULL from the right side, if there is no match.
// // <!> LeftJoin transforms an expr column to a raw column
// func LeftJoin(tables []DataTable, on []*JoinOn) (DataTable, error) {
// 	return newJoinImpl(leftJoin, tables, on).Compute()
// }

// // RightJoin returns all records from the right table (table2), and the matched records from the left table (table1).
// // The result is NULL from the left side, when there is no match.
// // <!> RightJoin transforms an expr column to a raw column
// func RightJoin(tables []DataTable, on []*JoinOn) (DataTable, error) {
// 	return newJoinImpl(rightJoin, tables, on).Compute()
// }

// // OuterJoin returns all records when there is a match in either left or right table
// // <!> OuterJoin transforms an expr column to a raw column
// func OuterJoin(tables []DataTable, on []*JoinOn) (DataTable, error) {
// 	return newJoinImpl(outerJoin, tables, on).Compute()
// }

// type JoinOn struct {
// 	Table string
// 	Field string
// }

// var rgOn = regexp.MustCompile(`^(?:\[([^]]+)\]\.)?(?:\[([^]]+)\])$`)

// func On(fields ...string) []*JoinOn {
// 	var jon []*JoinOn
// 	for _, f := range fields {
// 		matches := rgOn.FindStringSubmatch(f)

// 		switch len(matches) {
// 		case 0:
// 			jon = append(jon, &JoinOn{Table: "*", Field: f})
// 		case 3:
// 			t := matches[1]
// 			if len(t) == 0 {
// 				t = "*"
// 			}
// 			jon = append(jon, &JoinOn{Table: t, Field: matches[2]})
// 		default:
// 			return nil
// 		}
// 	}

// 	return jon
// }

// func Using(fields ...string) []*JoinOn {
// 	var jon []*JoinOn
// 	for _, f := range fields {
// 		jon = append(jon, &JoinOn{Table: "*", Field: f})
// 	}
// 	return jon
// }

// type joinType uint8

// const (
// 	innerJoin joinType = iota
// 	leftJoin
// 	rightJoin
// 	outerJoin
// )

// func colname(dt DataTable, col string) string {
// 	var sb strings.Builder
// 	sb.WriteString(dt.Name())
// 	sb.WriteString(".")
// 	sb.WriteString(col)
// 	return sb.String()
// }

// type joinClause struct {
// 	table         DataTable
// 	mcols         map[string][]string
// 	on            []string
// 	includeOnCols bool
// 	cmapper       [][2]string // [initial, output]
// 	hashtable     map[uint64][]int
// 	consumed      map[int]bool
// }

// func (jc *joinClause) copyColumnsTo(out DataTable) error {
// 	if out == nil {
// 		return errors.New("nil output datatable")
// 	}

// 	mon := make(map[string]bool, len(jc.on))
// 	for _, o := range jc.on {
// 		mon[o] = true
// 	}

// 	for _, name := range jc.table.Columns() {
// 		cname := name

// 		if _, found := mon[name]; found {
// 			if !jc.includeOnCols {
// 				continue
// 			}
// 		} else if v, ok := jc.mcols[name]; ok && len(v) > 1 {
// 			// commons col between table
// 			for _, tn := range v {
// 				if tn == jc.table.Name() {
// 					cname = colname(jc.table, name)
// 					break
// 				}
// 			}
// 		}

// 		col := jc.table.Column(name)
// 		ctyp := col.Type()
// 		if col.IsComputed() {
// 			ctyp = Raw
// 		}
// 		dc, err := out.AddColumn(cname, ctyp)
// 		if err != nil {
// 			return err
// 		}
// 		if label := col.Label(); len(label) > 0 {
// 			dc.Label(label)
// 		}
// 		dc.Hidden(col.Hidden())
// 		jc.cmapper = append(jc.cmapper, [2]string{name, cname})
// 	}

// 	return nil
// }

// func (jc *joinClause) initHashTable() {
// 	jc.hashtable = hasher.Table(jc.table, jc.on)
// 	jc.consumed = make(map[int]bool, jc.table.NumRows())
// }

// type joinImpl struct {
// 	mode    joinType
// 	tables  []DataTable
// 	on      []*JoinOn
// 	clauses []*joinClause
// 	mcols   map[string][]string
// }

// func newJoinImpl(mode joinType, tables []DataTable, on []*JoinOn) *joinImpl {
// 	return &joinImpl{
// 		mode:   mode,
// 		tables: tables,
// 		on:     on,
// 	}
// }

// func (j *joinImpl) Compute() (DataTable, error) {
// 	if err := j.checkInput(); err != nil {
// 		return nil, err
// 	}

// 	j.initColMapper()

// 	out := j.tables[0]
// 	for i := 1; i < len(j.tables); i++ {
// 		jdt, err := j.join(out, j.tables[i])
// 		if err != nil {
// 			return nil, err
// 		}
// 		out = jdt
// 	}

// 	if out == nil {
// 		return nil, errors.New("no output")
// 	}

// 	return out, nil
// }

// func (j *joinImpl) checkInput() error {
// 	if len(j.tables) < 2 {
// 		return errors.New("we need at least 2 datatables to compute a join")
// 	}
// 	for i, t := range j.tables {
// 		if t == nil || len(t.Name()) == 0 || t.NumCols() == 0 {
// 			return errors.Errorf("table #%d is nil", i)
// 		}
// 	}
// 	if len(j.on) == 0 {
// 		return errors.New("no on clauses")
// 	}
// 	for i, o := range j.on {
// 		if o == nil || len(o.Field) == 0 {
// 			return errors.Errorf("on #%d is nil", i)
// 		}
// 	}
// 	return nil
// }

// func (j *joinImpl) initColMapper() {
// 	mcols := make(map[string][]string)
// 	for _, t := range j.tables {
// 		for _, name := range t.Columns() {
// 			mcols[name] = append(mcols[name], t.Name())
// 		}
// 	}
// 	j.mcols = mcols
// }

// func (j *joinImpl) join(left, right DataTable) (DataTable, error) {
// 	if left == nil {
// 		return nil, errors.New("left is nil datatable")
// 	}
// 	if right == nil {
// 		return nil, errors.New("right is nil datatable")
// 	}

// 	clauses := [2]*joinClause{
// 		&joinClause{
// 			table:         left,
// 			mcols:         j.mcols,
// 			includeOnCols: true,
// 		},
// 		&joinClause{
// 			table: right,
// 			mcols: j.mcols,
// 		},
// 	}

// 	// find on clauses
// 	for _, o := range j.on {
// 		if o.Table == left.Name() {
// 			clauses[0].on = append(clauses[0].on, o.Field)
// 			continue
// 		}

// 		if o.Table == right.Name() {
// 			clauses[1].on = append(clauses[1].on, o.Field)
// 			continue
// 		}

// 		if o.Table == "*" || len(o.Table) == 0 {
// 			clauses[0].on = append(clauses[0].on, o.Field)
// 			clauses[1].on = append(clauses[1].on, o.Field)
// 		}
// 	}

// 	// create output
// 	out := New(left.Name())
// 	for _, clause := range clauses {
// 		if err := clause.copyColumnsTo(out); err != nil {
// 			return nil, err
// 		}
// 	}

// 	// mode
// 	var ref, join *joinClause
// 	switch j.mode {
// 	case innerJoin, leftJoin, outerJoin:
// 		ref, join = clauses[0], clauses[1]
// 	case rightJoin:
// 		ref, join = clauses[1], clauses[0]
// 	default:
// 		return nil, errors.Errorf("unknown mode '%v'", j.mode)
// 	}

// 	join.initHashTable()

// 	// Copy rows
// 	for _, refrow := range ref.table.Rows() {
// 		// Create hash
// 		hash := hasher.Row(refrow, ref.on)

// 		// Have we same hash in jointable ?
// 		if indexes, ok := join.hashtable[hash]; ok {
// 			for _, idx := range indexes {
// 				joinrow := join.table.Row(idx)
// 				row := out.NewRow()
// 				for _, cm := range ref.cmapper {
// 					row[cm[1]] = refrow.Get(cm[0])
// 				}
// 				for _, cm := range join.cmapper {
// 					row[cm[1]] = joinrow.Get(cm[0])
// 				}
// 				join.consumed[idx] = true
// 				out.AppendRow(row)
// 			}
// 		} else if j.mode != innerJoin {
// 			row := make(Row, len(refrow))
// 			for _, cm := range ref.cmapper {
// 				row[cm[1]] = refrow.Get(cm[0])
// 			}
// 			out.AppendRow(row)
// 		}
// 	}

// 	// Outer: we must copy rows not consummed in right (join) table
// 	if j.mode == outerJoin {
// 		for i, joinrow := range join.table.Rows() {
// 			if b, ok := join.consumed[i]; ok && b {
// 				continue
// 			}
// 			row := make(Row, len(joinrow))
// 			for _, cm := range join.cmapper {
// 				row[cm[1]] = joinrow.Get(cm[0])
// 			}
// 			out.Append(row)
// 		}
// 	}

// 	return out, nil
// }
