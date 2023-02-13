package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	left, right := 0, 1
	same := 0
	for right < len(nums) {
		if nums[right] == nums[left] {
			same++
			right++
			continue
		} else {
			left = right
			right++
		}
	}

	return len(nums) - same
}
