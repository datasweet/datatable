package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable"
)

func TestSwapRow(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", datatable.String, datatable.Values("Malzahar", "Xerath", "Teemo"))
	tb.AddColumn("champion", datatable.String, datatable.Expr("upper(`champ`)"))
	tb.AddColumn("win", datatable.Int, datatable.Values(10, 20, 666))
	tb.AddColumn("loose", datatable.Int, datatable.Values(6, 5, 666))
	tb.AddColumn("winRate", datatable.Float64, datatable.Expr("(`win` * 100 / (`win` + `loose`))"))
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
