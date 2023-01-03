package sqlg

import (
	"github.com/wwwangxc/sqlg/internal"
	"github.com/wwwangxc/sqlg/internal/expr"
)

// Expr expression
type Expr func(op internal.Operator, column string) internal.Expression

// EQ equal expression
func EQ(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewEQ(op, column, value)
	}
}

// NEQ not equal expression
func NEQ(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNEQ(op, column, value)
	}
}

// GT greater than expression
func GT(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewGT(op, column, value)
	}
}

// GTE greater than or equal expression
func GTE(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewGTE(op, column, value)
	}
}

// LT less than expression
func LT(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewLT(op, column, value)
	}
}

// LTE less than or equal expression
func LTE(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewLTE(op, column, value)
	}
}

// In expression
func In(values []interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewIn(op, column, values)
	}
}

// NIn not in expression
func NIn(values []interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNIn(op, column, values)
	}
}

// Between expression
func Between(value1, value2 interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewBetween(op, column, value1, value2)
	}
}

// NBetween not between expression
func NBetween(value1, value2 interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNBetween(op, column, value1, value2)
	}
}

// Like expression
func Like(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewLike(op, column, "%%%s%%", value)
	}
}

// NLike not like expression
func NLike(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNLike(op, column, "%%%s%%", value)
	}
}

// LikePrefix expression
func LikePrefix(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewLike(op, column, "%s%%", value)
	}
}

// NLikePrefix not like prefix expression
func NLikePrefix(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNLike(op, column, "%s%%", value)
	}
}

// LikeSuffix expression
func LikeSuffix(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewLike(op, column, "%%%s", value)
	}
}

// NLikeSuffix not like suffix expression
func NLikeSuffix(value interface{}) Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNLike(op, column, "%%%s", value)
	}
}

// Null expression
func Null() Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNull(op, column)
	}
}

// NNull not null expression
func NNull() Expr {
	return func(op internal.Operator, column string) internal.Expression {
		return expr.NewNNull(op, column)
	}
}
