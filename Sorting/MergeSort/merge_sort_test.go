package mergesort

// Arrays/LinearSearch/linear_search_test.go

import "testing"

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "already sorted",
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "not sorted",
			arr:      []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "some elements are same",
			arr:      []int{3, 1, 5, 1, 5, 7, 12},
			expected: []int{1, 1, 3, 5, 5, 7, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(&tt.arr)
			if !equal(tt.arr, tt.expected) {
				t.Errorf("MergeSort() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}

func equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
