package serie

func (s *serie) Head(size int) Serie {
	return s.Subset(1, size)
}

func (s *serie) Tail(size int) Serie {
	return s.Subset(len(s.values)-size+1, size)
}

func (s *serie) Subset(at, size int) Serie {
	ns := s.clone(false)
	from := at - 1
	to := from + size
	if from >= 0 && size > 0 && to <= len(s.values) {
		ns.values = make([]Value, size)
		for i, val := range s.values[from:to] {
			ns.values[i] = val.Clone()
		}
	}
	return ns
}

func (s *serie) Pick(at ...int) Serie {
	ns := s.clone(false)
	nilValue := NewNilValue(s.typ)
	for _, a := range at {
		a = a - 1
		if a < 0 || a >= len(s.values) {
			ns.values = append(ns.values, nilValue)
			continue
		}
		val := s.values[a]
		ns.values = append(ns.values, val.Clone())
	}
	return ns
}

func (s *serie) FindRows(where ValuePredicate) []int {
	if where == nil {
		return nil
	}
	var indexes []int
	for i, val := range s.values {
		// TODO: better search
		if where(val) {
			indexes = append(indexes, i+1)
		}
	}
	return indexes
}

func (s *serie) Filter(where ValuePredicate) Serie {
	rows := s.FindRows(where)
	return s.Pick(rows...)
}
