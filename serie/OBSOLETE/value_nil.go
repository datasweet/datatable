package serie

type nilValue struct {
	typ ValueType
}

func NewNilValue(typ ValueType) Value {
	return &nilValue{
		typ: typ,
	}
}

func (value *nilValue) Type() ValueType {
	return value.typ
}

func (value *nilValue) Val() interface{} {
	return nil
}

func (value *nilValue) Set(v interface{}) {
}

func (value *nilValue) IsValid() bool {
	return false
}

func (value *nilValue) Clone() Value {
	// same memory
	return value
}

func (value *nilValue) String() string {
	return nullValueStr
}
