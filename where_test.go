package datatable_test

import (
	"strings"
	"testing"
	"time"

	"github.com/datasweet/cast"
	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	// from join test
	customers, orders := sampleForJoin()
	dt, err := customers.LeftJoin(orders, datatable.On("[Customers].[id]", "[Orders].[user_id]"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)
	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
	)

	dt = dt.Where(func(row datatable.Row) bool {
		prenom, okp := cast.AsString(row["prenom"])
		num_facture, okf := cast.AsString(row["num_facture"])
		return okp && okf && (strings.ToLower(prenom) == "aimée" || num_facture == "A00106")
	})

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
	)
}
