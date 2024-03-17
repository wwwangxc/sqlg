package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*GT)(nil)

// GT greater than expression
type GT struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewGT create greater than expression structure
func NewGT(op internal.Operator, column string, value interface{}) *GT {
	return &GT{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return greater than expression
func (g *GT) ToSQL() (string, []interface{}) {
	if g == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s>?", g.op, internal.SafeName(g.column)), []interface{}{g.value}
}
