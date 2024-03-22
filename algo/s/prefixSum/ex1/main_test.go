package ex1

import (
	"reflect"
	"testing"
)

func Test_prefixSum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{"test 1", "12345", []int{1, 3, 6, 10, 15}},
		{"test 2", "99999", []int{9, 18, 27, 36, 45}},
		{"test 3", "54321", []int{5, 9, 12, 14, 15}},
		{"test 4", "00000", []int{0, 0, 0, 0, 0}},
		{"test 5", "98765", []int{9, 17, 24, 30, 35}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixSum(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("prefixSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkPrefixSums1(b *testing.B) {
	s := "1234567890"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixSum(s)
	}
}

func BenchmarkPrefixSums2(b *testing.B) {
	s := "9876543210"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixSum(s)
	}
}

func BenchmarkPrefixSums3(b *testing.B) {
	s := "5555555555"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixSum(s)
	}
}
