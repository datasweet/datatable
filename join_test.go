package datatable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://sql.sh/cours/jointures/inner-join
func sampleForJoin() (DataTable, DataTable) {
	customers := New("Customers")
	customers.AddColumn("id", Int)
	customers.AddColumn("prenom", String)
	customers.AddColumn("nom", String)
	customers.AddColumn("email", String)
	customers.AddColumn("ville", String)
	customers.AppendRow(1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris")
	customers.AppendRow(2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon")
	customers.AppendRow(3, "Marine", "Prevost", "m.prevost@example.com", "Lille")
	customers.AppendRow(4, "Luc", "Rolland", "lucrolland@example.com", "Marseille")

	orders := New("Orders")
	orders.AddColumn("user_id", Int)
	orders.AddColumn("date_achat", Datetime)
	orders.AddColumn("num_facture", String)
	orders.AddColumn("prix_total", Float)
	orders.AppendRow(1, "2013-01-23", "A00103", 203.14)
	orders.AppendRow(1, "2013-02-14", "A00104", 124.00)
	orders.AppendRow(2, "2013-02-17", "A00105", 149.45)
	orders.AppendRow(3, "2013-02-21", "A00106", 235.35)
	orders.AppendRow(5, "2013-03-02", "A00107", 47.58)

	return customers, orders
}

func TestInnerJoin(t *testing.T) {
	customers, orders := sampleForJoin()

	dt, err := InnerJoin(customers, orders, On("id", "user_id"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		int64(2), "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		int64(3), "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
	)
}

func TestLeftJoin(t *testing.T) {
	customers, orders := sampleForJoin()

	dt, err := LeftJoin(customers, orders, On("id", "user_id"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		int64(2), "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		int64(3), "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		int64(4), "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
	)
}

func TestRightJoin(t *testing.T) {
	customers, orders := sampleForJoin()

	dt, err := RightJoin(customers, orders, On("id", "user_id"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		int64(2), "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		int64(3), "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
		nil, nil, nil, nil, nil, time.Date(2013, time.March, 2, 0, 0, 0, 0, time.UTC), "A00107", 47.58,
	)
}

func TestJoinWithExpr(t *testing.T) {
	customers, orders := sampleForJoin()
	customers.AddExprColumn("upper_ville", "UPPER(ville)")

	dt, err := InnerJoin(customers, orders, On("id", "user_id"))
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	idx, dc := dt.GetColumn("upper_ville")
	assert.Equal(t, 5, idx)
	assert.Equal(t, Raw, dc.Type())
	assert.False(t, dc.Computed())
	checkTable(t, dt,
		"id", "prenom", "nom", "email", "ville", "upper_ville", "date_achat", "num_facture", "prix_total",
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "PARIS", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
		int64(1), "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "PARIS", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
		int64(2), "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", "LYON", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
		int64(3), "Marine", "Prevost", "m.prevost@example.com", "Lille", "LILLE", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
	)
}
