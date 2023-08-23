package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSortTable(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Given an sorted array, When sorted, Then I get the same array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Given an unsorted array, When I call MergeSort, Then I get a sorted array",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Given an unsorted array, When I call MergeSort, Then I get a sorted array",
			input:    []int{10, -1, 2, 5, 0, 6, 4, -5},
			expected: []int{-5, -1, 0, 2, 4, 5, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MergeSort(tt.input))
		})
	}
}
