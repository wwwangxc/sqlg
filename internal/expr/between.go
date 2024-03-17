package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*Between)(nil)

// Between expression
type Between struct {
	op     internal.Operator
	column string
	value1 interface{}
	value2 interface{}
	isNot  bool
}

// NewBetween create between expression structure
func NewBetween(op internal.Operator, column string, value1 interface{}, value2 interface{}) *Between {
	return &Between{
		op:     op,
		column: column,
		value1: value1,
		value2: value2,
		isNot:  false,
	}
}

// NewNBetween create not between expression structure
func NewNBetween(op internal.Operator, column string, value1 interface{}, value2 interface{}) *Between {
	return &Between{
		op:     op,
		column: column,
		value1: value1,
		value2: value2,
		isNot:  true,
	}
}

// ToSQL return between expression
func (b *Between) ToSQL() (string, []interface{}) {
	if b == nil {
		return "", nil
	}

	symbol := ""
	if b.isNot {
		symbol = "NOT "
	}

	return fmt.Sprintf("%s %s %sBETWEEN ? AND ?", b.op, internal.SafeName(b.column), symbol),
		[]interface{}{b.value1, b.value2}
}
