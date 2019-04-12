package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDataRow(t *testing.T) {
	dr := make(DataRow, 5)
	assert.NotNil(t, dr)
	dr["prenom"] = "Léon"
	dr["nom"] = "Dupuis"
	dr["ville"] = "Paris"
	dr["date_naissance"] = "1983-03-06"
	dr["total_achat"] = 135

	assert.Len(t, dr, 5)
	assert.Equal(t, "Léon", dr.Get("prenom"))
	assert.Equal(t, "Dupuis", dr.Get("nom"))
	assert.Equal(t, "Paris", dr.Get("ville"))
	assert.Equal(t, "1983-03-06", dr.Get("date_naissance"))
	assert.Equal(t, 135, dr.Get("total_achat"))

	// Change a valid value
	dr.Set("total_achat", 285)
	assert.Equal(t, 285, dr.Get("total_achat"))

	// Invalid cell
	dr.Set("dummy", "wrong")
	assert.Len(t, dr, 5)

	h, ok := dr.Hash()
	assert.True(t, ok)
	assert.NotZero(t, h)
}

func TestHashMustBeEqual(t *testing.T) {
	dr1 := make(DataRow, 5)
	dr1["prenom"] = "Léon"
	dr1["nom"] = "Dupuis"
	dr1["ville"] = "Paris"
	dr1["date_naissance"] = "1983-03-06"
	dr1["total_achat"] = 135

	h1, ok1 := dr1.Hash()
	assert.True(t, ok1)
	assert.NotZero(t, h1)

	dr2 := make(DataRow, 5)
	dr2["total_achat"] = 135
	dr2["date_naissance"] = "1983-03-06"
	dr2["ville"] = "Paris"
	dr2["nom"] = "Dupuis"
	dr2["prenom"] = "Léon"

	h2, ok2 := dr2.Hash()
	assert.True(t, ok2)
	assert.NotZero(t, h2)

	// h1 & h2 must have the same hash
	// despite the different ordering
	assert.Equal(t, h1, h2)
}

func TestHashMustBeNotEqual(t *testing.T) {
	dr1 := make(DataRow, 5)
	dr1["prenom"] = "Léon"
	dr1["nom"] = "Dupuis"
	dr1["ville"] = "Paris"
	dr1["date_naissance"] = "1983-03-06"
	dr1["total_achat"] = 135

	h1, ok1 := dr1.Hash()
	assert.True(t, ok1)
	assert.NotZero(t, h1)

	dr2 := make(DataRow, 5)
	dr2["prenom"] = "Léon"
	dr2["nom"] = "Dupuis"
	dr2["ville"] = "Paris"
	dr2["date_naissance"] = "1983-03-06"
	dr2["total_achat"] = 136

	h2, ok2 := dr2.Hash()
	assert.True(t, ok2)
	assert.NotZero(t, h2)

	// h1 & h2 must have the same hash
	// despite the different ordering
	assert.NotEqual(t, h1, h2)
}
