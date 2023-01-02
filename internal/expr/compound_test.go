package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestCompound_ToSQL(t *testing.T) {
	type fields struct {
		op    internal.Operator
		exprs []internal.Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "exprs empty",
			fields: fields{
				op:    internal.OperatorAnd,
				exprs: nil,
			},
			want:  "",
			want1: nil,
		},
		{
			name: "one expr",
			fields: fields{
				op:    internal.OperatorAnd,
				exprs: []internal.Expression{NewEQ(internal.OperatorOr, "col", "val")},
			},
			want:  "AND col=?",
			want1: []interface{}{"val"},
		},
		{
			name: "and",
			fields: fields{
				op: internal.OperatorAnd,
				exprs: []internal.Expression{
					NewEQ(internal.OperatorOr, "col1", "eq"),
					NewNEQ(internal.OperatorOr, "col2", "neq"),
					NewGT(internal.OperatorOr, "col3", 1),
					NewGTE(internal.OperatorOr, "col4", 2),
					NewLT(internal.OperatorOr, "col5", 3),
					NewLTE(internal.OperatorOr, "col6", 4),
					NewBetween(internal.OperatorOr, "col7", 111, 222),
					NewIn(internal.OperatorOr, "col8", []interface{}{"in1", "in2", "in3"}),
					NewLike(internal.OperatorOr, "col9", "%%%s%%", "like"),
					NewLike(internal.OperatorOr, "col10", "%s%%", "like_prefix"),
					NewLike(internal.OperatorOr, "col11", "%%%s", "like_suffix"),
				},
			},
			want: "AND (col1=? OR col2!=? OR col3>? OR col4>=? OR col5<? OR col6<=? OR col7 BETWEEN ? AND ? OR " +
				"col8 IN (?,?,?) OR col9 LIKE ? OR col10 LIKE ? OR col11 LIKE ?)",
			want1: []interface{}{"eq", "neq", 1, 2, 3, 4, 111, 222, "in1", "in2", "in3", "%like%", "like_prefix%", "%like_suffix"},
		},
		{
			name: "or",
			fields: fields{
				op: internal.OperatorOr,
				exprs: []internal.Expression{
					NewEQ(internal.OperatorAnd, "col1", "eq"),
					NewNEQ(internal.OperatorAnd, "col2", "neq"),
					NewGT(internal.OperatorAnd, "col3", 1),
					NewGTE(internal.OperatorAnd, "col4", 2),
					NewLT(internal.OperatorAnd, "col5", 3),
					NewLTE(internal.OperatorAnd, "col6", 4),
					NewBetween(internal.OperatorAnd, "col7", 111, 222),
					NewIn(internal.OperatorAnd, "col8", []interface{}{"in1", "in2", "in3"}),
					NewLike(internal.OperatorAnd, "col9", "%%%s%%", "like"),
					NewLike(internal.OperatorAnd, "col10", "%s%%", "like_prefix"),
					NewLike(internal.OperatorAnd, "col11", "%%%s", "like_suffix"),
				},
			},
			want: "OR (col1=? AND col2!=? AND col3>? AND col4>=? AND col5<? AND col6<=? AND col7 BETWEEN ? AND ? AND " +
				"col8 IN (?,?,?) AND col9 LIKE ? AND col10 LIKE ? AND col11 LIKE ?)",
			want1: []interface{}{"eq", "neq", 1, 2, 3, 4, 111, 222, "in1", "in2", "in3", "%like%", "like_prefix%", "%like_suffix"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Compound{
				op:    tt.fields.op,
				exprs: tt.fields.exprs,
			}
			got, got1 := c.ToSQL()
			if got != tt.want {
				t.Errorf("Compound.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Compound.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
