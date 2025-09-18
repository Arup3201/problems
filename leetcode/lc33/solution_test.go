package lc

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{2, 3, 0, 1},
			target:   3,
			expected: 1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2, 3},
			target:   2,
			expected: 6,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("search %v", tc.nums), func(t *testing.T) {
			if got := Search(tc.nums, tc.target); got != tc.expected {
				t.Errorf("expected %d but got %d", tc.expected, got)
			}
		})
	}
}
