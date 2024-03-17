package expr

import (
	"reflect"
	"testing"

	"github.com/wwwangxc/sqlg/internal"
)

func TestLike_ToSQL(t *testing.T) {
	type fields struct {
		op     internal.Operator
		column string
		format string
		value  interface{}
		isNot  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name: "and like",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%%%s%%",
				value:  "val",
				isNot:  false,
			},
			want:  "AND `col` LIKE ?",
			want1: []interface{}{"%val%"},
		},
		{
			name: "or like",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%%%s%%",
				value:  "val",
				isNot:  false,
			},
			want:  "OR `col` LIKE ?",
			want1: []interface{}{"%val%"},
		},
		{
			name: "and like prefix",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%s%%",
				value:  "val",
				isNot:  false,
			},
			want:  "AND `col` LIKE ?",
			want1: []interface{}{"val%"},
		},
		{
			name: "or like prefix",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%s%%",
				value:  "val",
				isNot:  false,
			},
			want:  "OR `col` LIKE ?",
			want1: []interface{}{"val%"},
		},
		{
			name: "and like suffix",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%%%s",
				value:  "val",
				isNot:  false,
			},
			want:  "AND `col` LIKE ?",
			want1: []interface{}{"%val"},
		},
		{
			name: "or like suffix",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%%%s",
				value:  "val",
				isNot:  false,
			},
			want:  "OR `col` LIKE ?",
			want1: []interface{}{"%val"},
		},
		{
			name: "and not like",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%%%s%%",
				value:  "val",
				isNot:  true,
			},
			want:  "AND `col` NOT LIKE ?",
			want1: []interface{}{"%val%"},
		},
		{
			name: "or not like",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%%%s%%",
				value:  "val",
				isNot:  true,
			},
			want:  "OR `col` NOT LIKE ?",
			want1: []interface{}{"%val%"},
		},
		{
			name: "and not like prefix",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%s%%",
				value:  "val",
				isNot:  true,
			},
			want:  "AND `col` NOT LIKE ?",
			want1: []interface{}{"val%"},
		},
		{
			name: "or not like prefix",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%s%%",
				value:  "val",
				isNot:  true,
			},
			want:  "OR `col` NOT LIKE ?",
			want1: []interface{}{"val%"},
		},
		{
			name: "and not like suffix",
			fields: fields{
				op:     internal.OperatorAnd,
				column: "col",
				format: "%%%s",
				value:  "val",
				isNot:  true,
			},
			want:  "AND `col` NOT LIKE ?",
			want1: []interface{}{"%val"},
		},
		{
			name: "or not like suffix",
			fields: fields{
				op:     internal.OperatorOr,
				column: "col",
				format: "%%%s",
				value:  "val",
				isNot:  true,
			},
			want:  "OR `col` NOT LIKE ?",
			want1: []interface{}{"%val"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Like{
				op:     tt.fields.op,
				column: tt.fields.column,
				format: tt.fields.format,
				value:  tt.fields.value,
				isNot:  tt.fields.isNot,
			}
			got, got1 := l.ToSQL()
			if got != tt.want {
				t.Errorf("Like.ToSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Like.ToSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
