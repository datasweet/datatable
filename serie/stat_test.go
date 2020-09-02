package serie_test

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/stat"
)

func TestAvg(t *testing.T) {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	s := serie.Float64(xs)
	assert.NotNil(t, s)
	assert.Equal(t, 9, s.Len())
	assert.Equal(t, stat.Mean(xs, nil), s.Avg())

	s = serie.Float64N(xs, "teemo", "nil")
	assert.NotNil(t, s)
	assert.Equal(t, 11, s.Len())
	assert.Equal(t, stat.Mean(xs, nil), s.Avg())
	assert.Greater(t, stat.Mean(xs, nil), s.Avg(serie.Missing(0)))
}

func TestCount(t *testing.T) {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	s := serie.Float64(xs)
	assert.NotNil(t, s)
	assert.Equal(t, 9, s.Len())
	assert.Equal(t, int64(9), s.Count())

	s = serie.Float64N(xs, "teemo", "nil")
	assert.NotNil(t, s)
	assert.Equal(t, 11, s.Len())
	assert.Equal(t, int64(9), s.Count())
}

func TestSum(t *testing.T) {
	s := serie.Float64N(1, "23", 3.14, "teemo", true, nil)
	assert.NotNil(t, s)
	assertSerieEq(t, s, float64(1), float64(23), float64(3.14), nil, float64(1), nil)
	assert.Equal(t, 28.14, s.Sum())
}

func TestMedian(t *testing.T) {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	fmt.Printf("data: %v\n", xs)

	// computes the weighted mean of the dataset.
	// we don't have any weights (ie: all weights are 1)
	// so we just pass a nil slice.
	mean := stat.Mean(xs, nil)
	variance := stat.Variance(xs, nil)
	stddev := math.Sqrt(variance)

	// stat.Quantile needs the input slice to be sorted.
	sort.Float64s(xs)
	fmt.Printf("data: %v (sorted)\n", xs)

	// computes the median of the dataset.
	// here as well, we pass a nil slice as weights.
	median := stat.Quantile(0.5, stat.Empirical, xs, nil)

	fmt.Printf("mean=     %v\n", mean)
	fmt.Printf("median=   %v\n", median)
	fmt.Printf("variance= %v\n", variance)
	fmt.Printf("std-dev=  %v\n", stddev)

}
