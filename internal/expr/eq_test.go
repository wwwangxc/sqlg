package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestEQ_ToSQL(t *testing.T) {
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
				value:  "val",
			},
			want:  "AND col=?",
			want1: []interface{}{"val"},
		},
		{
			name: "or",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				value:  "val",
			},
			want:  "OR col=?",
			want1: []interface{}{"val"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EQ{
				op:     tt.fields.op,
				column: tt.fields.column,
				value:  tt.fields.value,
			}
			got, got1 := e.ToSQL()
			if got != tt.want {
				t.Errorf("EQ.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EQ.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
