package expr

import (
	"github.com/datasweet/datatable/cast"
	"github.com/montanaflynn/stats"
)

// AVG(x...)
// Returns the average of all values of X
var avg = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Mean(values); err == nil {
			return s
		}
	}
	return nil
}

// COUNT(x...)
// Returns the number of values of X
var count = func(x ...interface{}) interface{} {
	values := flatten(x)
	return len(values)
}

// COUNT_DISTINCT(x...)
// Returns the number of unique values of X
var countdistinct = func(x ...interface{}) interface{} {
	values := flatten(x)
	m := make(map[float64]bool)
	for _, v := range values {
		cv, ok := cast.AsFloat(v)
		if !ok {
			return nil
		}
		m[cv] = true
	}
	return len(m)
}

// CUSUM(x...)
// Returns the cumulative sum of values of X
var cusum = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.CumulativeSum(values); err == nil {
			return s
		}
	}
	return nil
}

// MAX(x...)
// Returns the maximum value of X
var max = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Max(values); err == nil {
			return s
		}
	}
	return nil
}

// MEDIAN(x...)
// Returns the median of all values of X
var median = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Median(values); err == nil {
			return s
		}
	}
	return nil
}

// MIN(x...)
// Returns the minimum value of X
var min = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Min(values); err == nil {
			return s
		}
	}
	return nil
}

// PERCENTILE(r, x...)
// Returns the percentile rank R of X
var percentile = func(r interface{}, x ...interface{}) interface{} {
	if rank, rok := cast.AsFloat(r); rok {
		if values, ok := cast.AsFloatArray(flatten(x)...); ok {
			if s, err := stats.Percentile(values, rank); err == nil {
				return s
			}
		}
	}
	return nil
}

// STDDEV(x...)
// Returns the standard deviation of X
var stddev = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.StandardDeviation(values); err == nil {
			return s
		}
	}
	return nil
}

// SUM(x...)
// Returns the sum of all values of X
var sum = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Sum(values); err == nil {
			return s
		}
	}
	return nil
}

// VARIANCE(x...)
// Returns the variance of X
var variance = func(x ...interface{}) interface{} {
	if values, ok := cast.AsFloatArray(flatten(x)...); ok {
		if s, err := stats.Variance(values); err == nil {
			return s
		}
	}
	return nil
}
