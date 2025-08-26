package leetcode

import "testing"

func TestFindCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"flow", "flower", "flight"},
			output: "fl",
		},
		{
			input:  []string{"aab", "ab", "abc"},
			output: "a",
		},
		{
			input:  []string{"bde", "aba", "abc"},
			output: "",
		},
		{
			input:  []string{"", "aba", "abc"},
			output: "",
		},
	}

	for _, test := range tests {
		if got := FindCommonPrefix(test.input); got != test.output {
			t.Errorf("FindCommonPrefix(%v)=%v, expected %v", test.input, got, test.output)
		}
	}
}

func TestValidParenthesis(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "(){}[()]",
			output: true,
		},
		{
			input:  "()",
			output: true,
		},
		{
			input:  "(){]",
			output: false,
		},
		{
			input:  "[(])",
			output: false,
		},
	}

	for _, test := range tests {
		if got := ValidParenthesis(test.input); got != test.output {
			t.Errorf("ValidParenthesis(%v)=%v, expected %v", test.input, got, test.output)
		}
	}
}
