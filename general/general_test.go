package general

import (
	"slices"
	"testing"
)

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

func TestReverseDigit(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			input:  123,
			output: 321,
		},
		{
			input:  481,
			output: 184,
		},
		{
			input:  666,
			output: 666,
		},
	}

	for _, test := range tests {
		if got := ReverseDigits(test.input); got != test.output {
			t.Errorf("GCD(%d)=%d, expected %d", test.input, got, test.output)
		}
	}
}

func TestRotateByK(t *testing.T) {
	tests := []struct {
		array  []int
		k      int
		output []int
	}{
		{
			array:  []int{3, 1, 5, 2, 3},
			k:      2,
			output: []int{5, 2, 3, 3, 1},
		},
		{
			array:  []int{1, 2, 3, 4, 5},
			k:      1,
			output: []int{2, 3, 4, 5, 1},
		},
		{
			array:  []int{1, 2, 3, 4, 5},
			k:      4,
			output: []int{5, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		if got := RotateByK(test.array, test.k); !slices.Equal(got, test.output) {
			t.Errorf("RotateByK(%v, %d)=%d, expected %d", test.array, test.k, got, test.output)
		}
	}
}

func TestCharCount(t *testing.T) {
	tests := []struct {
		s      string
		ch     rune
		output int
	}{
		{
			s:      "Hello, world",
			ch:     'o',
			output: 2,
		},
		{
			s:      "Lord of mysteries",
			ch:     'r',
			output: 2,
		},
		{
			s:      "One piece",
			ch:     'r',
			output: 0,
		},
	}

	for _, test := range tests {
		if got := CharCount(test.s, test.ch); got != test.output {
			t.Errorf("CharCount(%s, %c)=%d, expected %d", test.s, test.ch, got, test.output)
		}
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{2, 4, 1, 5, 3, 10},
			expected: []int{1, 2, 3, 4, 5, 10},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums:     []int{6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			nums:     []int{2},
			expected: []int{2},
		},
	}

	for _, tc := range testCases {
		if got := quickSortStart(tc.nums); !slices.Equal(got, tc.expected) {
			t.Errorf("quickSort(%v)=%v, expected %v", tc.nums, got, tc.expected)
		}
	}
}
