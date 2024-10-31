package binarysearch


func BinarySearch(arr []int, target int) int {
	i := 0
	j := len(arr) - 1
	var m int

	for i <= j {
		m = (i + j) / 2
		if arr[m] < target {
			i = m + 1
		} else if arr[m] > target {
			j = m - 1
		} else {
      for m > 0 && arr[m-1] == target {
        m--
      }
			return m
		}
	}
	return -1
}
