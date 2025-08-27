package leetcode

import (
	"slices"
	"testing"
)

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

func TestAddBinary(t *testing.T) {
	tests := []struct {
		input  [2]string
		output string
	}{
		{
			input:  [2]string{"11", "1"},
			output: "100",
		},
		{
			input:  [2]string{"1010", "1011"},
			output: "10101",
		},
		{
			input:  [2]string{"110", "111"},
			output: "1101",
		},
	}

	for _, test := range tests {
		if got := AddBinary(test.input[0], test.input[1]); got != test.output {
			t.Errorf("AddBinary(%v)=%v, expected %v", test.input, got, test.output)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "A man, a plan, a canal: Panama",
			output: true,
		},
		{
			input:  "race a car",
			output: false,
		},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.output {
			t.Errorf("IsPalindrome(%v)=%v, expected %v", test.input, got, test.output)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		output []int
	}{
		{
			array:  []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			array:  []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			array:  []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, test := range tests {
		if got := TwoSum(test.array, test.target); !slices.Equal(got, test.output) {
			t.Errorf("TwoSum(%v, %d)=%v, expected %v", test.array, test.target, got, test.output)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		output int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			output: 0,
		},
	}

	for _, test := range tests {
		if got := MaxProfit(test.prices); got != test.output {
			t.Errorf("MaxProfit(%v)=%d, expected %d", test.prices, got, test.output)
		}
	}
}

func TestSumRange(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	tests := []struct {
		inputs [2]int
		output int
	}{
		{
			inputs: [2]int{0, 2},
			output: 1,
		},
		{
			inputs: [2]int{2, 5},
			output: -1,
		},
		{
			inputs: [2]int{0, 5},
			output: -3,
		},
		{
			inputs: [2]int{0, 0},
			output: -2,
		},
		{
			inputs: [2]int{3, 4},
			output: -3,
		},
	}

	for _, test := range tests {
		if got := numArray.SumRange(test.inputs[0], test.inputs[1]); got != test.output {
			t.Errorf("MaxProfit(%d, %d)=%d, expected %d", test.inputs[0], test.inputs[1], got, test.output)
		}
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 7, 3, 6, 5, 6},
			output: 3,
		},
		{
			input:  []int{2, 1, -1},
			output: 0,
		},
		{
			input:  []int{-1, 1, 2},
			output: 2,
		},
		{
			input:  []int{0, 0},
			output: 0,
		},
	}

	for _, test := range tests {
		if got := PivotIndex(test.input); got != test.output {
			t.Errorf("PivotIndex(%v)=%d, expected %d", test.input, got, test.output)
		}
	}
}
