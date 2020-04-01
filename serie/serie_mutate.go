package serie

import "github.com/pkg/errors"

// Append values to the serie.
// TODO: Append([]int) must flatten
func (s *serie) Append(v ...interface{}) {
	if s.err != nil {
		return
	}
	for _, val := range v {
		s.values = append(s.values, s.builder(val))
	}
}

func (s *serie) Prepend(v ...interface{}) {
	s.Insert(0, v...)
}

func (s *serie) Insert(at int, v ...interface{}) {
	if s.err != nil {
		return
	}

	if at < 0 || at >= len(s.values) {
		s.err = errors.Errorf("insert at [%d]: index out of range with length %d", at, len(s.values))
		return
	}

	if len(v) == 0 {
		return
	}

	for i := 0; i < len(v); i++ {
		s.values = append(s.values, nil)
	}

	copy(s.values[at+len(v):], s.values[at:])

	for i, val := range v {
		s.values[i+at] = s.builder(val)
	}
}

func (s *serie) Update(at int, v interface{}) {
	if s.err != nil {
		return
	}
	if at < 0 || at >= len(s.values) {
		s.err = errors.Errorf("update at [%d]: index out of range with length %d", at, len(s.values))
		return
	}
	s.values[at].Set(v)
}

func (s *serie) Delete(at int) {
	if s.err != nil {
		return
	}
	if at < 0 || at >= len(s.values) {
		s.err = errors.Errorf("delete at [%d]: index out of range with length %d", at, len(s.values))
		return
	}
	if at < len(s.values)-1 {
		copy(s.values[at:], s.values[at+1:])
	}
	s.values[len(s.values)-1] = nil
	s.values = s.values[:len(s.values)-1]
}
