package helper_test

import (
	"CYOLO/helper"
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, -5, 50, 5, 5}, []int{-5, 5, 5, 5, 50}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Sorting %v", test.input), func(t *testing.T) {
			helper.QuickSort(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, test.input)
			}
		})
	}
}
