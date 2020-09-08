package datatable

import (
	"github.com/pkg/errors"
)

// Errors in import/csv
var (
	ErrOpenFile           = errors.New("open file")
	ErrCantReadHeaders    = errors.New("can't read headers")
	ErrReadingLine        = errors.New("could not read line")
	ErrNilDatas           = errors.New("nil datas")
	ErrWrongNumberOfTypes = errors.New("expected different number of types")
	ErrAddingColumn       = errors.New("could not add column with given type")
)

// Errors in aggregate.go
var (
	ErrNoGroupBy      = errors.New("no groupby")
	ErrNoGroups       = errors.New("no groups")
	ErrNilDatatable   = errors.New("nil datatable")
	ErrColumnNotFound = errors.New("column not found")
	ErrUnknownAgg     = errors.New("unknown agg")
	ErrCantAddColumn  = errors.New("can't add column")
)

// Errors in column.go
var (
	ErrEmptyName         = errors.New("empty name")
	ErrNilFactory        = errors.New("nil factory")
	ErrTypeAlreadyExists = errors.New("type already exists")
	ErrUnknownColumnType = errors.New("unknown column type")
)

// Errors in concat.go
var (
	ErrNoTables = errors.New("no tables")
)

// Errors in eval_expr
var (
	ErrEvaluateExprSizeMismatch = errors.New("size mismatch")
)

// Errors in join.go
var (
	ErrNilOutputDatatable  = errors.New("nil output datatable")
	ErrNoOutput            = errors.New("no output")
	ErrNilTable            = errors.New("table is nil")
	ErrNotEnoughDatatables = errors.New("not enough datatables")
	ErrNoOnClauses         = errors.New("no on clauses")
	ErrOnClauseIsNil       = errors.New("on clause is nil")
	ErrUnknownMode         = errors.New("unknown mode")
)

// Errors in mutate_column.go
var (
	ErrNilColumn           = errors.New("nil column")
	ErrNilColumnName       = errors.New("nil column name")
	ErrNilColumnType       = errors.New("nil column type")
	ErrColumnAlreadyExists = errors.New("column already exists")
	ErrFormulaeSyntax      = errors.New("formulae syntax")
	ErrNilSerie            = errors.New("nil serie")
	ErrCreateSerie         = errors.New("create serie")
)

// Errors in mutate_rows.go
var (
	ErrLengthMismatch = errors.New("length mismatch")
	ErrUpdateRow      = errors.New("update row")
)
