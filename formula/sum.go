package formula

import "errors"

func init() {
	All["sum"] = sum
}

func sum(values ...interface{}) (interface{}, error) {
	total := float64(0)
	for _, v := range values {
		if f, ok := v.(float64); ok {
			total += f
		} else {
			return nil, errors.New("not a float64")
		}
	}
	return total, nil
}
