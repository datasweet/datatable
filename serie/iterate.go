package serie

// Iterator to creates a new iterator from the serie
func (s *serie) Iterator() Iterator {
	return &serieIterator{
		current: -1,
		serie:   s,
	}
}

// Iterator defines an iterator
// https://docs.microsoft.com/en-us/dotnet/api/system.collections.ienumerator.reset?view=netcore-3.1
type Iterator interface {
	Next() bool
	Current() interface{}
	Reset()
}

type serieIterator struct {
	current int
	serie   *serie
}

func (it *serieIterator) Next() bool {
	it.current++
	if it.current >= it.serie.Len() {
		return false
	}
	return true
}

func (it *serieIterator) Current() interface{} {
	return it.serie.Get(it.current)
}

func (it *serieIterator) Reset() {
	it.current = -1
}
