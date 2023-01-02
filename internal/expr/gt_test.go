package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestGT_ToSQL(t *testing.T) {
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
			want:  "AND col>?",
			want1: []interface{}{666},
		},
		{
			name: "or",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				value:  666,
			},
			want:  "OR col>?",
			want1: []interface{}{666},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GT{
				op:     tt.fields.op,
				column: tt.fields.column,
				value:  tt.fields.value,
			}
			got, got1 := g.ToSQL()
			if got != tt.want {
				t.Errorf("GT.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GT.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
