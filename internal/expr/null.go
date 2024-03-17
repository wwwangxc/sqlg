package expr

import (
	"fmt"

	"github.com/wwwangxc/sqlg/internal"
)

var _ (internal.Expression) = (*Null)(nil)

// Null is null expression
type Null struct {
	op     internal.Operator
	column string
	isNot  bool
}

// NewNull create is null expression structure
func NewNull(op internal.Operator, column string) *Null {
	return &Null{
		op:     op,
		column: column,
		isNot:  false,
	}
}

// NewNNull create is not null expression structure
func NewNNull(op internal.Operator, column string) *Null {
	return &Null{
		op:     op,
		column: column,
		isNot:  true,
	}
}

// ToSQL return is null expression
func (n *Null) ToSQL() (string, []interface{}) {
	if n == nil {
		return "", nil
	}

	symbol := ""
	if n.isNot {
		symbol = "NOT "
	}

	return fmt.Sprintf("%s %s IS %sNULL", n.op, internal.SafeName(n.column), symbol), nil
}
