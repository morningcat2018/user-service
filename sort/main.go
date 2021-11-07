package main

import (
	"fmt"
	s "sort1/quicksort"
)

func main() {
	arr := []int{34, 23, 6, 45, 67, 23, 124, 56, 9, 76, 23, 43, 9, 136, 765, 23, 4, 1, 98}
	fmt.Println(arr)
	sortArr := s.QuickSort(arr)
	fmt.Println(sortArr)
}
