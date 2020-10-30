package datatable_test

import (
	"encoding/json"
	"testing"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func sampleForExport(t *testing.T) *datatable.DataTable {
	customers := datatable.New("Customers")
	err := customers.AddColumn("id", datatable.Int)
	assert.NoError(t, err)

	err = customers.AddColumn("prenom", datatable.String)
	assert.NoError(t, err)

	err = customers.AddColumn("nom", datatable.String)
	assert.NoError(t, err)
	//dc.Hidden(true)

	err = customers.AddColumn("expr_nom", datatable.String, datatable.Expr("`prenom` ~ ' ' ~ UPPER(`nom`)"))
	assert.NoError(t, err)
	//dc.Label("nom")

	err = customers.AddColumn("email", datatable.String)
	assert.NoError(t, err)

	err = customers.AddColumn("ville", datatable.String)
	assert.NoError(t, err)

	customers.AppendRow(1, "Aimée", "Marechal", nil, "aime.marechal@example.com", "Paris")
	customers.AppendRow(2, "Esmée", "Lefort", nil, "esmee.lefort@example.com", "Lyon")
	customers.AppendRow(3, "Marine", "Prevost", nil, "m.prevost@example.com", "Lille")
	customers.AppendRow(4, "Luc", "Rolland", nil, "lucrolland@example.com", "Marseille")

	// Change structs
	assert.NoError(t, customers.RenameColumn("id", "Client ID"))
	customers.HideColumn("prenom")
	customers.HideColumn("nom")
	assert.Error(t, customers.RenameColumn("expr_nom", "nom"))
	assert.NoError(t, customers.RenameColumn("expr_nom", "Nom"))

	checkTable(t, customers,
		"Client ID", "Nom", "email", "ville",
		1, "Aimée MARECHAL", "aime.marechal@example.com", "Paris",
		2, "Esmée LEFORT", "esmee.lefort@example.com", "Lyon",
		3, "Marine PREVOST", "m.prevost@example.com", "Lille",
		4, "Luc ROLLAND", "lucrolland@example.com", "Marseille",
	)

	return customers
}

func TestToTable(t *testing.T) {
	dt := sampleForExport(t)
	out := dt.ToTable()
	assert.NotNil(t, out)

	expected := `[
		["Client ID", "Nom", "email", "ville"],
		[1, "Aimée MARECHAL", "aime.marechal@example.com", "Paris"],
		[2, "Esmée LEFORT", "esmee.lefort@example.com", "Lyon"],
		[3, "Marine PREVOST", "m.prevost@example.com", "Lille"],
		[4, "Luc ROLLAND", "lucrolland@example.com", "Marseille"]
	]`

	bytes, err := json.Marshal(out)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(bytes))

	out2 := dt.ToTable(datatable.ExportHidden(true))
	assert.NotNil(t, out2)

	expected2 := `[
		["Client ID", "prenom", "nom", "Nom", "email", "ville"],
		[1, "Aimée", "Marechal", "Aimée MARECHAL", "aime.marechal@example.com", "Paris"],
		[2, "Esmée", "Lefort", "Esmée LEFORT", "esmee.lefort@example.com", "Lyon"],
		[3, "Marine", "Prevost", "Marine PREVOST", "m.prevost@example.com", "Lille"],
		[4, "Luc", "Rolland", "Luc ROLLAND", "lucrolland@example.com", "Marseille"]
	]`
	bytes, err = json.Marshal(out2)
	assert.NoError(t, err)
	assert.JSONEq(t, expected2, string(bytes))
}

