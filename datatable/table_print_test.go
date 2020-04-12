package datatable_test

import (
	"os"
	"testing"

	"github.com/datasweet/datatable/datatable"
	"github.com/datasweet/datatable/serie"
)

func TestPrint(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", serie.String("Malzahar", "Xerath", "Teemo"))
	tb.AddExprColumn("champion", serie.String(), "upper(`champ`)")
	tb.AddColumn("win", serie.Int(10, 20, 666))
	tb.AddColumn("loose", serie.Int(6, 5, 666))
	tb.AddExprColumn("winRate", serie.String(), "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", serie.Int(), "sum(`win`)")
	tb.AddExprColumn("ok", serie.Bool(), "true")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696.0, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696.0, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696.0, true,
	)

	tb.Print(os.Stdout)

	// tb.Print(os.Stdout, datatable.PrintColumnType(false), datatable.PrintColumnName(false))
}
