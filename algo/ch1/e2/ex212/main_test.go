package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Case 1: empty array", []int{}, []int{}},
		{"Case 2: single element", []int{1}, []int{1}},
		{"Case 3: two elements in order", []int{1, 2}, []int{2, 1}},
		{"Case 4: two elements out of order", []int{2, 1}, []int{2, 1}},
		{"Case 5: multiple elements in order", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"Case 6: multiple elements out of order", []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{"Case 7: multiple elements random order", []int{3, 5, 1, 2, 4}, []int{5, 4, 3, 2, 1}},
		{"Case 8: duplicates", []int{1, 2, 2, 3, 1}, []int{3, 2, 2, 1, 1}},
		//{"Case 9: negative integers", []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1}},
		//{"Case 10: negative and positive", []int{-3, 0, 3, -2, 2}, []int{-3, -2, 0, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			insertionSort(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("For %v, expected %v, but got %v", tc.input, tc.expected, tc.input)
			}
		})
	}
}
