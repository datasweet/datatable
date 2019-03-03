package datatable

import (
	"testing"
)

func TestFormula(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String, "Malzahar", "Xerath", "Teemo")
	tb.AddColumn("win", Number, 10, 20, 666)
	tb.AddColumn("loose", Number, 6, 5, 666)
	tb.AddFormulaColumn("winRate", Number, "[win] * 100 / ([win] + [loose])")

	checkTable(t, tb,
		"champ", "win", "loose", "winRate",
		"Malzahar", 10.0, 6.0, 62.5,
		"Xerath", 20.0, 5.0, 80.0,
		"Teemo", 666.0, 666.0, 50.0,
	)
}
