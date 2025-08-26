package leetcode

import "github.com/Arup3201/stack"

func FindCommonPrefix(strs []string) string {
	prefix := strs[0]
	prefixLength := len(strs[0])

	for i := range strs {
		if prefixLength > len(strs[i]) {
			prefixLength = len(strs[i])
		}
	}

	found := false
	for prefix != "" && !found {
		found = true
		for i := range strs {
			if prefix != strs[i][:prefixLength] {
				prefixLength -= 1
				prefix = prefix[:prefixLength]
				found = false
			}
		}
	}

	return prefix
}

func ValidParenthesis(s string) bool {
	st := stack.NewStack(10000)

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			st.Push(int(ch))
		} else {
			cl, ok := st.Pop()

			if !ok {
				return false
			}

			switch {
			case cl == '(' && ch == ')':
			case cl == '{' && ch == '}':
			case cl == '[' && ch == ']':
				continue
			default:
				return false
			}
		}
	}

	if st.Top != 0 {
		return false
	}

	return true
}
