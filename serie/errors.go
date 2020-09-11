package serie

import (
	"github.com/pkg/errors"
)

// Errors in mutate.go
var (
	ErrOutOfRange                      = errors.New("out of range")
	ErrCantFlattenSliceWithSet         = errors.New("can't flatten slice with set")
	ErrGrowSizeMustBeStriclyPositive   = errors.New("grow: size must be > 0")
	ErrShrinkSizeMustBeStriclyPositive = errors.New("shrink: size must be > 0")
	ErrShrinkSizeMustBeLesserThanLen   = errors.New("shrink: size must be < len")
	ErrConcatTypeMismatch              = errors.New("concat: type mismatch")
)
