package lc

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("search %v", tc.matrix), func(t *testing.T) {
			if got := SearchMatrix(tc.matrix, tc.target); got != tc.expected {
				t.Errorf("expected %t but got %t", tc.expected, got)
			}
		})
	}
}
