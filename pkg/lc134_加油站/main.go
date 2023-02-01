package main

func main() {
}

func canCompleteCircuit(gas []int, cost []int) int {
	less, run, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		less += gas[i] - cost[i]
		run += gas[i] - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}

	if less >= 0 {
		return start
	}
	return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {

	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}

	length := len(gas)
	for i := 0; i < length; i++ {
		if gas[i] < cost[i] {
			continue
		}

		cur := i + 1
		less := gas[i] - cost[i]
		canComplete := true
		for j := 1; j < length; j++ {
			if cur >= length {
				cur = 0
			}
			less += gas[cur] - cost[cur]

			if less < 0 {
				canComplete = false
				break
			}

			cur++
		}
		if canComplete {
			return i
		}
	}

	return -1
}
