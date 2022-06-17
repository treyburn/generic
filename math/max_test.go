package math

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"testing"
)

type testCase[T constraints.Ordered] struct {
	name string
	x    T
	y    T
	want T
}

var intTestsMax = []testCase[int]{
	{"int1", 1, 2, 2},
	{"int2", 2, 1, 2},
	{"int3", 1, 1, 1},
	{"int4", 1, -1, 1},
}

var strTestsMax = []testCase[string]{
	{"str1", "a", "b", "b"},
	{"str2", "b", "a", "b"},
	{"str3", "b", "b", "b"},
}

func TestMax(t *testing.T) {
	for _, test := range intTestsMax {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := Max(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}

	for _, test := range strTestsMax {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := Max(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}
}
