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
