package lc21

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var i, j, m, curr *ListNode

	i = list1
	j = list2

	insert := func(n *ListNode) {
		if m == nil {
			m = n
		} else {
			curr.Next = n
		}
		curr = n
	}

	for i != nil && j != nil {
		if i.Val < j.Val {
			insert(i)
			i = i.Next
		} else {
			insert(j)
			j = j.Next
		}
	}

	for i != nil {
		insert(i)
		i = i.Next
	}
	for j != nil {
		insert(j)
		j = j.Next
	}

	return m
}
