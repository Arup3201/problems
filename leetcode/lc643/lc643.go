package lc643

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/

func FindMaxAverageSum(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum
	for i:=1;i<=n-k;i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
