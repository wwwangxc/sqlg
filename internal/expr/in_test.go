package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestIn_ToSQL(t *testing.T) {
	type fields struct {
		op     internal.Operator
		column string
		values []interface{}
		isNot  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "and in",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				values: []interface{}{"val1", "val2", "val3"},
				isNot:  false,
			},
			want:  "AND `col` IN (?,?,?)",
			want1: []interface{}{"val1", "val2", "val3"},
		},
		{
			name: "or in",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				values: []interface{}{"val1", "val2", "val3"},
				isNot:  false,
			},
			want:  "OR `col` IN (?,?,?)",
			want1: []interface{}{"val1", "val2", "val3"},
		},
		{
			name: "and not in",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				values: []interface{}{"val1", "val2", "val3"},
				isNot:  true,
			},
			want:  "AND `col` NOT IN (?,?,?)",
			want1: []interface{}{"val1", "val2", "val3"},
		},
		{
			name: "or not in",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				values: []interface{}{"val1", "val2", "val3"},
				isNot:  true,
			},
			want:  "OR `col` NOT IN (?,?,?)",
			want1: []interface{}{"val1", "val2", "val3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &In{
				op:     tt.fields.op,
				column: tt.fields.column,
				values: tt.fields.values,
				isNot:  tt.fields.isNot,
			}
			got, got1 := i.ToSQL()
			if got != tt.want {
				t.Errorf("In.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("In.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
