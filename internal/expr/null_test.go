package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestNull_ToSQL(t *testing.T) {
	type fields struct {
		op     internal.Operator
		column string
		isNot  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "and null",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				isNot:  false,
			},
			want:  "AND col IS NULL",
			want1: nil,
		},
		{
			name: "or null",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				isNot:  false,
			},
			want:  "OR col IS NULL",
			want1: nil,
		},
		{
			name: "and not null",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				isNot:  true,
			},
			want:  "AND col IS NOT NULL",
			want1: nil,
		},
		{
			name: "or not null",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				isNot:  true,
			},
			want:  "OR col IS NOT NULL",
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Null{
				op:     tt.fields.op,
				column: tt.fields.column,
				isNot:  tt.fields.isNot,
			}
			got, got1 := n.ToSQL()
			if got != tt.want {
				t.Errorf("Null.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Null.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
