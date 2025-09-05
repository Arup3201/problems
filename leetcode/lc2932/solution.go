package solution

import "slices"

func MaximumStrongPairXor(nums []int) int {
	left := 0
	maxXOR := 0

	slices.Sort(nums)

	for right := left + 1; right < len(nums); right++ {
		for left < right && nums[right] > 2*nums[left] {
			left++
		}
		if nums[right] <= 2*nums[left] {
			for i := left; i < right; i++ {
				maxXOR = max(maxXOR, nums[i]^nums[right])
			}
		}
	}

	return maxXOR
}
