package binarysearch

import (
	"fmt"
)

func BinarySearch(arr []int, target int) int {
	i := 0
	j := len(arr) - 1
	var m int

	for i < j {
		m = (i + j) / 2
		if arr[m] < target {
			i = m
		} else if arr[m] > target {
			j = m
		} else {
			return m
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 4))
}
