package main

func main() {
}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	left, right := start, end
	pivot := nums[(right+left)/2]
	for left <= right {
		for left <= right && nums[left] < pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	quickSort(nums, start, right)
	quickSort(nums, left, end)
}
