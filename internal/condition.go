package internal

import (
	"bytes"
	"fmt"
	"strings"
)

// Condition of statement
type Condition struct {
	exprs []Expression
}

// Append condition
func (c *Condition) Append(expr Expression) {
	c.exprs = append(c.exprs, expr)
}

// ToSQL return condition of statement
func (c *Condition) ToSQL() (string, []interface{}) {
	if c == nil {
		return "", nil
	}

	buffer := bytes.NewBuffer(nil)
	values := make([]interface{}, 0, len(c.exprs))
	for _, v := range c.exprs {
		sql, val := v.ToSQL()
		fmt.Fprintf(buffer, " %s", sql)
		values = append(values, val...)
	}

	sql := strings.TrimSpace(buffer.String())
	return sql[strings.Index(sql, " ")+1:], values
}
