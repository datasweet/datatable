package formula

import "errors"

func init() {
	All["avg"] = avg
}

func avg(values ...interface{}) (interface{}, error) {
	n := len(values)
	if n == 0 {
		return nil, errors.New("can't divide by 0")
	}

	total, err := sum(values...)
	if err != nil {
		return nil, err
	}

	return total.(float64) / float64(n), nil
}
