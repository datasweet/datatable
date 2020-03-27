package serie_test

import (
	"strings"
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func assertSerieEq(t *testing.T, s serie.Serie, val ...string) {
	assert.NotNil(t, s)
	assert.Equal(t,
		strings.TrimSpace(strings.Join(val, " ")),
		strings.TrimSpace(s.Print(serie.PrintType(false), serie.PrintRowNumber(false), serie.PrintValueSeparator(" "))),
	)
}
