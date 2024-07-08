package ternary

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args[T any] struct {
		condition  bool
		trueValue  T
		falseValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[any]{
		{
			name: "if true",
			args: args[any]{
				condition:  true,
				trueValue:  1,
				falseValue: 0,
			},
			want: 1,
		},
		{
			name: "if false",
			args: args[any]{
				condition:  false,
				trueValue:  "1",
				falseValue: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
