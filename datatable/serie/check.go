package serie

import (
	"reflect"

	"github.com/pkg/errors"
)

func (s *serie) mustBeSlice() (reflect.Value, error) {
	rv := reflect.Indirect(reflect.ValueOf(s.values))
	if rv.Kind() != reflect.Slice {
		return reflect.Value{}, errors.Errorf("values must be a slice")
	}
	return rv, nil
}
