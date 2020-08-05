package csv_test

import (
	"fmt"
	"testing"

	"github.com/datasweet/datatable/import/csv"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	dt, err := csv.Import("csv", "../../test/phone_data.csv",
		csv.HasHeader(true),
		csv.AcceptDate("02/01/06 15:04"),
		csv.AcceptDate("2006-01"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, dt)

	fmt.Println(dt)
}
