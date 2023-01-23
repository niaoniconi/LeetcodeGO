package arithmetic_leet

import "sort"

func TwoSum(nums []int, target int) []int {
	sort.Ints(nums)
	head := 0
	answer := make([]int, 2, 2)
	tail := len(nums) - 1
	for tail > head {
		if nums[tail]+nums[head] == target {
			answer[0] = head
			answer[1] = tail
			break
		}
		if tail+head > target {
			head++
		}
		if tail+head < target {
			tail--
		}

	}
	return answer
}
