package internal

const (
	OperatorEmpty Operator = iota
	OperatorAnd
	OperatorOr
)

var operatorToString = map[Operator]string{
	OperatorAnd: "AND",
	OperatorOr:  "OR",
}

// Operator of expression
type Operator uint8

func (o Operator) String() string {
	str, ok := operatorToString[o]
	if !ok {
		return "unknown-operator"
	}

	return str
}
