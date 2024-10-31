package binarysearch

import (
    "testing"
    "sort"
)

func TestBinarySearch(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        target   int
        expected int
    }{
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
            expected: 1, // or could be 2 or 3, depending on implementation
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

    // Random array test
    randomArr := make([]int, 1000)
    for i := range randomArr {
        randomArr[i] = i * 2 // even numbers
    }
    tests = append(tests, struct {
        name     string
        arr      []int
        target   int
        expected int
    }{
        name:     "large sorted array - element exists",
        arr:      randomArr,
        target:   500,
        expected: 250,
    })
    
    tests = append(tests, struct {
        name     string
        arr      []int
        target   int
        expected int
    }{
        name:     "large sorted array - element doesn't exist",
        arr:      randomArr,
        target:   501,
        expected: -1,
    })

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Verify array is sorted
            if !sort.IntsAreSorted(tt.arr) {
                t.Fatalf("Input array must be sorted")
            }
            
            got := BinarySearch(tt.arr, tt.target)
            if got != tt.expected {
                t.Errorf("BinarySearch() = %v, want %v", got, tt.expected)
            }
            
            // Verify the result if element was found
            if got != -1 {
                if got >= len(tt.arr) {
                    t.Errorf("Returned index %d out of bounds for array of length %d", got, len(tt.arr))
                }
                if tt.arr[got] != tt.target {
                    t.Errorf("Array[%d] = %d, but target was %d", got, tt.arr[got], tt.target)
                }
            }
        })
    }
}
