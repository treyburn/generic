package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var intTestsMin = []testCase[int]{
	{"int1", 1, 2, 1},
	{"int2", 2, 1, 1},
	{"int3", 1, 1, 1},
	{"int4", 1, -1, -1},
}

var strTestsMin = []testCase[string]{
	{"str1", "a", "b", "a"},
	{"str2", "b", "a", "a"},
	{"str3", "b", "b", "b"},
}

func TestMin(t *testing.T) {
	for _, test := range intTestsMin {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := Min(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}

	for _, test := range strTestsMin {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := Min(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}
}
