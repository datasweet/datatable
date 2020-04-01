package serie

import "github.com/datasweet/datatable/value"

func (s *serie) Head(size int) Serie {
	return s.Subset(0, size)
}

func (s *serie) Tail(size int) Serie {
	return s.Subset(len(s.values)-size, size)
}

func (s *serie) Subset(from, size int) Serie {
	cpy := s.clone(false)
	if cpy.err != nil {
		return cpy
	}
	to := from + size
	if from >= 0 && size > 0 && to <= len(s.values) {
		cpy.values = make([]value.Value, size)
		for i, val := range s.values[from:to] {
			cpy.values[i] = val.Clone()
		}
	}
	return cpy
}

func (s *serie) Pick(at ...int) Serie {
	ns := s.clone(false)
	if s.err != nil {
		return ns
	}
	nilValue := s.builder(nil)
	for _, a := range at {
		if a < 0 || a >= len(s.values) {
			ns.values = append(ns.values, nilValue)
			continue
		}
		val := s.values[a]
		ns.values = append(ns.values, val.Clone())
	}
	return ns
}

func (s *serie) FindRows(where value.Predicate) []int {
	if s.err != nil {
		return nil
	}
	if where == nil {
		return nil
	}
	var indexes []int
	for i, val := range s.values {
		// TODO: better search
		if where(val) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (s *serie) Filter(where value.Predicate) Serie {
	rows := s.FindRows(where)
	return s.Pick(rows...)
}

func (s *serie) Distinct() Serie {
	cpy := s.clone(false)
	if cpy.err != nil {
		return cpy
	}
	m := make(map[interface{}]bool)
	out := make([]value.Value, 0)
	for _, v := range s.values {
		if _, ok := m[v.Val()]; !ok {
			out = append(out, v.Clone())
			m[v.Val()] = true
		}
	}

	cpy.values = out
	return cpy
}
