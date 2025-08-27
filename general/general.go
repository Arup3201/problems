package general

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func ReverseDigits(num int) int {
	pos := 10
	rev := 0
	for num > 0 {
		rev = rev*pos + (num % 10)

		num /= 10
	}

	return rev
}
