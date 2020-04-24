package datatable

import (
	"fmt"
	"testing"
)

func dummy(s Serie) {
	s.Append(5, 6, 7, 8, 9)
}

func TestRefOrStruct(t *testing.T) {
	s := Serie{Name: "toto"}
	s.Append(1, 2, 3, 4, 5, 6)

	fmt.Println(s.Values)

	dummy(s)
	fmt.Println(s.Values)
}
