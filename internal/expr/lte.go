package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*LTE)(nil)

// LTE less than or equal expression
type LTE struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewLTE create less than or equal expression structure
func NewLTE(op internal.Operator, column string, value interface{}) *LTE {
	return &LTE{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return less than or equal expression
func (l *LTE) ToSQL() (string, []interface{}) {
	if l == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s<=?", l.op, internal.SafeName(l.column)), []interface{}{l.value}
}
