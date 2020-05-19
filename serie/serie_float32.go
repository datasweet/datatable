package serie

import (
	"github.com/spf13/cast"
)

func Float32(v ...interface{}) Serie {
	s, _ := New(float32(0), cast.ToFloat32, compareFloat32)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func Float32N(v ...interface{}) Serie {
	s, _ := New(NullFloat32{}, asNullFloat32, compareNullFloat32)
	if len(v) > 0 {
		s.Append(v...)
	}
	return s
}

func compareFloat32(a, b float32) int {
	if a == b {
		return Eq
	}
	if a < b {
		return Lt
	}
	return Gt
}

type NullFloat32 struct {
	Float32 float32
	Valid   bool
}

func (f NullFloat32) Interface() interface{} {
	if f.Valid {
		return f.Float32
	}
	return nil
}

func asNullFloat32(i interface{}) NullFloat32 {
	var ni NullFloat32
	if i == nil {
		return ni
	}

	if v, ok := i.(NullFloat32); ok {
		return v
	}

	if v, err := cast.ToFloat32E(i); err == nil {
		ni.Float32 = v
		ni.Valid = true
	}
	return ni
}

func compareNullFloat32(a, b NullFloat32) int {
	if !b.Valid {
		if !a.Valid {
			return Eq
		}
		return Gt
	}
	return compareFloat32(a.Float32, b.Float32)
}
