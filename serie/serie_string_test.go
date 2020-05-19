package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
)

func TestString(t *testing.T) {
	s := serie.String("A00103", 1, 3.14, true, nil)
	assertSerieEq(t, s, "A00103", "1", "3.14", "true", "")
}
