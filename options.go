package sqlg

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/wwwangxc/sqlg/internal"
)

// Options of SQL generator
type Options struct {
	where                *internal.Condition
	orderBy              []string
	groupBy              []string
	limit                uint32
	offset               uint32
	forceIndex           string
	onDuplicateKeyUpdate *AssExpr
	forUpdate            bool
}

func newOptions(opts ...Option) *Options {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func defaultOptions() *Options {
	return &Options{
		where:      &internal.Condition{},
		orderBy:    []string{},
		limit:      0,
		offset:     0,
		forceIndex: "",
	}
}

func (o *Options) genForceIndex() string {
	if o == nil || o.forceIndex == "" {
		return ""
	}

	return fmt.Sprintf("FORCE INDEX (%s)", internal.SafeName(o.forceIndex))
}

func (o *Options) genWhere() (string, []interface{}) {
	if o == nil || o.where.Empty() {
		return "", nil
	}

	sql, values := o.where.ToSQL()
	return fmt.Sprintf("WHERE %s", sql), values
}

func (o *Options) genGroupBy() string {
	if o == nil || len(o.groupBy) == 0 {
		return ""
	}

	return fmt.Sprintf("GROUP BY %s", strings.Join(internal.SafeNames(o.groupBy), ", "))
}

func (o *Options) genOrderBy() string {
	if o == nil || len(o.orderBy) == 0 {
		return ""
	}

	return fmt.Sprintf("ORDER BY %s", strings.Join(o.orderBy, ", "))
}

func (o *Options) genLimit() string {
	if o == nil || o.limit == 0 {
		return ""
	}

	return fmt.Sprintf("LIMIT %d", o.limit)
}

func (o *Options) genOffset() string {
	if o == nil || o.offset == 0 {
		return ""
	}

	return fmt.Sprintf("OFFSET %d", o.offset)
}

func (o *Options) genSet(assExpr *AssExpr) (string, []interface{}) {
	if o == nil || assExpr == nil {
		return "", nil
	}

	params := make([]interface{}, 0, assExpr.size())
	buffer := bytes.NewBuffer(nil)
	assExpr.each(func(column string, value interface{}) {
		fmt.Fprintf(buffer, ", %s=?", internal.SafeName(column))
		params = append(params, value)
	})

	sql := buffer.String()
	return fmt.Sprintf("SET %s", sql[strings.Index(sql, " ")+1:]), params
}

func (o *Options) genOnDuplicateKeyUpdate() (string, []interface{}) {
	if o == nil || o.onDuplicateKeyUpdate.empty() {
		return "", nil
	}

	params := make([]interface{}, 0, o.onDuplicateKeyUpdate.size())
	buffer := bytes.NewBuffer(nil)
	o.onDuplicateKeyUpdate.each(func(column string, value interface{}) {
		fmt.Fprintf(buffer, ", %s=?", internal.SafeName(column))
		params = append(params, value)
	})

	sql := buffer.String()
	return fmt.Sprintf("ON DUPLICATE KEY UPDATE %s", sql[strings.Index(sql, " ")+1:]), params
}

func (o *Options) genForUpdate() string {
	if o == nil || !o.forUpdate {
		return ""
	}

	return "FOR UPDATE"
}
