package sqlg

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var allColumns = []string{"*"}

// Generator of SQL statement
type Generator struct {
	table string
	opts  *Options
}

// NewGenerator create generator
func NewGenerator(table string, opts ...Option) *Generator {
	return &Generator{
		table: table,
		opts:  newOptions(opts...),
	}
}

// Select return select statement and params
func (g *Generator) Select(columns ...string) (string, []interface{}) {
	if g == nil {
		return "", nil
	}

	if len(columns) == 0 {
		columns = allColumns
	}

	conditions, params := g.opts.genWhere()
	sql := bytes.NewBufferString("SELECT")
	fmt.Fprintf(sql, " %s", strings.Join(columns, ", "))
	fmt.Fprintf(sql, " FROM %s", g.table)
	sql.WriteString(sqlOrEmpty(g.opts.genForceIndex()))
	sql.WriteString(sqlOrEmpty(conditions))
	sql.WriteString(sqlOrEmpty(g.opts.genOrderBy()))
	sql.WriteString(sqlOrEmpty(g.opts.genLimit()))
	sql.WriteString(sqlOrEmpty(g.opts.genOffset()))

	return sql.String(), params
}

// SelectByStruct return select statement and params
//
// The column of the query is obtained from the tag `sqlg` of the target structure
func (g *Generator) SelectByStruct(target interface{}) (string, []interface{}, error) {
	if g == nil {
		return "", nil, nil
	}

	columns, err := getColumns(target)
	if err != nil {
		return "", nil, err
	}

	sql, params := g.Select(columns...)
	return sql, params, nil
}

func getColumns(target interface{}) ([]string, error) {
	if target == nil {
		return nil, errors.New("target can not be empty")
	}
	targetType := reflect.TypeOf(target)

	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	var columns []string
	for i := 0; i < targetType.NumField(); i++ {
		columns = append(columns, targetType.Field(i).Tag.Get("sqlg"))
	}

	return columns, nil
}

func sqlOrEmpty(str string) string {
	if str == "" {
		return ""
	}

	return fmt.Sprintf(" %s", str)
}
