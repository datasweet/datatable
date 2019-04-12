package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTable(t *testing.T) {
	customers := New("Customers")
	dc, err := customers.AddColumn("id", Int)
	assert.NoError(t, err)
	assert.NotNil(t, dc)
	dc.Label("Client ID")

	dc, err = customers.AddColumn("prenom", String)
	assert.NoError(t, err)
	assert.NotNil(t, dc)
	dc.Hidden(true)

	dc, err = customers.AddColumn("nom", String)
	assert.NoError(t, err)
	assert.NotNil(t, dc)
	dc.Hidden(true)

	dc, err = customers.AddExprColumn("expr_nom", "`prenom` ~ ' ' ~ UPPER(`nom`)")
	assert.NoError(t, err)
	assert.NotNil(t, dc)
	dc.Label("nom")

	dc, err = customers.AddColumn("email", String)
	assert.NoError(t, err)
	assert.NotNil(t, dc)

	dc, err = customers.AddColumn("ville", String)
	assert.NoError(t, err)
	assert.NotNil(t, dc)

	customers.AppendRow(1, "Aimée", "Marechal", nil, "aime.marechal@example.com", "Paris")
	customers.AppendRow(2, "Esmée", "Lefort", nil, "esmee.lefort@example.com", "Lyon")
	customers.AppendRow(3, "Marine", "Prevost", nil, "m.prevost@example.com", "Lille")
	customers.AppendRow(4, "Luc", "Rolland", nil, "lucrolland@example.com", "Marseille")

	checkTable(t, customers,
		"Client ID", "nom", "email", "ville",
		int64(1), "Aimée MARECHAL", "aime.marechal@example.com", "Paris",
		int64(2), "Esmée LEFORT", "esmee.lefort@example.com", "Lyon",
		int64(3), "Marine PREVOST", "m.prevost@example.com", "Lille",
		int64(4), "Luc ROLLAND", "lucrolland@example.com", "Marseille",
	)

}
