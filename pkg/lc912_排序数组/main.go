package main

func main() {
}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	buildHeap(nums, (len(nums)-2)/2)
	heapSort(nums, len(nums)-1)
	return nums
}

func heapSort(nums []int, index int) {
	for i := index; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
}

func buildHeap(nums []int, index int) {
	for i := index; i >= 0; i-- {
		heapify(nums, len(nums), i)
	}
}

func heapify(nums []int, length, index int) {
	if index >= length {
		return
	}
	c1 := index*2 + 1
	c2 := index*2 + 2
	max := index
	if c1 < length && nums[c1] > nums[max] {
		max = c1
	}
	if c2 < length && nums[c2] > nums[max] {
		max = c2
	}
	if max != index {
		nums[max], nums[index] = nums[index], nums[max]
		heapify(nums, length, max)
	}
}

func sortArray1(nums []int) []int {
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
