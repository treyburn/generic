package array

import "golang.org/x/exp/constraints"

// MergeSorted assumes that both inputs are pre-sorted
func MergeSorted[T constraints.Ordered](arr1, arr2 []T) []T {
	n1 := len(arr1)
	n2 := len(arr2)
	arr3 := make([]T, n1+n2)
	var arr1idx, arr2idx, arr3idx int

	// traverse both arrays
	for arr1idx < n1 && arr2idx < n2 {
		// Check if current element of first array is smaller than current element of second array.
		// If yes, store first array element and increment first array index.
		//Otherwise, do same with second array
		switch arr1[arr1idx] < arr2[arr2idx] {
		case true:
			arr3[arr3idx] = arr1[arr1idx]
			arr3idx++
			arr1idx++
		case false:
			arr3[arr3idx] = arr2[arr2idx]
			arr3idx++
			arr2idx++
		}
	}

	// drain remaining elements of arr1 - if any
	for arr1idx < n1 {
		arr3[arr3idx] = arr1[arr1idx]
		arr3idx++
		arr1idx++
	}

	//drain remaining elements of arr2 - if any
	for arr2idx < n2 {
		arr3[arr3idx] = arr2[arr2idx]
		arr3idx++
		arr2idx++
	}

	return arr3
}
