package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestBetween_ToSQL(t *testing.T) {
	type fields struct {
		op     internal.Operator
		column string
		value1 interface{}
		value2 interface{}
		isNot  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "and between",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				value1: 100,
				value2: 200,
				isNot:  false,
			},
			want:  "AND `col` BETWEEN ? AND ?",
			want1: []interface{}{100, 200},
		},
		{
			name: "or between",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				value1: 100,
				value2: 200,
				isNot:  false,
			},
			want:  "OR `col` BETWEEN ? AND ?",
			want1: []interface{}{100, 200},
		},
		{
			name: "and not between",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				value1: 100,
				value2: 200,
				isNot:  true,
			},
			want:  "AND `col` NOT BETWEEN ? AND ?",
			want1: []interface{}{100, 200},
		},
		{
			name: "or not between",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				value1: 100,
				value2: 200,
				isNot:  true,
			},
			want:  "OR `col` NOT BETWEEN ? AND ?",
			want1: []interface{}{100, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Between{
				op:     tt.fields.op,
				column: tt.fields.column,
				value1: tt.fields.value1,
				value2: tt.fields.value2,
				isNot:  tt.fields.isNot,
			}
			got, got1 := b.ToSQL()
			if got != tt.want {
				t.Errorf("Between.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Between.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
