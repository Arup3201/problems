package solution

func MinimumRecolors(blocks string, k int) int {
	minOps, ops := 0, 0
	for i := range k {
		if blocks[i] == 'W' {
			ops ++
		}
	}

	minOps = ops
	for i:=1;i<=len(blocks)-k;i++ {
		if blocks[i-1] == 'W' {
			ops--
		}
		if blocks[i+k-1] == 'W' {
			ops++
		}

		minOps = min(minOps, ops)
	}

	return minOps
}
