package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*Like)(nil)

// Like expression
type Like struct {
	op     internal.Operator
	column string
	format string
	value  interface{}
	isNot  bool
}

// NewLike create like expression structure
func NewLike(op internal.Operator, column, format string, value interface{}) *Like {
	return &Like{
		op:     op,
		column: column,
		format: format,
		value:  value,
		isNot:  false,
	}
}

// NewNLike create not like expression structure
func NewNLike(op internal.Operator, column, format string, value interface{}) *Like {
	return &Like{
		op:     op,
		column: column,
		format: format,
		value:  value,
		isNot:  true,
	}
}

// ToSQL return like expression
func (l *Like) ToSQL() (string, []interface{}) {
	if l == nil {
		return "", nil
	}

	symbol := ""
	if l.isNot {
		symbol = "NOT "
	}

	return fmt.Sprintf("%s %s %sLIKE ?", l.op, internal.SafeName(l.column), symbol),
		[]interface{}{fmt.Sprintf(l.format, l.value)}
}
