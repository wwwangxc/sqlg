package sqlg

// AssExpr assignment expression
type AssExpr struct {
	m map[string]interface{}
	s []string
}

// NewAssExpr create assignment expression
func NewAssExpr() *AssExpr {
	return &AssExpr{
		m: map[string]interface{}{},
		s: []string{},
	}
}

// Put column and value into the assignment expression
func (a *AssExpr) Put(column string, value interface{}) {
	if a == nil {
		return
	}

	if a.m == nil {
		a.m = map[string]interface{}{}
	}

	if !a.exist(column) {
		a.s = append(a.s, column)
	}

	a.m[column] = value
}

func (a *AssExpr) exist(column string) bool {
	if a.empty() {
		return false
	}

	_, exist := a.m[column]
	return exist
}

func (a *AssExpr) each(f func(column string, value interface{})) {
	if a.empty() {
		return
	}

	for _, column := range a.s {
		value, exist := a.m[column]
		if !exist {
			continue
		}

		f(column, value)
	}
}

func (a *AssExpr) size() int {
	if a.empty() {
		return 0
	}

	return len(a.m)
}

func (a *AssExpr) empty() bool {
	return a == nil || len(a.m) == 0
}
