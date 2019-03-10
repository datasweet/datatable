package expr

import (
	"github.com/datasweet/datatable/cast"
)

// DATE_DIFF(x date, y date)
// Returns the difference in days between X and Y (X - Y)
var dateDiff binaryOperatorFunc = func(x interface{}, y interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	cy, oky := cast.AsDatetime(y)
	if !okx || !oky {
		return nil
	}
	return cx.Sub(cy).Hours() / 24
}

// DAY(x date)
// Returns the day of X
var day unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return cx.Day()
}

// HOUR(x date)
// Returns the hours in X in UTC timezone
var hour unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return cx.Hour()
}

// MINUTE(x date)
// Returns the minutes  in X in UTC timezone
var minute unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return cx.Minute()
}

// MONTH(x date)
// Returns the month in X
var month unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return int(cx.Month())
}

// QUARTER(x date)
// Returns the quarter of X
var quarter unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return ((int(cx.Month())-1)/3 + 1)
}

// SECOND(x date)
// Returns the seconds in X in UTC timezone
var second unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return cx.Second()
}

// WEEK(x date)
// Returns the week of X from start of year as per ISO 8601 standard
var week unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	_, w := cx.ISOWeek()
	return w
}

// WEEKDAY(x date)
// Returns the day of the week of X
var weekday unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return int(cx.Weekday())
}

// YEAR(x date)
// Returns the year of X
var year unaryOperatorFunc = func(x interface{}) interface{} {
	cx, okx := cast.AsDatetime(x)
	if !okx {
		return nil
	}
	return cx.Year()
}
