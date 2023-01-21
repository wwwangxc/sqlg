package sqlg

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
	"github.com/wwwangxc/sqlg/internal/expr"
)

// Options of SQL generator
type Options struct {
	where                *internal.Condition
	orderBy              []string
	limit                uint32
	offset               uint32
	forceIndex           string
	onDuplicateKeyUpdate *AssExpr
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

func (o *Options) genSet(assExpr *AssExpr) (string, []interface{}) {
	if o == nil || assExpr == nil {
		return "", nil
	}

	params := make([]interface{}, 0, assExpr.size())
	buffer := bytes.NewBuffer(nil)
	assExpr.each(func(column string, value interface{}) {
		fmt.Fprintf(buffer, ", %s=?", column)
		params = append(params, value)
	})

	sql := buffer.String()
	return fmt.Sprintf("SET %s", sql[strings.Index(sql, " ")+1:]), params
}

func (o *Options) genOnDuplicateKeyUpdate() (string, []interface{}) {
	if o == nil || o.onDuplicateKeyUpdate.empty() {
		return "", nil
	}

	params := make([]interface{}, 0, o.onDuplicateKeyUpdate.size())
	buffer := bytes.NewBuffer(nil)
	o.onDuplicateKeyUpdate.each(func(column string, value interface{}) {
		fmt.Fprintf(buffer, ", %s=?", column)
		params = append(params, value)
	})

	sql := buffer.String()
	return fmt.Sprintf("ON DUPLICATE KEY UPDATE %s", sql[strings.Index(sql, " ")+1:]), params
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

// WithExists append exists expression
//
// EXP:
//   AND EXISTS (SELECT * FROM ${table} WHERE %{expr1} AND ${expr2})
func WithExists(table string, m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || table == "" || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorAnd, column))
		})

		o.where.Append(expr.NewExists(internal.OperatorAnd, table, exprs...))
	}
}

// WithNExists append not exists expression
//
// EXP:
//   AND NOT EXISTS (SELECT * FROM ${table} WHERE %{expr1} AND ${expr2})
func WithNExists(table string, m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || table == "" || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorAnd, column))
		})

		o.where.Append(expr.NewNExists(internal.OperatorAnd, table, exprs...))
	}
}

// WithNExists append not exists expression
//
// EXP:
//   OR EXISTS (SELECT * FROM ${table} WHERE %{expr1} AND ${expr2})
func WithOrExists(table string, m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || table == "" || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorAnd, column))
		})

		o.where.Append(expr.NewExists(internal.OperatorOr, table, exprs...))
	}
}

// WithOrNExists append not exists expression
//
// EXP:
//   OR NOT EXISTS (SELECT * FROM ${table} WHERE %{expr1} AND ${expr2})
func WithOrNExists(table string, m *CompExpr) Option {
	return func(o *Options) {
		if m == nil || table == "" || m.empty() {
			return
		}

		exprs := make([]internal.Expression, 0, m.size())
		m.each(func(column string, expr Expr) {
			exprs = append(exprs, expr(internal.OperatorAnd, column))
		})

		o.where.Append(expr.NewNExists(internal.OperatorOr, table, exprs...))
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

// WithOnDuplicateKeyUpdate for generate insert statment
//
// EXP:
//   ON DUPLICATE KEY UPDATE ${column}=${value}, ${column}=${value}
func WithOnDuplicateKeyUpdate(assExpr *AssExpr) Option {
	return func(o *Options) {
		o.onDuplicateKeyUpdate = assExpr
	}
}
