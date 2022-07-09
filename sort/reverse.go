package sort

func Reverse[T any](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	for left < right {
		swap := arr[left]
		arr[left] = arr[right]
		arr[right] = swap
		left++
		right--
	}

	return arr
}
