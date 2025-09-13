package solution

import (
	"fmt"
	"testing"
)

func generateLinkedList(t testing.TB, elms []int) *ListNode {
	t.Helper()

	var root, curr, node *ListNode
	root = nil
	curr = root
	for _, elm := range elms {
		node = &ListNode{Val: elm, Next: nil}
		if root == nil {
			root = node
		} else {
			curr.Next = node
		}
		curr = node
	}

	return root
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		elms     []int
		expected bool
	}{
		{
			elms:     []int{1, 2, 2, 1},
			expected: true,
		},
		{
			elms:     []int{1, 2},
			expected: false,
		},
		{
			elms:     []int{1, 2, 1},
			expected: true,
		},
		{
			elms:     []int{1, 2, 3},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("head: %v", tc.elms), func(t *testing.T) {
			head := generateLinkedList(t, tc.elms)
			got := IsPalindrome(head)
			if got != tc.expected {
				t.Errorf("wrong answer, expected %v but got %v", tc.expected, got)
			}
		})
	}
}
