package serie

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

// An aggregate function performs a calculation on a set of values, and returns a single value.
// Aggregate functions ignore null values.

type StatOptions struct {
	Missing *float64 // replaces missing values with a value
}

type StatOption func(opts *StatOptions)

// Missing to treats all missing values (ie no-nils) as a value
func Missing(f float64) StatOption {
	return func(opts *StatOptions) {
		opts.Missing = &f
	}
}

func (s *serie) asFloats(opt ...StatOption) []float64 {
	var options StatOptions
	for _, o := range opt {
		o(&options)
	}
	conv := AsFloat64(s, options.Missing)
	return conv.Slice().([]float64)
}

// Avg returns the average of non-nil values
// returns NaN if no value
func (s *serie) Avg(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}
	return stat.Mean(src, nil)
}

// Count returns the number of non-nil values
func (s *serie) Count(opt ...StatOption) int64 {
	src := s.All()
	return int64(len(src))

}

// CountDistinct returns the number of unique non-nil values
func (s *serie) CountDistinct(opt ...StatOption) int64 {
	src := s.asFloats(opt...)
	set := make(map[float64]bool)

	for i := 0; i < len(src); i++ {
		set[src[i]] = true
	}
	return int64(len(set))
}

// Cusum returns the cumulative sum of non-nil values
// returns NaN if no value
func (s *serie) Cusum(opt ...StatOption) []float64 {
	opts := make([]StatOption, 0, len(opt)+1)
	opts = append(opts, Missing(0))
	opts = append(opts, opt...)
	src := s.asFloats(opts...)

	if len(src) == 0 {
		return src
	}

	dst := make([]float64, 0, len(src))
	floats.CumSum(dst, src)
	return dst
}

// Max returns the maximum of non-nil values
// returns NaN if no value
func (s *serie) Max(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}
	return floats.Max(src)
}

// Min returns the minimum of non-nil values
// returns NaN if no value
func (s *serie) Min(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}
	return floats.Min(src)
}

// Median returns the median value of non-nil values
// returns NaN if no value
func (s *serie) Median(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}

	// stat.Quantile needs the input slice to be sorted.
	sort.Float64s(src)

	// computes the median of the dataset.
	return stat.Quantile(0.5, stat.Empirical, src, nil)
}

// Stddev returns the standard deviation of non-nils values
// returns NaN if no value
func (s *serie) Stddev(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}
	return stat.StdDev(src, nil)
}

// Sum returns the sum of non-nil values
func (s *serie) Sum(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return 0
	}
	return floats.Sum(src)
}

// Variance returns the variance of non-nil values
// returns NaN if no value
func (s *serie) Variance(opt ...StatOption) float64 {
	src := s.asFloats(opt...)
	if len(src) == 0 {
		return math.NaN()
	}
	return stat.Variance(src, nil)
}
