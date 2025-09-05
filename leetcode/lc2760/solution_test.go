package solution

import "testing"

func TestLongestAlternatingSubstring(t *testing.T) {
	tests := []struct {
		nums      []int
		threshold int
		output    int
	}{
		{
			nums:      []int{3, 2, 5, 4},
			threshold: 5,
			output:    3,
		},
		{
			nums:      []int{1, 2},
			threshold: 2,
			output:    1,
		},
		{
			nums:      []int{3, 9},
			threshold: 4,
			output:    0,
		},
		{
			nums:      []int{2, 3, 4, 5, 6},
			threshold: 6,
			output:    5,
		},
		{
			nums:      []int{2, 10, 5},
			threshold: 7,
			output:    1,
		},
	}

	for _, test := range tests {
		if got := LongestAlternatingSubarray(test.nums, test.threshold); got != test.output {
			t.Errorf("LongestAlternatingSubarray(%v, %d)=%d, want=%d", test.nums, test.threshold, got, test.output)
		}
	}
}
