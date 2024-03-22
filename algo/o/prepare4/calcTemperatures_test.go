package prepare4

import (
	"reflect"
	"testing"
)

func Test_calcTemperatures(t *testing.T) {
	type args struct {
		t []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{[]int{73, 74, 75, 71, 69, 72, 76, 73}}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTemperatures(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
