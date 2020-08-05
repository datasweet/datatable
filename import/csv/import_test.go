package csv_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/datasweet/datatable"

	"github.com/datasweet/datatable/import/csv"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	dt, err := csv.Import("csv", "../../test/phone_data.csv",
		csv.HasHeader(true),
		csv.AcceptDate("02/01/06 15:04"),
		csv.AcceptDate("2006-01"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	dt.Print(os.Stdout, datatable.PrintMaxRows(24))

	dtc, err := dt.Aggregate(datatable.AggregateBy{datatable.Count, "index"})
	assert.NoError(t, err)
	fmt.Println(dtc)

	groups, err := dt.GroupBy(datatable.GroupBy{
		Name: "year",
		Type: datatable.Int,
		Keyer: func(row datatable.Row) (interface{}, bool) {
			if d, ok := row["date"]; ok {
				if tm, ok := d.(time.Time); ok {
					return tm.Year(), true
				}
			}
			return nil, false
		},
	})
	assert.NoError(t, err)
	out, err := groups.Aggregate(datatable.AggregateBy{datatable.Sum, "duration"}, datatable.AggregateBy{datatable.CountDistinct, "network"})
	assert.NoError(t, err)
	fmt.Println(out)

}
