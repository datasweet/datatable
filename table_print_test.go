package datatable_test

import (
	"os"
	"testing"

	"github.com/datasweet/datatable"
)

func TestPrint(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", datatable.String, "Malzahar", "Xerath", "Teemo")
	tb.AddExprColumn("champion", datatable.String, "upper(`champ`)")
	tb.AddColumn("win", datatable.Int, 10, 20, 666)
	tb.AddColumn("loose", datatable.Int, 6, 5, 666)
	tb.AddExprColumn("winRate", datatable.String, "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", datatable.Float64, "sum(`win`)")
	tb.AddExprColumn("ok", datatable.Bool, "true")
	tb.AddExprColumn("hidden", datatable.Bool, "false")
	tb.HideColumn("hidden")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696.0, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696.0, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696.0, true,
	)

	tb.Print(os.Stdout, datatable.PrintColumnType(false), datatable.PrintMaxRows(3))
}
