package intersection_of_sorted_arrays

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple test 1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "Simple test 2",
			args: args{
				nums1: []int{4, 5, 9},
				nums2: []int{4, 4, 4, 8, 9, 9},
			},
			want: []int{4, 9},
		},
		{
			name: "Simple test 3",
			args: args{
				nums1: []int{1, 2, 3, 3, 4, 5, 6},
				nums2: []int{3, 3, 5},
			},
			want: []int{3, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
