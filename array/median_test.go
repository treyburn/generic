package array

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"testing"
)

func TestMedian(t *testing.T) {
	type testCase[T constraints.Integer | constraints.Float] struct {
		name string
		nums []T
		want float64
	}

	var intTests = []testCase[int]{
		{"even", []int{1, 2, 3, 4}, 2.5},
		{"odd", []int{1, 2, 3}, 2.0},
		{"one element", []int{2}, 2.0},
	}

	for _, test := range intTests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := Median(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}

}
