package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/datasweet/datatable"
	"github.com/datasweet/datatable/import/csv"
)

func main() {
	dt, err := csv.Import("csv", "phone_data.csv",
		csv.HasHeader(true),
		csv.AcceptDate("02/01/06 15:04"),
		csv.AcceptDate("2006-01"),
	)
	if err != nil {
		log.Fatalf("reading csv: %v", err)
	}

	dt.Print(os.Stdout, datatable.PrintMaxRows(24))

	dt2, err := dt.Aggregate(datatable.AggregateBy{Type: datatable.Count, Field: "index"})
	if err != nil {
		log.Fatalf("aggregate COUNT('index'): %v", err)
	}
	fmt.Println(dt2)

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
	if err != nil {
		log.Fatalf("GROUP BY 'year': %v", err)
	}
	dt3, err := groups.Aggregate(
		datatable.AggregateBy{Type: datatable.Sum, Field: "duration"},
		datatable.AggregateBy{Type: datatable.CountDistinct, Field: "network"},
	)
	if err != nil {
		log.Fatalf("Aggregate SUM('duration'), COUNT_DISTINCT('network') GROUP BY 'year': %v", err)
	}
	fmt.Println(dt3)

}
