package expr

import (
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*In)(nil)

// In expression
type In struct {
	op     internal.Operator
	column string
	values []interface{}
	isNot  bool
}

// NewIn create in expression structure
func NewIn(op internal.Operator, column string, values []interface{}) *In {
	return &In{
		op:     op,
		column: column,
		values: values,
		isNot:  false,
	}
}

// NewNIn create not in expression structure
func NewNIn(op internal.Operator, column string, values []interface{}) *In {
	return &In{
		op:     op,
		column: column,
		values: values,
		isNot:  true,
	}
}

// ToSQL return in expression
func (i *In) ToSQL() (string, []interface{}) {
	if i == nil {
		return "", nil
	}

	var placeholder string
	if len(i.values) > 0 {
		placeholder = strings.Repeat(",?", len(i.values))[1:]
	}

	symbol := ""
	if i.isNot {
		symbol = "NOT "
	}

	return fmt.Sprintf("%s %s %sIN (%s)", i.op, i.column, symbol, placeholder), i.values
}
