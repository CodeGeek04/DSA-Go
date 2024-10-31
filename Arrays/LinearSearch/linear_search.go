package linearsearch

func LinearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}