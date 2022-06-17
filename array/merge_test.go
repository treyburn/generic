package array

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"testing"
)

func Test_merge(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}

	var intTests = []testCase[int]{
		{"even", []int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{"odd", []int{1, 3}, []int{2}, []int{1, 2, 3}},
		{"first empty", []int{}, []int{2}, []int{2}},
		{"second empty", []int{2}, []int{}, []int{2}},
		{"both empty", []int{}, []int{}, []int{}},
	}

	for _, test := range intTests {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := MergeSorted(tc.nums1, tc.nums2)
			assert.Equal(t, tc.want, got)
		})
	}
}
