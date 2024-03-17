package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*EQ)(nil)

// EQ equal expression
type EQ struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewEQ create equal expression structure
func NewEQ(op internal.Operator, column string, value interface{}) *EQ {
	return &EQ{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return equal expression
func (e *EQ) ToSQL() (string, []interface{}) {
	if e == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s=?", e.op, internal.SafeName(e.column)), []interface{}{e.value}
}
