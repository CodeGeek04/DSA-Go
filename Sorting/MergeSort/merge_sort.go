package mergesort

// MergeSort sorts an array using the merge sort algorithm
func MergeSort(pArr *[]int) {
	arr := *pArr
	if len(arr) <= 1 {
		return
	}

	var mergeSort func([]int, int, int)
	var merge func([]int, int, int, int)

	// merge combines two sorted subarrays into one sorted array
	merge = func(arr []int, low, mid, high int) {
		// Create temporary arrays
		leftSize := mid - low + 1
		rightSize := high - mid

		left := make([]int, leftSize)
		right := make([]int, rightSize)

		// Copy data to temporary arrays
		for i := 0; i < leftSize; i++ {
			left[i] = arr[low+i]
		}
		for i := 0; i < rightSize; i++ {
			right[i] = arr[mid+1+i]
		}

		// Merge the temporary arrays back into arr[low..high]
		i := 0   // Initial index of left subarray
		j := 0   // Initial index of right subarray
		k := low // Initial index of merged subarray

		// Compare and merge elements
		for i < leftSize && j < rightSize {
			if left[i] <= right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
			k++
		}

		// Copy remaining elements of left[], if any
		for i < leftSize {
			arr[k] = left[i]
			i++
			k++
		}

		// Copy remaining elements of right[], if any
		for j < rightSize {
			arr[k] = right[j]
			j++
			k++
		}
	}

	// mergeSort recursively divides the array and merges sorted subarrays
	mergeSort = func(arr []int, low, high int) {
		if low < high {
			// Find the middle point
			mid := low + (high-low)/2

			// Sort first and second halves
			mergeSort(arr, low, mid)
			mergeSort(arr, mid+1, high)

			// Merge the sorted halves
			merge(arr, low, mid, high)
		}
	}

	// Start the merge sort
	mergeSort(arr, 0, len(arr)-1)
}
