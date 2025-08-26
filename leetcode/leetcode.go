package leetcode

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
