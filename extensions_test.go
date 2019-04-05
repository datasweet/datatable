package datatable

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://sql.sh/cours/jointures/inner-join
func sampleForJoin() (DataTable, DataTable) {
	customers := New("Customers")
	customers.AddColumn("id", Number)
	customers.AddColumn("prenom", String)
	customers.AddColumn("nom", String)
	customers.AddColumn("email", String)
	customers.AddColumn("ville", String)
	customers.AppendRow(1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris")
	customers.AppendRow(2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon")
	customers.AppendRow(3, "Marine", "Prevost", "m.prevost@example.com", "Lille")
	customers.AppendRow(4, "Luc", "Rolland", "lucrolland@example.com", "Marseille")

	orders := New("Orders")
	orders.AddColumn("user_id", Number)
	orders.AddColumn("date_achat", Datetime)
	orders.AddColumn("num_facture", String)
	orders.AddColumn("prix_total", Number)
	orders.AppendRow(1, "2013-01-23", "A00103", 203.14)
	orders.AppendRow(1, "2013-02-14", "A00104", 124.00)
	orders.AppendRow(2, "2013-02-17", "A00105", 149.45)
	orders.AppendRow(2, "2013-02-21", "A00106", 235.35)
	orders.AppendRow(5, "2013-03-02", "A00107", 47.58)

	return customers, orders
}

// Sample from https://sql.sh/cours/union
func sampleForUnion(t *testing.T) (DataTable, DataTable) {
	a := New("magasin1")
	a.AddColumn("prenom", String)
	a.AddColumn("nom", String)
	a.AddColumn("ville", String)
	a.AddColumn("date_naissance", Datetime)
	a.AddColumn("total_achat", Number)

	a.AppendRow("Léon", "Dupuis", "Paris", "1983-03-06", 135)
	a.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	a.AppendRow("Sophie", "Dupond", "Marseille", "1986-02-22", 27)
	a.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

	checkTable(t, a,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), 135.0,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), 75.0,
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), 27.0,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), 39.0,
	)

	b := New("magasin2")
	b.AddColumn("prenom", String)
	b.AddColumn("nom", String)
	b.AddColumn("ville", String)
	b.AddColumn("date_naissance", Datetime)
	b.AddColumn("total_achat", Number)

	b.AppendRow("Marion", "Leroy", "Lyon", "1982-10-27", 285)
	b.AppendRow("Paul", "Moreau", "Lyon", "1976-04-19", 133)
	b.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	b.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

	checkTable(t, b,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), 285.0,
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), 133.0,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), 75.0,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), 39.0,
	)

	return a, b
}

func TestUnionAll(t *testing.T) {
	a, b := sampleForUnion(t)
	dt, err := UnionAll(a, b)
	assert.NoError(t, err)
	assert.Equal(t, "unionall-magasin1-magasin2", dt.Name())
	assert.Equal(t, 8, dt.NumRows())
	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), 135.0,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), 75.0,
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), 27.0,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), 39.0,
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), 285.0,
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), 133.0,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), 75.0,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), 39.0,
	)
}

func TestUnion(t *testing.T) {
	a, b := sampleForUnion(t)
	dt, err := Union(a, b)
	assert.NoError(t, err)
	assert.Equal(t, "union-magasin1-magasin2", dt.Name())
	assert.Equal(t, 6, dt.NumRows())
	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), 135.0,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), 75.0,
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), 27.0,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), 39.0,
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), 285.0,
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), 133.0,
	)
}

func TestInnerJoin(t *testing.T) {
	customers, orders := sampleForJoin()
	dt := InnerJoin(customers, orders, func(l, r DataRow) bool {
		return l.Get("id") == r.Get("user_id")
	})

	assert.NotNil(t, dt)

	fmt.Println(ToTable(dt, true))

	// checkTable(t, dt,
	// 	"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
	// 	1.0, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "2013-01-23", "A00103", 203.14,
	// 	1.0, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "2013-02-14", "A00104", 124.00,
	// 	2.0, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", "2013-02-17", "A00105", 149.45,
	// 	2.0, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", "2013-02-21", "A00106", 235.35,
	// )
}
