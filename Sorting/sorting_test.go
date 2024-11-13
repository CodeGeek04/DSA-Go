package sorting

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"

	bubblesort "github.com/CodeGeek04/go-data-structures/Sorting/BubbleSort"
	mergesort "github.com/CodeGeek04/go-data-structures/Sorting/MergeSort"
)

// SortFunc is a function type that represents any sorting algorithm
type SortFunc func(*[]int)

// TestCase represents a single test case for sorting functions
type TestCase struct {
	name     string
	input    []int
	expected []int
}

// getTestCases returns a slice of test cases used for all sorting algorithms
func getTestCases() []TestCase {
	tests := []TestCase{
		// Empty and single element arrays
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element array",
			input:    []int{1},
			expected: []int{1},
		},

		// Basic cases
		{
			name:     "already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random order small array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},

		// Arrays with duplicate elements
		{
			name:     "all elements same",
			input:    []int{4, 4, 4, 4, 4},
			expected: []int{4, 4, 4, 4, 4},
		},
		{
			name:     "many duplicates",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},

		// Arrays with negative numbers
		{
			name:     "negative numbers only",
			input:    []int{-5, -3, -1, -4, -2},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-3, 1, -4, 1, -5, 9, -2, 6, -5, 3, 5},
			expected: []int{-5, -5, -4, -3, -2, 1, 1, 3, 5, 6, 9},
		},

		// Arrays with zero
		{
			name:     "array with zeros",
			input:    []int{0, 0, 1, 0, -1, 2, 0, -2},
			expected: []int{-2, -1, 0, 0, 0, 0, 1, 2},
		},

		// Large numbers
		{
			name:     "large numbers",
			input:    []int{1000000, 500000, 750000, 250000},
			expected: []int{250000, 500000, 750000, 1000000},
		},

		// Edge cases for stability testing
		{
			name:     "repeated pairs",
			input:    []int{2, 1, 2, 1, 2, 1},
			expected: []int{1, 1, 1, 2, 2, 2},
		},
	}

	// Add test cases with various array sizes
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		// Random array
		rand.Seed(time.Now().UnixNano())
		randomArr := make([]int, size)
		for i := range randomArr {
			randomArr[i] = rand.Intn(1000)
		}
		expected := make([]int, len(randomArr))
		copy(expected, randomArr)
		sort.Ints(expected)

		tests = append(tests, TestCase{
			name:     "random array of size " + strconv.Itoa(size),
			input:    randomArr,
			expected: expected,
		})

		// Nearly sorted array
		nearlySorted := make([]int, size)
		for i := range nearlySorted {
			nearlySorted[i] = i
		}
		// Swap a few elements to make it nearly sorted
		for i := 0; i < size/10; i++ {
			j := rand.Intn(size)
			k := rand.Intn(size)
			nearlySorted[j], nearlySorted[k] = nearlySorted[k], nearlySorted[j]
		}
		expected = make([]int, len(nearlySorted))
		copy(expected, nearlySorted)
		sort.Ints(expected)

		tests = append(tests, TestCase{
			name:     "nearly sorted array of size " + strconv.Itoa(size),
			input:    nearlySorted,
			expected: expected,
		})
	}

	return tests
}

// testSortFunction runs all test cases for a given sort function
func testSortFunction(t *testing.T, sortFunc SortFunc, name string) {
	tests := getTestCases()

	for _, tt := range tests {
		t.Run(name+"/"+tt.name, func(t *testing.T) {
			// Make a copy of the input array to avoid modifying the original
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// Time the sorting operation
			start := time.Now()
			sortFunc(&input)
			duration := time.Since(start)

			// Check if the array is actually sorted
			if !sort.IntsAreSorted(input) {
				t.Errorf("%s: array is not sorted", name)
			}

			// Check if the sorted array matches the expected result
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("%s = %v, want %v", name, input, tt.expected)
			}

			// Additional checks
			if len(input) != len(tt.input) {
				t.Errorf("Length changed during sort: got %v, want %v", len(input), len(tt.input))
			}

			// Report timing for larger arrays
			if len(input) >= 100 {
				t.Logf("%s/%s took %v", name, tt.name, duration)
			}

			// Check if all original elements are present (no losses during sorting)
			inputMap := make(map[int]int)
			for _, v := range tt.input {
				inputMap[v]++
			}
			for _, v := range input {
				inputMap[v]--
				if inputMap[v] < 0 {
					t.Errorf("Element count mismatch: too many %d in result", v)
				}
			}
			for k, v := range inputMap {
				if v > 0 {
					t.Errorf("Element count mismatch: missing %d in result", k)
				}
			}
		})
	}
}

// TestAllSortFunctions runs all sorting algorithm tests
func TestAllSortFunctions(t *testing.T) {
	// Test Bubble Sort
	t.Run("BubbleSort", func(t *testing.T) {
		testSortFunction(t, bubblesort.BubbleSort, "BubbleSort")
	})

	// Add other sorting algorithms here as needed, for example:
	// t.Run("QuickSort", func(t *testing.T) {
	//     testSortFunction(t, QuickSort.QuickSort, "QuickSort")
	// })
	t.Run("MergeSort", func(t *testing.T) {
		testSortFunction(t, mergesort.MergeSort, "MergeSort")
	})
}
