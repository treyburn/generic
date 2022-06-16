package generic

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}
