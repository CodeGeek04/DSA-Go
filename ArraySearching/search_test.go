package searching

import (
	"sort"
	"testing"

	binarysearch "github.com/CodeGeek04/go-data-structures/ArraySearching/BinarySearch"
	linearsearch "github.com/CodeGeek04/go-data-structures/ArraySearching/LinearSearch"
)

// SearchFunc is a function type that represents any search algorithm
type SearchFunc func([]int, int) int

// TestCase represents a single test case for search functions
type TestCase struct {
	name     string
	arr      []int
	target   int
	expected int
}

// getTestCases returns a slice of test cases used for all search algorithms
func getTestCases() []TestCase {
	tests := []TestCase{
		// Basic cases
		{
			name:     "element found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "element found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "element found in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},

		// Edge cases
		{
			name:     "empty array",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "single element array - element found",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element array - element not found",
			arr:      []int{1},
			target:   2,
			expected: -1,
		},

		// Duplicate elements
		{
			name:     "duplicate elements - first occurrence",
			arr:      []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 1,
		},

		// Large numbers
		{
			name:     "large numbers",
			arr:      []int{1000000, 2000000, 3000000},
			target:   2000000,
			expected: 1,
		},

		// Negative numbers
		{
			name:     "negative numbers",
			arr:      []int{-5, -4, -3, -2, -1},
			target:   -3,
			expected: 2,
		},
		{
			name:     "mixed positive and negative",
			arr:      []int{-2, -1, 0, 1, 2},
			target:   0,
			expected: 2,
		},

		// Not found cases
		{
			name:     "element smaller than all array elements",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "element larger than all array elements",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "element between existing elements",
			arr:      []int{1, 3, 5, 7},
			target:   4,
			expected: -1,
		},
	}

	// Add random array test
	randomArr := make([]int, 1000)
	for i := range randomArr {
		randomArr[i] = i * 2 // even numbers
	}
	tests = append(tests,
		TestCase{
			name:     "large sorted array - element exists",
			arr:      randomArr,
			target:   500,
			expected: 250,
		},
		TestCase{
			name:     "large sorted array - element doesn't exist",
			arr:      randomArr,
			target:   501,
			expected: -1,
		},
	)

	return tests
}

// testSearchFunction runs all test cases for a given search function
func testSearchFunction(t *testing.T, searchFunc SearchFunc, name string, requireSorted bool) {
	tests := getTestCases()

	for _, tt := range tests {
		t.Run(name+"/"+tt.name, func(t *testing.T) {
			// Make a copy of the array to avoid modifying the original
			arr := make([]int, len(tt.arr))
			copy(arr, tt.arr)

			// If the search requires sorted array, sort it
			if requireSorted {
				if !sort.IntsAreSorted(arr) {
					sort.Ints(arr)
				}
			}

			got := searchFunc(arr, tt.target)

			// For linear search, we need to adjust the expected index if the array was sorted
			expected := tt.expected
			if requireSorted && expected != -1 {
				// Find the actual expected index in the sorted array
				for i, v := range arr {
					if v == tt.target {
						expected = i
						break
					}
				}
			}

			if got != expected {
				t.Errorf("%s() = %v, want %v", name, got, expected)
			}

			// Verify the result if element was found
			if got != -1 {
				if got >= len(arr) {
					t.Errorf("Returned index %d out of bounds for array of length %d", got, len(arr))
				}
				if arr[got] != tt.target {
					t.Errorf("Array[%d] = %d, but target was %d", got, arr[got], tt.target)
				}
			}
		})
	}
}

func TestAllSearchFunctions(t *testing.T) {
	// Test Linear Search
	t.Run("LinearSearch", func(t *testing.T) {
		testSearchFunction(t, linearsearch.LinearSearch, "LinearSearch", false)
	})

	// Test Binary Search
	t.Run("BinarySearch", func(t *testing.T) {
		testSearchFunction(t, binarysearch.BinarySearch, "BinarySearch", true)
	})
}
