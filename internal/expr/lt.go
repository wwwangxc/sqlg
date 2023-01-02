package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*LT)(nil)

// LT less than expression
type LT struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewLT create less than expression structure
func NewLT(op internal.Operator, column string, value interface{}) *LT {
	return &LT{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return less than expression
func (l *LT) ToSQL() (string, []interface{}) {
	if l == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s<?", l.op, l.column), []interface{}{l.value}
}
