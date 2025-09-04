package lc643

import (
	"fmt"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := [] struct {
		nums []int
		k int
		output string
	}{
		{
			nums: []int{1,12,-5,-6,50,3}, 
			k: 4,
			output: "12.75", 
		}, 
		{
			nums: []int{5}, 
			k: 1, 
			output: "5.00", 
		}, 
	}

	for _, test := range tests {
		if got:=FindMaxAverageSum(test.nums, test.k); fmt.Sprintf("%.2f", got) != test.output {
			t.Errorf("FindMaxAverageSum(%v, %d)=%.2f, want=%s", test.nums, test.k, got, test.output)
		}
	}
}
