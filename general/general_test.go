package general

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		input  [2]int
		output int
	}{
		{
			input:  [2]int{3, 6},
			output: 3,
		},
		{
			input:  [2]int{6, 8},
			output: 2,
		},
	}

	for _, test := range tests {
		if got := GCD(test.input[0], test.input[1]); got != test.output {
			t.Errorf("GCD(%d, %d)=%d, expected %d", test.input[0], test.input[1], got, test.output)
		}
	}

}
