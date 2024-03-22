package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple test",
			args: args{
				nums: []int{1, 3, 3, 3, 3, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{3, 1},
		},
		{
			name: "Simple test",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "Simple test",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Simple test",
			args: args{
				nums: []int{3, 0, 1, 0},
				k:    1,
			},
			want: []int{0},
		},
		{
			name: "Simple test",
			args: args{
				nums: []int{3, 0, 1, 0, 0, 3},
				k:    2,
			},
			want: []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
