package solution

import "testing"

func TestMaximumStrongPairXor(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5},
			output: 7,
		},
		{
			nums:   []int{10, 100},
			output: 0,
		},
		{
			nums:   []int{5, 6, 25, 30},
			output: 7,
		},
		{
			nums:   []int{1, 1, 2, 3, 5},
			output: 6,
		},
	}

	for _, test := range tests {
		if got := MaximumStrongPairXor(test.nums); got != test.output {
			t.Errorf("LongestAlternatingSubarray(%v)=%d, want=%d", test.nums, got, test.output)
		}
	}
}
