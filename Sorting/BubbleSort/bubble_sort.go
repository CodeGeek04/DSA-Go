package bubblesort

func BubbleSort(arrP *[]int) {
	sorted := false
	arr := *arrP
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}
