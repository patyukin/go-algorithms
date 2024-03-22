package main

import "testing"

func Test_compressedUnicode(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"AAAAABBBCCCDDDEEAA"}, "A5B3C3D3E2A2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressedUnicode(tt.args.s); got != tt.want {
				t.Errorf("compressedUnicode() = %v, want %v", got, tt.want)
			}

			if got := compressed(tt.args.s); got != tt.want {
				t.Errorf("compressed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressModify(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"AAAAABBBCCCDDDEEAA"}, "A5B3C3D3E2A2"},
		{"test 2", args{"AAAABBBCCDDEE"}, "A4B3C2D2E2"},
		{"test 3", args{"AAAABBBCCDDEER"}, "A4B3C2D2E2R"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressModify(tt.args.s); got != tt.want {
				t.Errorf("compressModify() = %v, want %v", got, tt.want)
			}
		})
	}
}
