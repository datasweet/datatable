package datatable_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func TestAggregate(t *testing.T) {
	customers, orders := sampleForJoin()

	dt, err := customers.LeftJoin(orders, datatable.On("[Customers].[id]", "[Orders].[user_id]"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)
	fmt.Println(dt)

	// Aggregate by SUM
	out, err := dt.Aggregate(datatable.AggregateBy{datatable.Sum, "prix_total"})
	assert.NoError(t, err)
	assert.NotNil(t, out)
	fmt.Println(out)

	// Aggregate by SUM(prix_total), COUNT_DISTINCT(ville)
	out, err = dt.Aggregate(datatable.AggregateBy{datatable.Sum, "prix_total"}, datatable.AggregateBy{datatable.CountDistinct, "ville"})
	assert.NoError(t, err)
	assert.NotNil(t, out)
	fmt.Println(out)

	groups, err := dt.GroupBy(
		datatable.GroupBy{
			Name: "Year",
			Type: datatable.Int64,
			Keyer: func(row datatable.Row) (interface{}, bool) {
				t, ok := row["date_achat"].(time.Time)
				if !ok {
					return 0, false
				}
				return t.Year(), true
			},
		},
		datatable.GroupBy{
			Name: "Month",
			Type: datatable.Int,
			Keyer: func(row datatable.Row) (interface{}, bool) {
				t, ok := row["date_achat"].(time.Time)
				if !ok {
					return 0, false
				}
				return int(t.Month()), true
			},
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, groups)

	gdt, err := groups.Aggregate(datatable.AggregateBy{datatable.Sum, "prix_total"})
	assert.NoError(t, err)
	fmt.Println(gdt)
}
