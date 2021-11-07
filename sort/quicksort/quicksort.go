package quicksort

func QuickSort(source []int) []int {
	if len(source) < 2 {
		return source
	}

	pivot := source[0]
	var left, right []int
	for _, value := range source[1:] {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	mid := append([]int{pivot}, QuickSort(right)...)
	return append(QuickSort(left), mid...)
}
