package sqlg

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
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

	where, params := g.opts.genWhere()
	sql := bytes.NewBufferString("SELECT")
	fmt.Fprintf(sql, " %s", strings.Join(internal.SafeNames(columns), ", "))
	fmt.Fprintf(sql, " FROM %s", internal.SafeName(g.table))
	sql.WriteString(sqlOrEmpty(g.opts.genForceIndex()))
	sql.WriteString(sqlOrEmpty(where))
	sql.WriteString(sqlOrEmpty(g.opts.genGroupBy()))
	sql.WriteString(sqlOrEmpty(g.opts.genOrderBy()))
	sql.WriteString(sqlOrEmpty(g.opts.genLimit()))
	sql.WriteString(sqlOrEmpty(g.opts.genOffset()))
	sql.WriteString(sqlOrEmpty(g.opts.genForUpdate()))

	return sql.String(), params
}

// SelectByStruct return select statement and params
//
// The column of the query is obtained from the tag `db` of the target structure
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

// Update return update statement and params
func (g *Generator) Update(assExpr *AssExpr) (string, []interface{}) {
	if g == nil || assExpr.empty() {
		return "", nil
	}

	set, params := g.opts.genSet(assExpr)
	where, whereParams := g.opts.genWhere()
	params = append(params, whereParams...)

	sql := bytes.NewBufferString(fmt.Sprintf("UPDATE %s", internal.SafeName(g.table)))
	sql.WriteString(sqlOrEmpty(set))
	sql.WriteString(sqlOrEmpty(where))
	sql.WriteString(sqlOrEmpty(g.opts.genOrderBy()))
	sql.WriteString(sqlOrEmpty(g.opts.genLimit()))
	sql.WriteString(sqlOrEmpty(g.opts.genOffset()))

	return sql.String(), params
}

// Delete return delete statement and params
func (g *Generator) Delete() (string, []interface{}) {
	if g == nil {
		return "", nil
	}

	where, params := g.opts.genWhere()
	sql := bytes.NewBufferString("DELETE")
	fmt.Fprintf(sql, " FROM %s", internal.SafeName(g.table))
	sql.WriteString(sqlOrEmpty(where))
	sql.WriteString(sqlOrEmpty(g.opts.genOrderBy()))
	sql.WriteString(sqlOrEmpty(g.opts.genLimit()))
	sql.WriteString(sqlOrEmpty(g.opts.genOffset()))

	return sql.String(), params
}

// Insert return insert statement and params
func (g *Generator) Insert(columns []string, records ...[]interface{}) (string, []interface{}) {
	if g == nil || len(columns) == 0 || len(records) == 0 {
		return "", nil
	}

	switch {
	case !g.opts.where.Empty():
		return g.insertWithWhereCond(columns, records[0])
	default:
		return g.insertNormal(columns, records...)
	}
}

func (g *Generator) insertNormal(columns []string, records ...[]interface{}) (string, []interface{}) {
	if g == nil || len(columns) == 0 || len(records) == 0 {
		return "", nil
	}

	onDuplicateKeyUpdate, updateParams := g.opts.genOnDuplicateKeyUpdate()
	sql := bytes.NewBufferString(fmt.Sprintf("INSERT INTO %s", internal.SafeName(g.table)))
	fmt.Fprintf(sql, " (%s)", strings.Join(internal.SafeNames(columns), ", "))
	fmt.Fprintf(sql, " VALUES (%s)", strings.Repeat(",?", len(records[0]))[1:])
	for i := 1; i < len(records); i++ {
		fmt.Fprintf(sql, ", (%s)", strings.Repeat(",?", len(records[i]))[1:])
	}
	sql.WriteString(sqlOrEmpty(onDuplicateKeyUpdate))

	var params []interface{}
	for _, v := range records {
		params = append(params, v...)
	}
	params = append(params, updateParams...)

	return sql.String(), params
}

func (g *Generator) insertWithWhereCond(columns []string, record []interface{}) (string, []interface{}) {
	if g == nil || len(columns) == 0 || len(record) == 0 {
		return "", nil
	}

	where, whereParams := g.opts.genWhere()
	sql := bytes.NewBufferString(fmt.Sprintf("INSERT INTO %s", internal.SafeName(g.table)))
	fmt.Fprintf(sql, " (%s)", strings.Join(internal.SafeNames(columns), ", "))
	fmt.Fprintf(sql, " SELECT %s FROM dual", strings.Repeat(",?", len(record))[1:])
	sql.WriteString(sqlOrEmpty(where))

	var params []interface{}
	params = append(params, record...)
	params = append(params, whereParams...)

	return sql.String(), params
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
		column := targetType.Field(i).Tag.Get("db")
		if column == "" || column == "-" {
			continue
		}
		column = internal.SafeName(column)

		if expr := targetType.Field(i).Tag.Get("expr"); expr != "" {
			column = expr
		}

		columns = append(columns, column)
	}

	return columns, nil
}

func sqlOrEmpty(str string) string {
	if str == "" {
		return ""
	}

	return fmt.Sprintf(" %s", str)
}
