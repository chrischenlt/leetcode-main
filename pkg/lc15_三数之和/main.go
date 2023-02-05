package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	//sum := threeSum(nums)
	fmt.Println(nums)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[left], nums[right], -target})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)
	mm := make(map[int]int)
	set := make(map[string]struct{})
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	for _, num := range nums {
		v, exist := mm[num]
		if exist {
			mm[num] = v + 1
		} else {
			mm[num] = 1
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < 0 {
				break
			}
			mm[nums[i]] = mm[nums[i]] - 1
			mm[nums[j]] = mm[nums[j]] - 1
			target := -(nums[i] + nums[j])
			v, exist := mm[target]
			if !exist {
				mm[nums[i]] = mm[nums[i]] + 1
				mm[nums[j]] = mm[nums[j]] + 1
				continue
			}
			if v > 0 {
				b := []int{nums[i], nums[j], target}
				sort.Slice(b, func(i, j int) bool {
					return b[i] < b[j]
				})
				s := getString(b)
				_, ok := set[s]
				if !ok {
					res = append(res, b)
					set[s] = struct{}{}
				}
			}
			mm[nums[i]] = mm[nums[i]] + 1
			mm[nums[j]] = mm[nums[j]] + 1

		}
	}
	return res
}

func getString(nums []int) string {
	var res string
	for _, num := range nums {
		res = res + strconv.Itoa(num)
	}
	return res
}
