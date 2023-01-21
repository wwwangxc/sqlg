package expr

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*Exists)(nil)

// Exists expression
type Exists struct {
	op       internal.Operator
	table    string
	compExpr *Compound
	isNot    bool
}

// NewExists create exists expression structure
func NewExists(op internal.Operator, table string, exprs ...internal.Expression) *Exists {
	return &Exists{
		op:       op,
		table:    table,
		compExpr: NewCompound(internal.OperatorEmpty, exprs...),
		isNot:    false,
	}
}

// NewNExists create not exists expression structure
func NewNExists(op internal.Operator, table string, exprs ...internal.Expression) *Exists {
	return &Exists{
		op:       op,
		table:    table,
		compExpr: NewCompound(internal.OperatorEmpty, exprs...),
		isNot:    true,
	}
}

// ToSQL return exists expression
func (e *Exists) ToSQL() (string, []interface{}) {
	if e == nil || e.compExpr == nil || len(e.compExpr.exprs) == 0 {
		return "", nil
	}

	symbol := ""
	if e.isNot {
		symbol = "NOT "
	}

	cond, params := e.compExpr.ToSQL()
	cond = strings.TrimRight(strings.TrimLeft(removeFirstOp(cond), "("), ")")
	return fmt.Sprintf("%s %sEXISTS (SELECT * FROM %s WHERE %s)", e.op, symbol, e.table, cond), params
}
