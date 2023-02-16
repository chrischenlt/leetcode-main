package main

func largestRectangleArea1(heights []int) int {

	if len(heights) == 0 {
		return 0
	}
	var res int
	for i := 0; i < len(heights); i++ {
		m := heights[i]
		for j := i; j < len(heights); j++ {
			a := min(heights[i], heights[j])
			if a < m {
				m = a
			}
			area := (j - i + 1) * m
			if area > res {
				res = area
			}
		}
	}

	return res
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
