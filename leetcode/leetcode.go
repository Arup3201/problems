package leetcode

import (
	"slices"
	"sort"
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

/*
LC 283 - Move zeros

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/

func MoveZeros(nums []int) {
	k := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[k], nums[j] = nums[j], nums[k]
			k++
		}
	}
}

/*
LC 141 - Linked list cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for slow != nil || fast != nil {
		if slow.Next == nil {
			break
		}
		if fast.Next == nil {
			break
		}
		if fast.Next.Next == nil {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

/*
LC 387 - First Unique Character in a String

Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
*/

func FirstUniqueChar(s string) int {
	freq := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, ok := freq[s[i]]
		if !ok {
			freq[s[i]] = 0
		} else {
			freq[s[i]]++
		}
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]] == 1 {
			return i
		}
	}

	return -1
}

/*
LC 242 - Valid anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.
*/

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqS, freqT := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		freqS[s[i]]++
		freqT[t[i]]++
	}

	for bt, fs := range freqS {
		if freqT[bt] != fs {
			return false
		}
	}

	return true
}

/*
LC 49 - Group anagrams

Given an array of strings strs, group the anagrams together. You can return the answer in any order.
*/

func getSortedWord(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func GroupAnagram(strs []string) [][]string {
	sortedFreqs := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		sorted := getSortedWord(strs[i])
		sortedFreqs[sorted] = append(sortedFreqs[sorted], strs[i])
	}

	groupedAnagrams := [][]string{}
	for _, anagrams := range sortedFreqs {
		groupedAnagrams = append(groupedAnagrams, anagrams)
	}

	return groupedAnagrams
}

/*
LC 268 - Missing numbers

Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
*/

func MissingNumber(nums []int) int {
	sum := 0
	mx := len(nums)
	for _, n := range nums {
		sum += n
	}

	return mx*(mx+1)/2 - sum
}

/*
LC 3 - Longest substring without repeating characters

Given a string s, find the length of the longest substring without duplicate characters.
*/

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	longest := 0

	freq := map[byte]int{}
	for right < len(s) {
		freq[s[right]]++

		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}

		longest = max(longest, right-left+1)

		right++
	}

	return longest
}

/*
LC 42 - Trapping rain water

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/

func Trap(height []int) int {
	n := len(height)
	lmax, rmax := make([]int, n), make([]int, n)
	lmax[0] = height[0]
	rmax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		lmax[i] = max(height[i], lmax[i-1])
	}

	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(height[i], rmax[i+1])
	}

	water := 0
	for i := 1; i < n-1; i++ {
		water += min(lmax[i], rmax[i]) - height[i]
	}

	return water
}

/*
LC 1422 - Maximum score after splitting a string
*/
func MaxScore(s string) int {
	n := len(s)
	pSum := make([]int, n)

	if s[0] == '1' {
		pSum[0] = 1
	} else {
		pSum[0] = 0
	}

	for i := 1; i < n; i++ {
		if s[i] == '1' {
			pSum[i] = pSum[i-1] + 1
		} else {
			pSum[i] = pSum[i-1]
		}
	}

	maxScore := 0
	left := 1      // left substring size
	for left < n { // left-1 < n-1
		lZeros := left - pSum[left-1]
		rOnes := pSum[n-1] - pSum[left-1]
		maxScore = max(maxScore, lZeros+rOnes)
		left++
	}

	return maxScore
}

/*
LC 1732 - Find the highest altitude

There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i and i + 1 for all (0 <= i < n). Return the highest altitude of a point.
*/

func LargestAltitude(gain []int) int {
	currAlt := gain[0]
	maxAlt := max(0, currAlt)
	for i := 1; i < len(gain); i++ {
		currAlt += gain[i]
		maxAlt = max(maxAlt, currAlt)
	}

	return maxAlt
}

/*
LC 3432 - Count partitions with even sum difference

You are given an integer array nums of length n.

A partition is defined as an index i where 0 <= i < n - 1, splitting the array into two non-empty subarrays such that:

Left subarray contains indices [0, i].
Right subarray contains indices [i + 1, n - 1].
Return the number of partitions where the difference between the sum of the left and right subarrays is even.
*/
func CountPartitions(nums []int) int {
	n := len(nums)
	pSum := make([]int, n)
	pSum[0] = nums[0]

	for i := 1; i < n; i++ {
		pSum[i] = pSum[i-1] + nums[i]
	}

	pCount := 0
	for left := 0; left < n-1; left++ {
		lSum := pSum[left]
		rSum := pSum[n-1] - pSum[left]
		if (rSum-lSum)%2 == 0 {
			pCount++
		}
	}

	return pCount
}

/*
LC 349 - Given two integer arrays nums1 and nums2, return an array of their . Each element in the result must be unique and you may return the result in any order.
*/

func Intersection(nums1, nums2 []int) []int {
	f1 := map[int]int{}
	for _, n1 := range nums1 {
		f1[n1]++
	}

	intersect := []int{}
	for _, n := range nums2 {
		if f1[n]>0 {
			intersect = append(intersect, n)
			f1[n] = 0
		}
	}

	return intersect
}

/*
LC 653 - Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inOrder(node *TreeNode, items []int) []int {
	if node==nil {
		return items
	}

	items = inOrder(node.Left, items)
	items = append(items, node.Val)
	items = inOrder(node.Right, items)
	return items
}

func FindTarget(root *TreeNode, k int) bool {
	nodeValues := []int{}
	nodeValues = inOrder(root, nodeValues)

	left, right := 0, len(nodeValues)-1
	for left<right {
		if nodeValues[left]+nodeValues[right]==k {
			return true
		} else if nodeValues[left]+nodeValues[right]>k {
			right--
		} else {
			left++
		}
	}

	return false
}

/*
LC 594 - We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.
*/

func FindLHS(nums []int) int {
	slices.Sort(nums)
	maxSubseq := 0
	left, subseqLen := 0, 1

	for right:=1;right<len(nums);right++ {
		subseqLen += 1
		if nums[right]-nums[left] == 1 {
			maxSubseq = max(maxSubseq, subseqLen)
		} else {
			for left<right && nums[right]-nums[left] > 1 {
				subseqLen -= 1
				left++
			}
		}
	}

	return maxSubseq
}
