package main

//class Solution {
//    public int largestRectangleArea(int[] heights) {
//        Stack<Integer> stack = new Stack<>();
//        stack.push(-1);
//        int maxarea = 0;
//        for (int i = 0; i < heights.length; i++) {
//            while (stack.peek() != -1 && heights[stack.peek()] >= heights[i]) {
//                maxarea = Math.max(maxarea, heights[stack.pop()] * (i - stack.peek() - 1));
//            }
//            stack.push(i);
//        }
//        while (stack.peek() != -1) {
//            maxarea = Math.max(maxarea, heights[stack.pop()] * (heights.length - 1 - stack.peek()));
//        }
//        return maxarea;
//    }
//}

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	var maxarea int
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[i] <= heights[stack[len(stack)-1]] {
			maxarea = max(maxarea, (i-stack[len(stack)-2]-1)*heights[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		maxarea = max(maxarea, (len(heights)-1-stack[len(stack)-2])*heights[stack[len(stack)-1]])
		stack = stack[:len(stack)-1]
	}
	return maxarea
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// timeout
func largestRectangleArea11(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	var res int
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left >= 0 {
			left--
			if left < 0 || heights[left] < heights[i] {
				break
			}
		}
		for right <= len(heights)-1 {
			right++
			if right >= len(heights) || heights[right] < heights[i] {
				break
			}
		}

		area := (right - left - 1) * heights[i]
		if area > res {
			res = area
		}
	}
	return res
}

// timeout
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
