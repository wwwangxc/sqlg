package expr

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
)

var _ internal.Expression = (*Compound)(nil)

// Compound expression
type Compound struct {
	op    internal.Operator
	exprs []internal.Expression
}

// NewCompound create compound expression structure
func NewCompound(op internal.Operator, exprs ...internal.Expression) *Compound {
	return &Compound{
		op:    op,
		exprs: exprs,
	}
}

// ToSQL return compound expression
func (c *Compound) ToSQL() (string, []interface{}) {
	if c == nil || len(c.exprs) == 0 {
		return "", nil
	}

	if len(c.exprs) == 1 {
		sql, values := c.exprs[0].ToSQL()
		return fmt.Sprintf("%s %s", c.op, removeFirstOp(sql)), values
	}

	buffer := bytes.NewBuffer(nil)
	values := make([]interface{}, 0, len(c.exprs))
	for _, v := range c.exprs {
		sql, vals := v.ToSQL()
		fmt.Fprintf(buffer, " %s", sql)
		values = append(values, vals...)
	}

	return fmt.Sprintf("%s (%s)", c.op, removeFirstOp(buffer.String())), values

}

func removeFirstOp(str string) string {
	str = strings.TrimSpace(str)
	return str[strings.Index(str, " ")+1:]
}
