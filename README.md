
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
- Apply formulas
- Selects (head, tail, subset)
- Sorting
- InnerJoin, LeftJoin, RightJoin, OuterJoin, Concats
- Export to map, slice


### Creating a DataTable
```go

import (
	"github.com/datasweet/datatable"
	"github.com/datasweet/datatable/serie"
)

func main() {
  tb := datatable.New("test")
	tb.AddColumn("champ", serie.String("Malzahar", "Xerath", "Teemo"))
	tb.AddExprColumn("champion", serie.String(), "upper(`champ`)")
	tb.AddColumn("win", serie.Int(10, 20, 666))
	tb.AddColumn("loose", serie.Int(6, 5, 666))
	tb.AddExprColumn("winRate", serie.String(), "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", serie.Float64(), "sum(`win`)")
  tb.AddExprColumn("ok", serie.Bool(), "true")

  fmt.Println(tb)
}

/*
CHAMP <STRING>	CHAMPION <STRING>	WIN <INT>	LOOSE <INT>	WINRATE <STRING>	SUM <FLOAT64>	OK <BOOL> 
Malzahar      	MALZAHAR         	10       	6          	62.5 %          	696          	true     	
Xerath        	XERATH           	20       	5          	80 %            	696          	true     	
Teemo         	TEEMO            	666      	666        	50 %            	696          	true     	
*/
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
