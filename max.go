package generic

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}
	return y
}
