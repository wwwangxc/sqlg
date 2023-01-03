package sqlg

// CompExpr compound expression
type CompExpr struct {
	m map[string]Expr
	s []string
}

// NewCompExpr create compound expression
func NewCompExpr() *CompExpr {
	return &CompExpr{
		m: map[string]Expr{},
		s: []string{},
	}
}

// Put expression into the compound expression
func (e *CompExpr) Put(column string, expr Expr) {
	if e == nil {
		return
	}

	if e.m == nil {
		e.m = map[string]Expr{}
	}

	if !e.exist(column) {
		e.s = append(e.s, column)
	}

	e.m[column] = expr
}

func (e *CompExpr) exist(column string) bool {
	if e.empty() {
		return false
	}

	_, exist := e.m[column]
	return exist
}

func (e *CompExpr) each(f func(column string, expr Expr)) {
	if e.empty() {
		return
	}

	for _, column := range e.s {
		e, exist := e.m[column]
		if !exist {
			continue
		}

		f(column, e)
	}
}

func (e *CompExpr) size() int {
	if e.empty() {
		return 0
	}

	return len(e.m)
}

func (e *CompExpr) empty() bool {
	return e == nil || len(e.m) == 0
}
