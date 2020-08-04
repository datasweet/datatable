package datatable

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cespare/xxhash"
	"github.com/datasweet/datatable/serie"
	"github.com/pkg/errors"
)

// GroupBy defines the group by  configuration
// Name is the name of the output column
// Type is the type of the output column
// Keyer is our main function to aggregate
type GroupBy struct {
	Name  string
	Type  ColumnType
	Keyer func(row Row) (interface{}, bool)
}

// AggregationType defines the avalaible aggregation
type AggregationType uint8

const (
	Avg AggregationType = iota
	Count
	CountDistinct
	Cusum
	Max
	Min
	Median
	Stddev
	Sum
	Variance
)

func (a AggregationType) String() string {
	switch a {
	case Avg:
		return "avg"
	case Count:
		return "count"
	case CountDistinct:
		return "count_distinct"
	case Cusum:
		return "cusum"
	case Max:
		return "max"
	case Min:
		return "min"
	case Median:
		return "median"
	case Stddev:
		return "stddev"
	case Sum:
		return "sum"
	case Variance:
		return "variance"
	default:
		panic("unkwown aggregation type")
	}
}

// AggregateBy defines the aggregation
type AggregateBy struct {
	Type  AggregationType
	Field string
}

// GroupBy splits our datatable by group
func (dt *DataTable) GroupBy(by ...GroupBy) (*Groups, error) {
	if len(by) == 0 {
		return nil, errors.New("no groupby")
	}

	var groups []*group
	gindex := make(map[uint64]int)

	for pos := 0; pos < dt.nrows; pos++ {
		row := dt.Row(pos)
		buf := bytes.NewBuffer(nil)
		enc := gob.NewEncoder(buf)

		buckets := make([]interface{}, len(by))

		for i, k := range by {
			k := &k
			if v, ok := k.Keyer(row); ok {
				buckets[i] = v
				enc.Encode(v)
			}
		}

		hash := xxhash.Sum64(buf.Bytes())

		if at, ok := gindex[hash]; ok {
			groups[at].Rows = append(groups[at].Rows, pos)
		} else {
			gindex[hash] = len(groups)
			groups = append(groups, &group{
				Key:     hash,
				Buckets: buckets,
				Rows:    []int{pos},
			})
		}
	}
	return &Groups{dt: dt, groups: groups, by: by}, nil
}

// Aggregate aggregates some field
func (dt *DataTable) Aggregate(by ...AggregateBy) (*DataTable, error) {
	g := &Groups{
		dt: dt,
		groups: []*group{
			&group{TakeAll: true},
		},
	}
	return g.Aggregate(by...)
}

// Groups
type Groups struct {
	dt     *DataTable
	by     []GroupBy
	groups []*group
}

type group struct {
	Key     uint64
	Buckets []interface{}
	Rows    []int
	TakeAll bool
}

// Aggregate our groups
func (g *Groups) Aggregate(aggs ...AggregateBy) (*DataTable, error) {
	if g == nil {
		return nil, errors.New("no groups")
	}

	if g.dt == nil {
		return nil, errors.New("nil datatable")
	}

	// check cols
	series := make(map[string]serie.Serie)
	for _, agg := range aggs {
		col := g.dt.Column(agg.Field)
		if col == nil {
			return nil, errors.Errorf("column '%s' not found", agg.Field)
		}
		switch agg.Type {
		case Avg, Count, CountDistinct, Cusum, Max, Min, Median, Stddev, Sum, Variance:
			series[agg.Field] = col.(*column).serie
		default:
			return nil, errors.New("unknown agg")
		}
	}

	out := New(g.dt.name)

	// create columns
	for _, by := range g.by {
		typ := by.Type
		if len(typ) == 0 {
			typ = Raw
		}
		if err := out.AddColumn(by.Name, typ); err != nil {
			return nil, errors.Wrapf(err, "can't add column '%s'", by.Name)
		}
	}
	for _, agg := range aggs {
		name := fmt.Sprintf("%s_%s", agg.Type, agg.Field)
		typ := Float64
		switch agg.Type {
		case Count, CountDistinct:
			typ = Int64
		default:
		}
		if err := out.AddColumn(name, typ); err != nil {
			return nil, errors.Wrapf(err, "can't add column '%s'", name)
		}
	}

	// aggregate the series
	for _, group := range g.groups {
		values := make([]interface{}, 0, len(group.Buckets)+len(aggs))
		values = append(values, group.Buckets...)

		for _, agg := range aggs {
			serie := series[agg.Field]

			if !group.TakeAll {
				serie = serie.Pick(group.Rows...)
			}

			fmt.Println("stop serie", serie)

			switch agg.Type {
			case Avg:
				values = append(values, serie.Avg())
			case Count:
				values = append(values, serie.Count())
			case CountDistinct:
				values = append(values, serie.CountDistinct())
			case Cusum:
				values = append(values, serie.Cusum())
			case Max:
				values = append(values, serie.Max())
			case Min:
				values = append(values, serie.Min())
			case Median:
				values = append(values, serie.Median())
			case Stddev:
				values = append(values, serie.Stddev())
			case Sum:
				values = append(values, serie.Sum())
			case Variance:
				values = append(values, serie.Variance())
			}
		}
		out.AppendRow(values...)
	}

	return out, nil
}

// func ByColumn(name string) {

// 	return func(dt *DataTable) GroupBy {
// 		return GroupBy{
// 			Name: name,
// 			Type: dt.Column(name).Type(),
// 			Keyer:
// 		}

// 	}
// }
