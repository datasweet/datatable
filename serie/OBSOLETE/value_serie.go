package serie

const Series = ValueType("serie")

type serieValue struct {
	val   Serie
	valid bool
}

func NewSerieValue(v interface{}) Value {
	value := &serieValue{}
	value.Set(v)
	return value
}

func (value *serieValue) Type() ValueType {
	return Series
}

func (value *serieValue) Val() interface{} {
	if value.valid {
		return value.val
	}
	return nil
}

func (value *serieValue) Set(v interface{}) {
	value.val = nil
	value.valid = false

	if casted, ok := v.(Serie); ok {
		value.val = casted
		value.valid = casted != nil
	}
}

func (value *serieValue) IsValid() bool {
	return value.valid
}

func (value *serieValue) Clone() Value {
	var cpy Serie
	if value.valid {
		cpy = value.val.Clone(true)
	}
	return &serieValue{
		valid: value.valid,
		val:   cpy,
	}
}

func (value *serieValue) String() string {
	if value.valid {
		return value.val.Print()
	}
	return nullValueStr
}
