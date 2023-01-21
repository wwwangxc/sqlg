package sqlg

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
	"github.com/wwwangxc/sqlg/internal/expr"
)

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

// WithGroupBy append group by condition
//
// EXP:
//   GROUP BY ${column1}, ${column2}
func WithGroupBy(columns ...string) Option {
	return func(o *Options) {
		o.groupBy = append(o.groupBy, columns...)
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

// ForceIndex set force index
//
// EXP:
//   FORCE INDEX(${index})
func ForceIndex(index string) Option {
	return func(o *Options) {
		o.forceIndex = index
	}
}

// OnDuplicateKeyUpdate for generate insert statment
//
// EXP:
//   ON DUPLICATE KEY UPDATE ${column}=${value}, ${column}=${value}
func OnDuplicateKeyUpdate(assExpr *AssExpr) Option {
	return func(o *Options) {
		o.onDuplicateKeyUpdate = assExpr
	}
}

// ForUpdate set for update symbol
//
// EXP:
//   FOR UPDATE
func ForUpdate() Option {
	return func(o *Options) {
		o.forUpdate = true
	}
}
