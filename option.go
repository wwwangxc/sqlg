package sqlg

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
	"github.com/wwwangxc/sqlg/internal/expr"
)

// Options of SQL generator
type Options struct {
	where      *internal.Condition
	orderBy    []string
	limit      uint32
	offset     uint32
	forceIndex string
}

func newOptions(opts ...Option) *Options {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func defaultOptions() *Options {
	return &Options{
		where:      &internal.Condition{},
		orderBy:    []string{},
		limit:      0,
		offset:     0,
		forceIndex: "",
	}
}

func (o *Options) genForceIndex() string {
	if o == nil || o.forceIndex == "" {
		return ""
	}

	return fmt.Sprintf("FORCE INDEX (%s)", o.forceIndex)
}

func (o *Options) genWhere() (string, []interface{}) {
	if o == nil || o.where == nil {
		return "", nil
	}

	sql, values := o.where.ToSQL()
	return fmt.Sprintf("WHERE %s", sql), values
}

func (o *Options) genOrderBy() string {
	if o == nil || len(o.orderBy) == 0 {
		return ""
	}

	return fmt.Sprintf("ORDER BY %s", strings.Join(o.orderBy, ", "))
}

func (o *Options) genLimit() string {
	if o == nil || o.limit == 0 {
		return ""
	}

	return fmt.Sprintf("LIMIT %d", o.limit)
}

func (o *Options) genOffset() string {
	if o == nil || o.offset == 0 {
		return ""
	}

	return fmt.Sprintf("OFFSET %d", o.offset)
}

// Option is optional for the SQL generator
type Option func(*Options)

// WithAnd append AND expression into the condition
//
// EXP:
//   AND ${expr}
func WithAnd(column string, expr Expr) Option {
	return func(o *Options) {
		o.where.Append(expr(internal.OperatorAnd, column))
	}
}

// WithOr append OR expression into the condition
//
// EXP:
//   OR ${expr}
func WithOr(column string, expr Expr) Option {
	return func(o *Options) {
		o.where.Append(expr(internal.OperatorOr, column))
	}
}

// WithAndExprs append compound expression
//
// EXP:
//   AND (${expr1} OR ${expr2})
func WithAndExprs(m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorOr, column))
		})

		o.where.Append(expr.NewCompound(internal.OperatorAnd, exprs...))
	}
}

// WithOrExprs append compound expression
//
// EXP:
//   OR (${expr1} AND ${expr2})
func WithOrExprs(m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorAnd, column))
		})

		o.where.Append(expr.NewCompound(internal.OperatorOr, exprs...))
	}
}

// WithOrderBy append order by condition
//
// EXP:
//   ORDER BY ${column} ASC
func WithOrderBy(column string) Option {
	return func(o *Options) {
		o.orderBy = append(o.orderBy, fmt.Sprintf("%s ASC", column))
	}
}

// WithOrderByDESC append order by condition
//
// EXP:
//   ORDER BY ${column} DESC
func WithOrderByDESC(column string) Option {
	return func(o *Options) {
		o.orderBy = append(o.orderBy, fmt.Sprintf("%s DESC", column))
	}
}

// WithLimit set limit
//
// EXP:
//   LIMIT ${limit}
func WithLimit(limit uint32) Option {
	return func(o *Options) {
		o.limit = limit
	}
}

// WithOffset set offset
//
// EXP:
//   OFFSET ${offset}
func WithOffset(offset uint32) Option {
	return func(o *Options) {
		o.offset = offset
	}
}

// WithForceIndex set force index
//
// EXP:
//   FORCE INDEX(${index})
func WithForceIndex(index string) Option {
	return func(o *Options) {
		o.forceIndex = index
	}
}
