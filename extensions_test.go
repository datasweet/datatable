package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	tb := New("test")
	tb.AddColumn("id", String, "0001", "0002", "0003", "0004", "0005", "0006")
	tb.AddColumn("type", String, "donut", "donut", "donut", "bar", "twist", "filled")
	tb.AddColumn("name", String, "Cake", "Raised", "Old Fashioned", "Bar", "Twist", "Filled")
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.False(t, tb.Swap("toto", "name"))
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.False(t, tb.Swap("name", "toto"))
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.True(t, tb.Swap("id", "name"))
	checkTable(t, tb,
		"name", "type", "id",
		"Cake", "donut", "0001",
		"Raised", "donut", "0002",
		"Old Fashioned", "donut", "0003",
		"Bar", "bar", "0004",
		"Twist", "twist", "0005",
		"Filled", "filled", "0006",
	)

}
