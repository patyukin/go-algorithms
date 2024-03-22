package main

import (
	"reflect"
	"runtime"
	"testing"
)

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "big test",
		args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
							Left: &TreeNode{
								Val: 16,
							},
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 10,
						},
						Right: &TreeNode{
							Val: 11,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 12,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 14,
						},
					},
				},
			},
		},
		want: []int{1, 3, 7, 14},
	},
	{
		name: "Simple test",
		args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}},
		want: []int{1, 3, 4},
	},
	{
		name: "Empty tree",
		args: args{root: nil},
		want: []int{},
	},
	{
		name: "One element tree",
		args: args{root: &TreeNode{Val: 1}},
		want: []int{1},
	},
	{
		name: "Two elements tree",
		args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
		want: []int{1, 2},
	},
}

func Test_rightSideView(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
			if got := rightSideView1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRightSideView(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rightSideView(tests[0].args.root)
	}
}

func BenchmarkRightSideView1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rightSideView1(tests[0].args.root)
	}
}

func TestMyFunctionAllocs(t *testing.T) {
	runtime.GC()

	allocations := testing.AllocsPerRun(1000, func() {
		rightSideView(tests[0].args.root)
	})

	t.Logf("Allocations: %.2f per run", allocations)
}

func TestMyFunctionAllocs1(t *testing.T) {
	runtime.GC()

	allocations := testing.AllocsPerRun(1000, func() {
		rightSideView1(tests[0].args.root)
	})

	t.Logf("Allocations: %.2f per run", allocations)
}
