
# datatable
[![Circle CI](https://circleci.com/gh/datasweet/datatable.svg?style=svg)](https://circleci.com/gh/datasweet/datatable) [![Go Report Card](https://goreportcard.com/badge/github.com/datasweet/datatable)](https://goreportcard.com/report/github.com/datasweet/datatable) [![GoDoc](https://godoc.org/github.com/datasweet/datatable?status.png)](https://godoc.org/github.com/datasweet/datatable) [![GitHub stars](https://img.shields.io/github/stars/datasweet/datatable.svg)](https://github.com/datasweet/datatable/stargazers)
[![GitHub license](https://img.shields.io/github/license/datasweet/datatable.svg)](https://github.com/datasweet/datatable/blob/master/LICENSE)

[![datasweet-logo](https://www.datasweet.fr/wp-content/uploads/2019/02/datasweet-black.png)](http://www.datasweet.fr)

datatable is a Go package to manipulate tabular data, like an excel spreadsheet. 
datatable is inspired by the pandas python package and the data.frame R structure.
Although it's production ready, be aware that we're still working on API improvements

## Installation
```
go get github.com/datasweet/datatable
```

## Features
- Create custom Series (ie custom columns). Currently available, serie.Int, serie.String, serie.Time, serie.Float64. 
- Apply expressions
- Selects (head, tail, subset)
- Sorting
- InnerJoin, LeftJoin, RightJoin, OuterJoin, Concats
- Aggregate
- Import from CSV
- Export to map, slice


### Creating a DataTable
```go
package main

import (
	"fmt"

	"github.com/datasweet/datatable"
)

func main() {
	dt := datatable.New("test")
	dt.AddColumn("champ", datatable.String, datatable.Values("Malzahar", "Xerath", "Teemo"))
	dt.AddColumn("champion", datatable.String, datatable.Expr("upper(`champ`)"))
	dt.AddColumn("win", datatable.Int, datatable.Values(10, 20, 666))
	dt.AddColumn("loose", datatable.Int, datatable.Values(6, 5, 666))
	dt.AddColumn("winRate", datatable.Float64, datatable.Expr("`win` * 100 / (`win` + `loose`)"))
	dt.AddColumn("winRate %", datatable.String, datatable.Expr(" `winRate` ~ \" %\""))
	dt.AddColumn("sum", datatable.Float64, datatable.Expr("sum(`win`)"))

	fmt.Println(dt)
}

/*
CHAMP <NULLSTRING>      CHAMPION <NULLSTRING>   WIN <NULLINT>   LOOSE <NULLINT> WINRATE <NULLFLOAT64>   WINRATE % <NULLSTRING>  SUM <NULLFLOAT64> 
Malzahar                MALZAHAR                10              6               62.5                    62.5 %                  696              
Xerath                  XERATH                  20              5               80                      80 %                    696              
Teemo                   TEEMO                   666             666             50                      50 %                    696    
*/
```


### Reading a CSV and aggregate
```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/datasweet/datatable"
	"github.com/datasweet/datatable/import/csv"
)

func main() {
	dt, err := csv.Import("csv", "phone_data.csv",
		csv.HasHeader(true),
		csv.AcceptDate("02/01/06 15:04"),
		csv.AcceptDate("2006-01"),
	)
	if err != nil {
		log.Fatalf("reading csv: %v", err)
	}

	dt.Print(os.Stdout, datatable.PrintMaxRows(24))

	dt2, err := dt.Aggregate(datatable.AggregateBy{Type: datatable.Count, Field: "index"})
	if err != nil {
		log.Fatalf("aggregate COUNT('index'): %v", err)
	}
	fmt.Println(dt2)

	groups, err := dt.GroupBy(datatable.GroupBy{
		Name: "year",
		Type: datatable.Int,
		Keyer: func(row datatable.Row) (interface{}, bool) {
			if d, ok := row["date"]; ok {
				if tm, ok := d.(time.Time); ok {
					return tm.Year(), true
				}
			}
			return nil, false
		},
	})
	if err != nil {
		log.Fatalf("GROUP BY 'year': %v", err)
	}
	dt3, err := groups.Aggregate(
		datatable.AggregateBy{Type: datatable.Sum, Field: "duration"},
		datatable.AggregateBy{Type: datatable.CountDistinct, Field: "network"},
	)
	if err != nil {
		log.Fatalf("Aggregate SUM('duration'), COUNT_DISTINCT('network') GROUP BY 'year': %v", err)
	}
	fmt.Println(dt3)
}
```

### Creating a custom serie

To create a custom serie you must provide:
- a caster function, to cast a generic value to your serie value. The signature must be func(i interface{}) T
- a comparator, to compare your serie value. The signature must be func(a, b T) int

Example with a NullInt

```go
// IntN is an alis to create the custom Serie to manage IntN
func IntN(v ...interface{}) Serie {
	s, _ := New(NullInt{}, asNullInt, compareNullInt)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

type NullInt struct {
	Int   int
	Valid bool
}

// Interface() to render the current struct as a value.
// If not provided, the serie.All() or serie.Get() wills returns the embedded value
// IE: NullInt{}
func (i NullInt) Interface() interface{} {
	if i.Valid {
		return i.Int
	}
	return nil
}

// asNullInt is our caster function
func asNullInt(i interface{}) NullInt {
	var ni NullInt
	if i == nil {
		return ni
	}

	if v, ok := i.(NullInt); ok {
		return v
	}

	if v, err := cast.ToIntE(i); err == nil {
		ni.Int = v
		ni.Valid = true
	}
	return ni
}

// compareNullInt is our comparator function
// used to sort
func compareNullInt(a, b NullInt) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	if !a.Valid {
		return Lt
  }
  if a.Int == b.Int {
		return Eq
	}
	if a.Int < b.Int {
		return Lt
	}
	return Gt
}
```

## Who are we ?
We are Datasweet, a french startup providing full service (big) data solutions.

## Questions ? problems ? suggestions ?
If you find a bug or want to request a feature, please create a [GitHub Issue](https://github.com/datasweet/datatable/issues/new).

## License
```
This software is licensed under the Apache License, version 2 ("ALv2"), quoted below.

Copyright 2017-2020 Datasweet <http://www.datasweet.fr>

Licensed under the Apache License, Version 2.0 (the "License"); you may not
use this file except in compliance with the License. You may obtain a copy of
the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations under
the License.
```
