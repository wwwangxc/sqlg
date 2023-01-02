package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*GTE)(nil)

// GTE greater than or equal expression
type GTE struct {
	op     internal.Operator
	column string
	value  interface{}
}

// NewGTE create greater than or equal expression structure
func NewGTE(op internal.Operator, column string, value interface{}) *GTE {
	return &GTE{
		op:     op,
		column: column,
		value:  value,
	}
}

// ToSQL return greater than or equal expression
func (g *GTE) ToSQL() (string, []interface{}) {
	if g == nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s>=?", g.op, g.column), []interface{}{g.value}
}