func TestToMap(t *testing.T) {
	dt := sampleForExport(t)
	out := dt.ToMap()
	assert.NotNil(t, out)

	expected := `[
	{ "Client ID":1, "Nom":"Aimée MARECHAL", "email":"aime.marechal@example.com", "ville":"Paris" },
	{ "Client ID":2, "Nom":"Esmée LEFORT", "email":"esmee.lefort@example.com", "ville":"Lyon" },
	{ "Client ID":3, "Nom":"Marine PREVOST", "email":"m.prevost@example.com", "ville":"Lille" },
	{ "Client ID":4, "Nom":"Luc ROLLAND", "email":"lucrolland@example.com", "ville":"Marseille" }
]`

	bytes, err := json.Marshal(out)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(bytes))

	out2 := dt.ToMap(datatable.ExportHidden(true))
	assert.NotNil(t, out2)

	expected2 := `[
	{ "Client ID":1, "prenom": "Aimée", "nom": "Marechal", "Nom":"Aimée MARECHAL", "email":"aime.marechal@example.com", "ville":"Paris" },
	{ "Client ID":2, "prenom": "Esmée", "nom": "Lefort", "Nom":"Esmée LEFORT", "email":"esmee.lefort@example.com", "ville":"Lyon" },
	{ "Client ID":3, "prenom": "Marine", "nom": "Prevost", "Nom":"Marine PREVOST", "email":"m.prevost@example.com", "ville":"Lille" },
	{ "Client ID":4, "prenom": "Luc", "nom": "Rolland", "Nom":"Luc ROLLAND", "email":"lucrolland@example.com", "ville":"Marseille" }
]`

	bytes, err = json.Marshal(out2)
	assert.NoError(t, err)
	assert.JSONEq(t, expected2, string(bytes))
}

func TestToSchema(t *testing.T) {
	dt := sampleForExport(t)
	schema := dt.ToSchema()
	assert.NotNil(t, schema)
	assert.Equal(t, "Customers", schema.Name)
	assert.Equal(t, []datatable.SchemaColumn{
		datatable.SchemaColumn{"Client ID", "NullInt"},
		datatable.SchemaColumn{"Nom", "NullString"},
		datatable.SchemaColumn{"email", "NullString"},
		datatable.SchemaColumn{"ville", "NullString"},
	}, schema.Columns)
	assert.Len(t, schema.Rows, 4)
	assert.Equal(t, []interface{}{1, "Aimée MARECHAL", "aime.marechal@example.com", "Paris"}, schema.Rows[0])
	assert.Equal(t, []interface{}{2, "Esmée LEFORT", "esmee.lefort@example.com", "Lyon"}, schema.Rows[1])
	assert.Equal(t, []interface{}{3, "Marine PREVOST", "m.prevost@example.com", "Lille"}, schema.Rows[2])
	assert.Equal(t, []interface{}{4, "Luc ROLLAND", "lucrolland@example.com", "Marseille"}, schema.Rows[3])

	schema2 := dt.ToSchema(datatable.ExportHidden(true))
	assert.NotNil(t, schema2)
	assert.Equal(t, "Customers", schema2.Name)
	assert.Equal(t, []datatable.SchemaColumn{
		datatable.SchemaColumn{"Client ID", "NullInt"},
		datatable.SchemaColumn{"prenom", "NullString"},
		datatable.SchemaColumn{"nom", "NullString"},
		datatable.SchemaColumn{"Nom", "NullString"},
		datatable.SchemaColumn{"email", "NullString"},
		datatable.SchemaColumn{"ville", "NullString"},
	}, schema2.Columns)
	assert.Len(t, schema2.Rows, 4)
	assert.Equal(t, []interface{}{1, "Aimée", "Marechal", "Aimée MARECHAL", "aime.marechal@example.com", "Paris"}, schema2.Rows[0])
	assert.Equal(t, []interface{}{2, "Esmée", "Lefort", "Esmée LEFORT", "esmee.lefort@example.com", "Lyon"}, schema2.Rows[1])
	assert.Equal(t, []interface{}{3, "Marine", "Prevost", "Marine PREVOST", "m.prevost@example.com", "Lille"}, schema2.Rows[2])
	assert.Equal(t, []interface{}{4, "Luc", "Rolland", "Luc ROLLAND", "lucrolland@example.com", "Marseille"}, schema2.Rows[3])
}
