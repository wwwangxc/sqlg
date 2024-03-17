package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestLTE_ToSQL(t *testing.T) {
	type fields struct {
		op     internal.Operator
		column string
		value  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "and",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				value:  666,
			},
			want:  "AND `col`<=?",
			want1: []interface{}{666},
		},
		{
			name: "or",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				value:  666,
			},
			want:  "OR `col`<=?",
			want1: []interface{}{666},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LTE{
				op:     tt.fields.op,
				column: tt.fields.column,
				value:  tt.fields.value,
			}
			got, got1 := l.ToSQL()
			if got != tt.want {
				t.Errorf("LTE.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LTE.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
