package lc21

import (
	"fmt"
	"slices"
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

func fromLinkedList(t testing.TB, head *ListNode) []int {
	t.Helper()

	elms := []int{}
	curr := head
	for curr != nil {
		elms = append(elms, curr.Val)
		curr = curr.Next
	}

	return elms
}

func assert(t testing.TB, got, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("wrong answer, expected %v but got %v", want, got)
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			list1:    []int{1, 3},
			list2:    []int{2, 3},
			expected: []int{1, 2, 3, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v+%v", tc.list1, tc.list2), func(t *testing.T) {
			l1Head := generateLinkedList(t, tc.list1)
			l2Head := generateLinkedList(t, tc.list2)

			m := MergeTwoLists(l1Head, l2Head)
			assert(t, fromLinkedList(t, m), tc.expected)
		})
	}
}
