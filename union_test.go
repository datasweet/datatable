package datatable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Sample from https://sql.sh/cours/union
func sampleForUnion(t *testing.T) (DataTable, DataTable) {
	a := New("magasin1")
	a.AddColumn("prenom", String)
	a.AddColumn("nom", String)
	a.AddColumn("ville", String)
	a.AddColumn("date_naissance", Datetime)
	a.AddColumn("total_achat", Int)

	a.AppendRow("Léon", "Dupuis", "Paris", "1983-03-06", 135)
	a.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	a.AppendRow("Sophie", "Dupond", "Marseille", "1986-02-22", 27)
	a.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

	b := New("magasin2")
	b.AddColumn("prenom", String)
	b.AddColumn("nom", String)
	b.AddColumn("ville", String)
	b.AddColumn("date_naissance", Datetime)
	b.AddColumn("total_achat", Int)

	b.AppendRow("Marion", "Leroy", "Lyon", "1982-10-27", 285)
	b.AppendRow("Paul", "Moreau", "Lyon", "1976-04-19", 133)
	b.AppendRow("Marie", "Bernard", "Paris", "1993-07-03", 75)
	b.AppendRow("Marcel", "Martin", "Paris", "1976-11-24", 39)

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

func TestUnion(t *testing.T) {
	a, b := sampleForUnion(t)
	dt, err := Union(a, b)
	assert.NoError(t, err)
	assert.Equal(t, "union-magasin1-magasin2", dt.Name())
	assert.Equal(t, 6, dt.NumRows())
	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), int64(135),
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75),
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), int64(27),
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39),
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), int64(285),
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), int64(133),
	)
}

func TestUnionWithExpr(t *testing.T) {
	a, b := sampleForUnion(t)
	a.AddExprColumn("upper_ville", "UPPER(ville)")
	b.AddExprColumn("upper_ville", "UPPER(ville)")

	dt, err := Union(a, b)
	assert.NoError(t, err)
	assert.Equal(t, "union-magasin1-magasin2", dt.Name())
	assert.Equal(t, 6, dt.NumRows())

	idx, dc := dt.GetColumn("upper_ville")
	assert.Equal(t, 5, idx)
	assert.Equal(t, Raw, dc.Type())
	assert.False(t, dc.Computed())

	checkTable(t, dt,
		"prenom", "nom", "ville", "date_naissance", "total_achat", "upper_ville",
		"Léon", "Dupuis", "Paris", time.Date(1983, time.March, 6, 0, 0, 0, 0, time.UTC), int64(135), "PARIS",
		"Marie", "Bernard", "Paris", time.Date(1993, time.July, 3, 0, 0, 0, 0, time.UTC), int64(75), "PARIS",
		"Sophie", "Dupond", "Marseille", time.Date(1986, time.February, 22, 0, 0, 0, 0, time.UTC), int64(27), "MARSEILLE",
		"Marcel", "Martin", "Paris", time.Date(1976, time.November, 24, 0, 0, 0, 0, time.UTC), int64(39), "PARIS",
		"Marion", "Leroy", "Lyon", time.Date(1982, time.October, 27, 0, 0, 0, 0, time.UTC), int64(285), "LYON",
		"Paul", "Moreau", "Lyon", time.Date(1976, time.April, 19, 0, 0, 0, 0, time.UTC), int64(133), "LYON",
	)
}
