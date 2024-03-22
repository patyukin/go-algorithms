package main

import (
	"reflect"
	"testing"
)

func Test_findAndDeleteFromEnd(t *testing.T) {
	type args struct {
		root                     *ListNode
		numberElementFromEndList int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test case 1",
			args: args{
				root:                     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				numberElementFromEndList: 2,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndDeleteFromEnd(tt.args.root, tt.args.numberElementFromEndList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndDeleteFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
