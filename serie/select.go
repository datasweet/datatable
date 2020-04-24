package serie

import "github.com/datasweet/datatable/value"

func (s *serie) Head(size int) Serie {
	return s.Subset(0, size)
}

func (s *serie) Tail(size int) Serie {
	return s.Subset(len(s.values)-size, size)
}

func (s *serie) Subset(from, size int) Serie {
	cpy := &serie{typ: s.typ, builder: s.builder}
	to := from + size
	if from >= 0 && size > 0 && to <= len(s.values) {
		cpy.values = s.values[from:to]
	}
	return cpy
}

func (s *serie) Pick(at ...int) Serie {
	cpy := &serie{typ: s.typ, builder: s.builder}
	for _, a := range at {
		if a < 0 || a >= len(s.values) {
			cpy.values = append(cpy.values, s.builder(nil))
			continue
		}
		cpy.values = append(cpy.values, s.values[a])
	}
	return cpy
}

func (s *serie) FindRows(where Predicate) []int {
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

func (s *serie) Filter(where Predicate) Serie {
	rows := s.FindRows(where)
	return s.Pick(rows...)
}

func (s *serie) Distinct() Serie {
	cpy := &serie{typ: s.typ, builder: s.builder}
	m := make(map[interface{}]bool)
	out := make([]value.Value, 0)
	for _, v := range s.values {
		if _, ok := m[v.Val()]; !ok {
			out = append(out, v)
			m[v.Val()] = true
		}
	}
	cpy.values = out
	return cpy
}
