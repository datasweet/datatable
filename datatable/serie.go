package datatable

type Serie struct {
	Name   string
	Values []int
}

func (s *Serie) Append(values ...int) {
	s.Values = append(s.Values, values...)

}
