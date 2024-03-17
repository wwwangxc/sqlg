package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*NEQ)(nil)

// NEQ not equal expression
type NEQ struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewNEQ create not equal expression structure
func NewNEQ(op internal.Operator, column string, value interface{}) *NEQ {
	return &NEQ{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return not equal expression
func (n *NEQ) ToSQL() (string, []interface{}) {
	if n == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s!=?", n.op, internal.SafeName(n.column)), []interface{}{n.value}
}
