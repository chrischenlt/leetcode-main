package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		a := (right - left) * min(height[left], height[right])
		if a > max {
			max = a
		}

		if height[left] >= height[right] {
			right--
		} else {
			left++
		}
	}

	return max
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
