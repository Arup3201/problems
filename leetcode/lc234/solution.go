package solution

/*
LC 234: Palindrome linked list

Given the head of a singly linked list, return true if it is a or false otherwise.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type element struct {
	data int
	next *element
}

type stack struct {
	top *element
}

func (s *stack) push(d int) {
	elm := &element{data: d, next: nil}

	if s.top == nil {
		s.top = elm
		return
	}

	elm.next = s.top
	s.top = elm
}

func (s *stack) pop() (int, bool) {
	if s.top == nil {
		return -1, false
	}

	elm := s.top

	s.top = elm.next
	elm.next = nil

	return elm.data, true
}

func IsPalindrome(head *ListNode) bool {
	s := &stack{}

	curr := head
	for curr != nil {
		s.push(curr.Val)
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		d, _ := s.pop()
		if d != curr.Val {
			return false
		}
		curr = curr.Next
	}

	return true
}
