package serie

import "errors"

func (s *serie) Append(v ...interface{}) {
	for _, val := range v {
		s.values = append(s.values, s.builder(val))
	}
}

func (s *serie) Prepend(v ...interface{}) {
	s.Insert(1, v...)
}

func (s *serie) Insert(at int, v ...interface{}) error {
	at = at - 1

	if at < 0 || at >= len(s.values) {
		return errors.New("slice bounds out of range")
	}

	if len(v) == 0 {
		return nil
	}

	for i := 0; i < len(v); i++ {
		s.values = append(s.values, nil)
	}

	copy(s.values[at+len(v):], s.values[at:])

	for i, val := range v {
		s.values[i+at] = s.builder(val)
	}

	return nil
}
