package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func TestNewSerie(t *testing.T) {
	s, err := serie.New(nil, nil, nil)
	assert.Nil(t, s)
	assert.Error(t, err)

	s, err = serie.New(1, nil, nil)
	assert.Nil(t, s)
	assert.Error(t, err)

	s, err = serie.New(1, cast.ToFloat32, func(i, j int) int {
		return serie.Eq
	})
	assert.Nil(t, s)
	assert.Error(t, err)

	s, err = serie.New(1, cast.ToInt, nil)
	assert.Nil(t, s)
	assert.Error(t, err)

	s, err = serie.New(1, cast.ToInt, func(i, j int) int {
		return serie.Eq
	})
	assert.NotNil(t, s)
	assert.NoError(t, err)
}
