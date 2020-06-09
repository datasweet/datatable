package datatable_test

import (
	"testing"
	"time"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

// Sample from https://sql.sh/cours/union
func sampleForConcat(t *testing.T) (*datatable.DataTable, *datatable.DataTable, *datatable.DataTable) {
	a := datatable.New("magasin1")
	a.AddStringColumn("prenom")
	a.AddStringColumn("nom")
	a.AddStringColumn("ville")
	a.AddTimeColumn("date_naissance")
	a.AddInt64Column("total_achat")

	a.AppendRow("Léon", "Dupuis", "Paris", "1983-03-06", 135)
	a.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	a.AppendRow("Sophie", "Dupond", "Marseille", "1986-02-22", 27)
	a.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

	b := datatable.New("magasin2")
	b.AddStringColumn("prenom")
	b.AddStringColumn("nom")
	b.AddStringColumn("ville")
	b.AddTimeColumn("date_naissance")
	b.AddInt64Column("total_achat")

	b.AppendRow("Marion", "Leroy", "Lyon", "1982-10-27", 285)
	b.AppendRow("Paul", "Moreau", "Lyon", "1976-04-19", 133)
	b.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	b.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

	c := datatable.New("magasin3")
	c.AddStringColumn("prenom")
	c.AddStringColumn("nom")
	c.AddStringColumn("ville")
	c.AddTimeColumn("date_naissance")
	c.AddFloat64Column("marge")

	c.AppendRow("Marion", "Leroy", "Lyon", "1982-10-27", 5.2)
	c.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 0.8)

	return a, b, c
}

func TestSimpleConcat(t *testing.T) {
	a, b, _ := sampleForConcat(t)
	dt, err := a.Concat(b)
	assert.NoError(t, err)
	assert.Equal(t, "magasin1", dt.Name())
	assert.Equal(t, 8, dt.NumRows())

	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), int64(135),
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75),
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), int64(27),
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39),
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), int64(285),
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), int64(133),
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75),
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39),
	)
}

func TestGrowColConcat(t *testing.T) {
	a, b, c := sampleForConcat(t)
	dt, err := a.Concat(b, c)
	assert.NoError(t, err)
	assert.Equal(t, "magasin1", dt.Name())
	assert.Equal(t, 10, dt.NumRows())

	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat", "marge",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), int64(135), nil,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75), nil,
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), int64(27), nil,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39), nil,
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), int64(285), nil,
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), int64(133), nil,
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75), nil,
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39), nil,
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), nil, float64(5.2),
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), nil, float64(0.8),
	)
}

func TestConcatWithExpr(t *testing.T) {
	a, b, _ := sampleForConcat(t)
	a.AddStringExprColumn("upper_ville", "UPPER(ville)")
	b.AddStringExprColumn("upper_ville", "UPPER(ville)")

	dt, err := a.Concat(b)
	assert.NoError(t, err)
	assert.Equal(t, "magasin1", dt.Name())
	assert.Equal(t, 8, dt.NumRows())

	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat", "upper_ville",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), int64(135), "PARIS",
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75), "PARIS",
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), int64(27), "MARSEILLE",
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39), "PARIS",
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), int64(285), "LYON",
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), int64(133), "LYON",
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75), "PARIS",
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39), "PARIS",
	)
}
