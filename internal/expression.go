package internal

// Expression of condition
type Expression interface {
	// ToSQL return sql expression
	ToSQL() (string, []interface{})
}
