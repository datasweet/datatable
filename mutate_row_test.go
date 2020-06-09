package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable"
	"github.com/datasweet/datatable/serie"
)

func TestSwapRow(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", serie.String("Malzahar", "Xerath", "Teemo"))
	tb.AddExprColumn("champion", serie.String(), "upper(`champ`)")
	tb.AddColumn("win", serie.Int(10, 20, 666))
	tb.AddColumn("loose", serie.Int(6, 5, 666))
	tb.AddExprColumn("winRate", serie.Float64(), "(`win` * 100 / (`win` + `loose`))")
	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate",
		"Malzahar", "MALZAHAR", 10, 6, 62.5,
		"Xerath", "XERATH", 20, 5, 80.0,
		"Teemo", "TEEMO", 666, 666, 50.0,
	)

	tb.SwapRow(0, 2)

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate",
		"Teemo", "TEEMO", 666, 666, 50.0,
		"Xerath", "XERATH", 20, 5, 80.0,
		"Malzahar", "MALZAHAR", 10, 6, 62.5,
	)
}
