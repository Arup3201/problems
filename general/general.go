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

func reverse(arr []int, l, h int) []int {
	for l < h {
		arr[l], arr[h] = arr[h], arr[l]
		l++
		h--
	}

	return arr
}

func RotateByK(arr []int, k int) []int {
	arr = reverse(arr, 0, k%len(arr)-1)
	arr = reverse(arr, k%len(arr), len(arr)-1)
	arr = reverse(arr, 0, len(arr)-1)
	return arr
}
