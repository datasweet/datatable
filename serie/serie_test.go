package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/datasweet/datatable/serie"
// 	"github.com/stretchr/testify/assert"
// )

func TestNewInt64Serie(t *testing.T) {
	s := serie.NewInt64()
	assert.NotNil(t, s)
	assert.Equal(t, serie.Int64, s.Type())
	assert.Equal(t, 0, s.Len())
	assertSerieEq(t, s, "nil")

	s.Append(1, 2, 3, 4, 5, nil, "10", "teemo", true)
	assert.Equal(t, 9, s.Len())
	assertSerieEq(t, s, "1 2 3 4 5 #NULL! 10 #NULL! 1")

	s.Prepend(-1, -2, -3, -4, "-10", "teemo", false)
	assert.Equal(t, 16, s.Len())
	assertSerieEq(t, s, "-1 -2 -3 -4 -10 #NULL! 0 1 2 3 4 5 #NULL! 10 #NULL! 1")

	assert.NoError(t, s.Insert(7, 100, 101, 102))
	assert.Equal(t, 19, s.Len())
	assertSerieEq(t, s, "-1 -2 -3 -4 -10 #NULL! 100 101 102 0 1 2 3 4 5 #NULL! 10 #NULL! 1")
}

// func TestNewScalar(t *testing.T) {
// 	s := serie.New(180604)
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Int, s.Type())
// 	assert.Equal(t, 1, s.Len())
// 	fmt.Println(s)

// 	// Len
// 	s = serie.New(180604, serie.Len(5))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Int, s.Type())
// 	assert.Equal(t, 5, s.Len())
// 	fmt.Println(s)

// 	// Force type
// 	s = serie.New(180604, serie.Type(serie.Bool))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Bool, s.Type())
// 	assert.Equal(t, 1, s.Len())
// 	fmt.Println(s)

// 	// No detect type
// 	s = serie.New(180604, serie.DetectType(false))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Raw, s.Type())
// 	assert.Equal(t, 1, s.Len())
// 	fmt.Println(s)
// }

// func TestNewSlice(t *testing.T) {
// 	s := serie.New([]int{1, 2, 3, 4, 5})
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Int, s.Type())
// 	assert.Equal(t, 5, s.Len())
// 	fmt.Println(s)

// 	s = serie.New([]int{1, 2, 3, 4, 5}, serie.Len(4))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Int, s.Type())
// 	assert.Equal(t, 4, s.Len())
// 	fmt.Println(s)

// 	s = serie.New([]int{1, 2, 3, 4, 5}, serie.Len(7))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Int, s.Type())
// 	assert.Equal(t, 7, s.Len())
// 	fmt.Println(s)

// 	s = serie.New([]int{1, 2, 3, 4, 5}, serie.Type(serie.Bool))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Bool, s.Type())
// 	assert.Equal(t, 5, s.Len())
// 	fmt.Println(s)

// 	s = serie.New([]int{1, 2, 3, 4, 5}, serie.DetectType(false))
// 	assert.NotNil(t, s)
// 	assert.Equal(t, serie.Raw, s.Type())
// 	assert.Equal(t, 5, s.Len())
// 	fmt.Println(s)
// }

// // id:  [1,2,3,4]
// // prenom: ["Aimée", "Marechal", "Lefort", "Prevost"]
// // ville: ["Paris, "Lyon", "Lille", "Marseille"]

// // user_id: [1,1,2,3,5]
// // prix: [203.14,124,149.45,235.35,47.57]

// // JOIN
// //
