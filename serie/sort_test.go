package serie_test

import (
	"sort"
	"testing"

	"github.com/datasweet/datatable/serie"
)

func TestSortInt(t *testing.T) {
	random := []interface{}{
		31, 23, 98, 3, 59, 67, 5, 5, 87, 18,
		3, 88, 7, 63, 29, 62, 37, 66, 87, 26,
		24, 5, 62, 75, 69, 56, 15, 59, 40, 34,
		68, 32, 34, 29, 90, 21, 8, 8, 100, 64,
		30, 56, 73, 2, 65, 74, 3, 26, 92, 46,
		6, 100, 35, 17, 91, 55, 99, 87, 9, 25,
		55, 76, 39, 78, 43, 99, 35, 90, 36, 27,
		52, 65, 33, 49, 84, 87, 42, 92, 27, 65,
		48, 47, 74, 98, 76, 88, 18, 100, 69, 57,
		69, 90, 74, 25, 64, 37, 63, 61, 85, 12,
	}
	s := serie.Int(random...)

	sort.Slice(random, func(i, j int) bool {
		return random[i].(int) < random[j].(int)
	})
	s.SortAsc()
	assertSerieEq(t, s, random...)

	sort.Slice(random, func(i, j int) bool {
		return random[i].(int) > random[j].(int)
	})
	s.SortDesc()
	assertSerieEq(t, s, random...)
}

func TestSortString(t *testing.T) {
	s := serie.String("A00103", "A00105", "A00104", "A00106", "A00104", nil)
	s.SortAsc()
	assertSerieEq(t, s, "", "A00103", "A00104", "A00104", "A00105", "A00106")

	s.SortDesc()
	assertSerieEq(t, s, "A00106", "A00105", "A00104", "A00104", "A00103", "")
}
