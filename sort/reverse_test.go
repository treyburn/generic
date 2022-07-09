package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	type testCase struct {
		name  string
		input []int
		want  []int
	}

	var tests = []testCase{
		{"0 len", []int{}, []int{}},
		{"1 len", []int{2}, []int{2}},
		{"2 len", []int{1, 2}, []int{2, 1}},
		{"3 len", []int{1, 2, 3}, []int{3, 2, 1}},
		{"4 len", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := Reverse(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
