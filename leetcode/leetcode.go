package leetcode

import (
	"strings"
	"unicode"

	"github.com/Arup3201/stack"
)

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

func addBit(a, b, c rune) (rune, rune) {
	rbit, carry := '0', '0'
	switch {
	case a == '0' && b == '1' && c == '0':
		rbit = '1'
		carry = '0'
	case a == '1' && b == '0' && c == '0':
		rbit = '1'
		carry = '0'
	case a == '0' && b == '0' && c == '1':
		rbit = '1'
		carry = '0'
	case a == '0' && b == '1' && c == '1':
		rbit = '0'
		carry = '1'
	case a == '1' && b == '0' && c == '1':
		rbit = '0'
		carry = '1'
	case a == '1' && b == '1' && c == '0':
		rbit = '0'
		carry = '1'
	case a == '1' && b == '1' && c == '1':
		rbit = '1'
		carry = '1'
	}

	return rbit, carry
}

func revert(runes []rune) []rune {
	low, high := 0, len(runes)-1
	for low < high {
		runes[low], runes[high] = runes[high], runes[low]
		low++
		high--
	}
	return runes
}

func AddBinary(a, b string) string {
	ra, rb := []rune(a), []rune(b)
	carry := '0'
	result := []rune{}

	var cb rune
	i, j := len(ra)-1, len(rb)-1
	for i >= 0 && j >= 0 {
		cb, carry = addBit(ra[i], rb[j], carry)
		result = append(result, cb)
		i--
		j--
	}

	for i >= 0 {
		cb, carry = addBit(ra[i], '0', carry)
		result = append(result, cb)
		i--
	}

	for j >= 0 {
		cb, carry = addBit('0', rb[j], carry)
		result = append(result, cb)
		j--
	}

	if carry == '1' {
		result = append(result, carry)
	}

	result = revert(result)
	return string(result)
}

/*
LC - 125
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric
characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	filtered := []rune{}

	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}

	low, high := 0, len(filtered)-1
	for low < high {
		if filtered[low] != filtered[high] {
			return false
		}

		low++
		high--
	}

	return true
}

/*
LC-1

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func TwoSum(nums []int, target int) []int {
	numToIndex := map[int]int{}
	var num2, index int
	var ok bool

	for i, num := range nums {
		num2 = target - num
		index, ok = numToIndex[num2]
		if !ok {
			numToIndex[num] = i
		} else {
			return []int{index, i}
		}
	}

	return []int{}
}

/*
LC-121: Best time to buy and sell stocks

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func MaxProfit(prices []int) int {
	buy := prices[0]
	profit := 0

	for _, p := range prices {
		if buy > p {
			buy = p
		}

		profit = max(profit, p-buy)
	}

	return profit
}

/*
LC 303 - Range sum query - immutable

Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/

type NumArray struct {
	nums      []int
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, 0, len(nums))

	prefixSum = append(prefixSum, nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
	}

	return NumArray{
		nums:      nums,
		prefixSum: prefixSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}

	if left == 0 && right == 0 {
		return this.nums[left]
	}

	return this.prefixSum[right] - this.prefixSum[left-1]
}

/*
LC 724 - Find pivot index

Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
*/

func PivotIndex(nums []int) int {
	n := len(nums)
	pSum := make([]int, 0, n)

	pSum = append(pSum, nums[0])
	for i := 1; i < n; i++ {
		pSum = append(pSum, pSum[i-1]+nums[i])
	}

	if len(pSum) == 1 {
		return 0
	}

	if len(pSum) > 1 && pSum[n-1]-pSum[0] == 0 {
		return 0
	}

	for i := 1; i < n; i++ {
		if i == n-1 && pSum[i-1] == 0 {
			return i
		}

		if pSum[i-1] == pSum[n-1]-pSum[i] {
			return i
		}
	}

	return -1
}

/*
LC 26 - Remove duplicates from sorted array

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

func RemoveDuplicates(nums []int) int {
	k := 1
	for j := 1; j < len(nums); j++ {
		if nums[k-1] != nums[j] {
			nums[k], nums[j] = nums[j], nums[k]
			k++
		}
	}

	return k
}

/*
LC 27 - Remove element

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

func RemoveElement(nums []int, val int) int {
	k := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[k], nums[j] = nums[j], nums[k]
			k++
		}
	}

	return k
}
