package datatable_test

import (
	"os"
	"testing"

	"github.com/datasweet/datatable"
)

func TestPrint(t *testing.T) {
	tb := New(t)

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696.0, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696.0, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696.0, true,
	)

	tb.Print(os.Stdout, datatable.PrintColumnType(false), datatable.PrintMaxRows(3))
}
