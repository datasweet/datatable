package serie

import "github.com/pkg/errors"

// Append values to the serie.
// TODO: Append([]int) must flatten
func (s *serie) Append(v ...interface{}) {
	if len(v) == 0 {
		s.values = append(s.values, s.builder())
		return
	}

	for _, val := range v {
		s.values = append(s.values, s.builder(val))
	}
}

// Prepend values to the serie
func (s *serie) Prepend(v ...interface{}) {
	s.Insert(0, v...)
}

// Insert values to the serie at index
func (s *serie) Insert(at int, v ...interface{}) error {
	if at < 0 || at >= len(s.values) {
		return errors.Errorf("insert at [%d]: index out of range with length %d", at, len(s.values))
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

// Update a value at index
func (s *serie) Update(at int, v interface{}) error {
	if at < 0 || at >= len(s.values) {
		return errors.Errorf("update at [%d]: index out of range with length %d", at, len(s.values))
	}
	s.values[at].Set(v)
	return nil
}

// Delete a value at index
func (s *serie) Delete(at int) error {
	if at < 0 || at >= len(s.values) {
		return errors.Errorf("delete at [%d]: index out of range with length %d", at, len(s.values))
	}
	if at < len(s.values)-1 {
		copy(s.values[at:], s.values[at+1:])
	}
	s.values[len(s.values)-1] = nil
	s.values = s.values[:len(s.values)-1]
	return nil
}

// Grow the serie with size
// Grow will create zero value
func (s *serie) Grow(size int) error {
	if size < 0 {
		return errors.Errorf("grow: size '%d' must be > 0", size)
	}
	for i := 0; i < size; i++ {
		s.values = append(s.values, s.builder())
	}
	return nil
}

// Shrink the serie with size
func (s *serie) Shrink(size int) error {
	if size < 0 {
		return errors.Errorf("shrink: size '%d' must be > 0", size)
	}
	len := len(s.values)
	if size > len {
		return errors.Errorf("shrink: size '%d' must be < length '%d'", size, len)
	}
	from := len - size
	to := len
	for i := size; i < to; i++ {
		s.values[i] = nil
	}
	s.values = s.values[:from]
	return nil
}

// Concat the serie (mutate) with others series
// series provided must be the same type as the source serie
func (s *serie) Concat(series ...Serie) error {
	if len(series) == 0 {
		return nil
	}

	// check type
	// for _, os := range series {
	// 	if os.Type() != s.Type() {
	// 		return errors.Errorf("can't concat serie of type '%s' to '%s'", os.Type())
	// 	}
	// }

	return nil
}
