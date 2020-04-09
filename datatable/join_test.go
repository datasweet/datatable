package datatable

import (
	"github.com/datasweet/datatable/serie"
)

// https://sql.sh/cours/jointures/inner-join
func sampleForJoin() (DataTable, DataTable) {
	customers := New("Customers")
	customers.AddColumn("id", serie.Int())
	customers.AddColumn("prenom", serie.String())
	customers.AddColumn("nom", serie.String())
	customers.AddColumn("email", serie.String())
	customers.AddColumn("ville", serie.String())
	customers.AppendRow(1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris")
	customers.AppendRow(2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon")
	customers.AppendRow(3, "Marine", "Prevost", "m.prevost@example.com", "Lille")
	customers.AppendRow(4, "Luc", "Rolland", "lucrolland@example.com", "Marseille")

	orders := New("Orders")
	orders.AddColumn("user_id", serie.Int(1, 1, 2, 3, 5))
	orders.AddColumn("date_achat", serie.String("2013-01-23", "2013-02-14", "2013-02-17", "2013-02-21", "2013-03-02"))
	orders.AddColumn("num_facture", serie.String("A00103", "A00104", "A00105", "A00106", "A00107"))
	orders.AddColumn("prix_total", serie.Float64(2013.14, 124.00, 149.45, 235.35, 47.58))
	// orders.AppendRow(1, "2013-01-23", "A00103", 203.14)
	// orders.AppendRow(1, "2013-02-14", "A00104", 124.00)
	// orders.AppendRow(2, "2013-02-17", "A00105", 149.45)
	// orders.AppendRow(3, "2013-02-21", "A00106", 235.35)
	// orders.AppendRow(5, "2013-03-02", "A00107", 47.58)

	return customers, orders
}

// func sampleMultipleForJoin() []DataTable {
// 	count := New("count")
// 	count.AddColumn("terms_speaker", String)
// 	count.AddColumn("line_id", Int)
// 	count.AppendRow("IAGO", 1161)
// 	count.AppendRow("OTHELLO", 928)
// 	count.AppendRow("DESDEMONA", 404)
// 	count.AppendRow("CASSIO", 308)
// 	count.AppendRow("LODOVICO", 252)

// 	min := New("min")
// 	min.AddColumn("terms_speaker", String)
// 	min.AddColumn("line_id", Int)
// 	min.AppendRow("GRATIANO", 75203)
// 	min.AppendRow("LODOVICO", 74664)
// 	min.AppendRow("CASSIO", 75750)
// 	min.AppendRow("Gentleman", 73634)
// 	min.AppendRow("First Musician", 7356)

// 	max := New("max")
// 	max.AddColumn("terms_speaker", String)
// 	max.AddColumn("line_id", Int)
// 	max.AppendRow("LODOVICO", 75762)
// 	max.AppendRow("BIANCA", 74379)
// 	max.AppendRow("OTHELLO", 75748)
// 	max.AppendRow("GRATIANO", 75745)
// 	max.AppendRow("IAGO", 7568)

// 	return []DataTable{count, min, max}
// }

// func TestJoinOn(t *testing.T) {
// 	on := On("[customers].[id]")
// 	assert.NotNil(t, on)
// 	assert.Len(t, on, 1)
// 	assert.Equal(t, "customers", on[0].Table)
// 	assert.Equal(t, "id", on[0].Field)

// 	on = On("[id]")
// 	assert.NotNil(t, on)
// 	assert.Len(t, on, 1)
// 	assert.Equal(t, "*", on[0].Table)
// 	assert.Equal(t, "id", on[0].Field)

// 	on = On("id")
// 	assert.NotNil(t, on)
// 	assert.Len(t, on, 1)
// 	assert.Equal(t, "*", on[0].Table)
// 	assert.Equal(t, "id", on[0].Field)

// 	on = On("customers.[id]")
// 	assert.NotNil(t, on)
// 	assert.Len(t, on, 1)
// 	assert.Equal(t, "*", on[0].Table)
// 	assert.Equal(t, "customers.[id]", on[0].Field)
// }

// func TestInnerJoin(t *testing.T) {
// 	customers, orders := sampleForJoin()

// 	dt, err := InnerJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 	)
// }

// func TestLeftJoin(t *testing.T) {
// 	customers, orders := sampleForJoin()

// 	dt, err := LeftJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
// 	)
// }

// func TestRightJoin(t *testing.T) {
// 	customers, orders := sampleForJoin()

// 	dt, err := RightJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 		nil, nil, nil, nil, nil, time.Date(2013, time.March, 2, 0, 0, 0, 0, time.UTC), "A00107", 47.58,
// 	)
// }

// func TestOuterJoin(t *testing.T) {
// 	customers, orders := sampleForJoin()

// 	dt, err := OuterJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"id", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 		4, "Luc", "Rolland", "lucrolland@example.com", "Marseille", nil, nil, nil,
// 		nil, nil, nil, nil, nil, time.Date(2013, time.March, 2, 0, 0, 0, 0, time.UTC), "A00107", 47.58,
// 	)
// }

// func TestJoinWithExpr(t *testing.T) {
// 	customers, orders := sampleForJoin()
// 	customers.AddExprColumn("upper_ville", "UPPER(ville)")

// 	dt, err := InnerJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	idx, dc := dt.GetColumn("upper_ville")
// 	assert.Equal(t, 5, idx)
// 	assert.Equal(t, Raw, dc.Type())
// 	assert.False(t, dc.Computed())

// 	checkTable(t, dt,
// 		"id", "prenom", "nom", "email", "ville", "upper_ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "PARIS", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", "PARIS", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", "LYON", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", "LILLE", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 	)
// }

// func TestJoinWithColumnName(t *testing.T) {
// 	customers, orders := sampleForJoin()
// 	if _, dc := customers.GetColumn("id"); dc != nil {
// 		dc.Label("ClientID")
// 	}

// 	dt, err := InnerJoin([]DataTable{customers, orders}, On("[Customers].[id]", "[Orders].[user_id]"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"ClientID", "prenom", "nom", "email", "ville", "date_achat", "num_facture", "prix_total",
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.January, 23, 0, 0, 0, 0, time.UTC), "A00103", 203.14,
// 		1, "Aimée", "Marechal", "aime.marechal@example.com", "Paris", time.Date(2013, time.February, 14, 0, 0, 0, 0, time.UTC), "A00104", 124.00,
// 		2, "Esmée", "Lefort", "esmee.lefort@example.com", "Lyon", time.Date(2013, time.February, 17, 0, 0, 0, 0, time.UTC), "A00105", 149.45,
// 		3, "Marine", "Prevost", "m.prevost@example.com", "Lille", time.Date(2013, time.February, 21, 0, 0, 0, 0, time.UTC), "A00106", 235.35,
// 	)
// }

// func TestMultipleInnerJoin(t *testing.T) {
// 	tables := sampleMultipleForJoin()
// 	dt, err := InnerJoin(tables, Using("terms_speaker"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"terms_speaker", "count.line_id", "min.line_id", "max.line_id",
// 		"LODOVICO", 252, 74664, 75762,
// 	)
// }

// func TestMultipleLeftJoin(t *testing.T) {
// 	tables := sampleMultipleForJoin()
// 	dt, err := LeftJoin(tables, Using("terms_speaker"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"terms_speaker", "count.line_id", "min.line_id", "max.line_id",
// 		"IAGO", 1161, nil, 7568,
// 		"OTHELLO", 928, nil, 75748,
// 		"DESDEMONA", 404, nil, nil,
// 		"CASSIO", 308, 75750, nil,
// 		"LODOVICO", 252, 74664, 75762,
// 	)
// }

// func TestMultipleRightJoin(t *testing.T) {
// 	tables := sampleMultipleForJoin()
// 	dt, err := RightJoin(tables, Using("terms_speaker"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"terms_speaker", "count.line_id", "min.line_id", "max.line_id",
// 		"LODOVICO", 252, 74664, 75762,
// 		nil, nil, nil, 74379,
// 		nil, nil, nil, 75748,
// 		nil, nil, nil, 75745,
// 		nil, nil, nil, 7568,
// 	)
// }

// func TestMultipleFullOuterJoin(t *testing.T) {
// 	tables := sampleMultipleForJoin()
// 	dt, err := OuterJoin(tables, Using("terms_speaker"))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, dt)

// 	checkTable(t, dt,
// 		"terms_speaker", "count.line_id", "min.line_id", "max.line_id",
// 		"IAGO", 1161, nil, 7568,
// 		"OTHELLO", 928, nil, 75748,
// 		"DESDEMONA", 404, nil, nil,
// 		"CASSIO", 308, 75750, nil,
// 		"LODOVICO", 252, 74664, 75762,
// 		nil, nil, 75203, nil,
// 		nil, nil, 73634, nil,
// 		nil, nil, 7356, nil,
// 		nil, nil, nil, 74379,
// 		nil, nil, nil, 75745,
// 	)
// }
