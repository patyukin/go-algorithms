package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple test",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "Simple test 2",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "Simple test 4",
			args: args{s: " "},
			want: true,
		},
		{
			name: "Simple test 5",
			args: args{s: "0P"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}

			if got := isPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
