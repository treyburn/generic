package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	var lesser []T
	var greater []T

	pivot := arr[0]
	for _, val := range arr[1:] {
		if val >= pivot {
			greater = append(greater, val)
		} else {
			lesser = append(lesser, val)
		}
	}

	return append(append(QuickSort(lesser), pivot), QuickSort(greater)...)
}
