package solution

/*
You are given a 0-indexed integer array nums and an integer threshold.

Find the length of the longest subarray of nums starting at index l and ending at index r (0 <= l <= r < nums.length) that satisfies the following conditions:

nums[l] % 2 == 0
For all indices i in the range [l, r - 1], nums[i] % 2 != nums[i + 1] % 2
For all indices i in the range [l, r], nums[i] <= threshold
Return an integer denoting the length of the longest such subarray.

Note: A subarray is a contiguous non-empty sequence of elements within an array.
*/

func LongestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	left := 0
	longest := 0

	for left < n && (nums[left]%2 != 0 || nums[left] > threshold) {
		left++
	}

	if left >= n {
		return longest
	}

	longest = 1
	for right := left + 1; right < n; right++ {
		if nums[right-1]%2 != nums[right]%2 && nums[right] <= threshold {
			longest = max(longest, right-left+1)
		} else {
			left = right
			for left < n && (nums[left]%2 != 0 || nums[left] > threshold) {
				left++
			}
		}
	}

	return longest
}
