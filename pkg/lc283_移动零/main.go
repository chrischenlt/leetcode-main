package main

func main() {
	moveZeroes([]int{0, 1, 0, 2, 0})
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroes1(nums []int) {
	if len(nums) == 0 {
		return
	}

	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[a] = nums[i]
			a++
		}
	}

	for a < len(nums) {
		nums[a] = 0
		a++
	}
}
