package datatable_test

import (
	"testing"
	"time"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	// from join test
	customers, orders := sampleForJoin()

	dt, err := datatable.LeftJoin([]*datatable.DataTable{customers, orders}, datatable.On("[Customers].[id]", "[Orders].[user_id]"))
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

	dt.Sort(datatable.By{Column: "num_facture", Desc: true})

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
	)

	dt.Sort(datatable.By{Column: "ville"}, datatable.By{Column: "id", Desc: true})

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
	)

	dt.Sort(datatable.By{Column: "ville"}, datatable.By{Column: "prix_total", Desc: true})

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
	)
}
