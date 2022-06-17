package array

import "golang.org/x/exp/constraints"

// Median returns the median value of a sorted array with at least 1 element
func Median[T constraints.Integer | constraints.Float](arr []T) float64 {
	l := len(arr)
	if l < 1 {
		panic("cannot get median of empty array")
	}

	switch l % 2 {
	default: // array has an odd number of elements
		return float64(arr[l/2])
	case 0: // array has an even number of elements
		floor := arr[l/2]
		ceil := arr[l/2-1]
		return float64(floor+ceil) / 2
	}
}
