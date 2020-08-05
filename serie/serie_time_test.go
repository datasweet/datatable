package serie_test

import (
	"testing"
	"time"

	"github.com/datasweet/datatable/serie"
)

func TestSerieTime(t *testing.T) {
	s := serie.Time()
	s.Append(1551435220270, "2019-03-01", "2019-03-01 10:13:40", "2019-03-01T10:13:40.27Z", "01/03/2019", "01/03/2019 10:13:40", "wrong")

	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	assertSerieEq(t, s, timestamp, date, datetime, timestamp, date, datetime, time.Time{})

	s = serie.Time("02/01/06", "02/01/06 15:04:05")
	s.Append("01/03/19", "01/03/19 10:13:40", "wrong")
	assertSerieEq(t, s, date, datetime, time.Time{})

}

func TestSerieTimeN(t *testing.T) {
	s := serie.TimeN()
	s.Append(1551435220270, "2019-03-01", "2019-03-01 10:13:40", "2019-03-01T10:13:40.27Z", "01/03/2019", "01/03/2019 10:13:40", "wrong")

	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	assertSerieEq(t, s, timestamp, date, datetime, timestamp, date, datetime, nil)

	s = serie.TimeN("02/01/06", "02/01/06 15:04:05")
	s.Append("01/03/19", "01/03/19 10:13:40", "wrong")
	assertSerieEq(t, s, date, datetime, nil)

}
