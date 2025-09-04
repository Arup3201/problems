package solution

import "testing"

func TestMinimumRecolors(t *testing.T) {
	tests := [] struct {
		blocks string
		k int
		output int
	}{
		{
			blocks: "WBBWWBBWBW", 
			k: 7, 
			output: 3,
		}, 
		{
			blocks: "WBWBBBW", 
			k: 2, 
			output: 0,
		},
	}

	for _, test := range tests {
		if got:=MinimumRecolors(test.blocks, test.k); got != test.output {
			t.Errorf("FindMaxAverageSum(%v, %d)=%d, want=%d", test.blocks, test.k, got, test.output)
		}
	}
}
