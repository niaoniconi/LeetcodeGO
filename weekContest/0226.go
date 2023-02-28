package weekContest

import "sort"

//周赛 https://leetcode.cn/circle/discuss/ldt5CI/
func leftRightDifference(nums []int) []int {
	answer := make([]int, len(nums))
	leftSum := 0
	rightSum := 0
	for i, v := range nums {
		answer[i] += leftSum
		answer[len(nums)-1-i] -= rightSum
		leftSum += v
		rightSum += nums[len(nums)-1-i]
	}
	for i, v := range answer {
		if v < 0 {
			answer[i] = -1 * v
		}
	}
	return answer
}

/**
"91221181269244172125025075166510211202115152121212341281327"
21
[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,1,0,1,0,0,0,0,0,0,0,0,0,1]

*/

//string 范围
func divisibilityArray(word string, m int) []int {
	answer := make([]int, len(word))
	prefix := 0
	for i, v := range word {
		prefix = prefix * 10
		prefix += int(v - '0')
		prefix %= m
		if prefix == 0 {
			answer[i] = 1
		}

	}
	return answer
}

//var queue []int
//current := 0
//count := 0
//for i, v := range nums {
//	if current*2 <= v || len(queue) == 0 {
//		queue = append(queue, v)
//		current = v
//	} else if queue[0]*2 <= v {
//		queue = queue[1:]
//		count += 2
//	}
//}
//count += len(queue) - len(queue)%2
//有想到是贪心，但没想到贪心的方向
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	count := 0
	var mark [][]int
	current := 0
	for i := 1; i < len(nums); i++ {
		for current < i && nums[current]*2 <= nums[i] {
			mark = append(mark, []int{current, i})
			current++
		}
	}
	sort.Slice(mark, func(i, j int) bool {
		return mark[i][1] < mark[j][1]
	})

	//这一半处理肯定有问题
	tail := len(nums) - 1
	used := make(map[int]int)
	for i := len(mark) - 1; i > -1; i-- {
		for used[tail] == -1 {
			tail--
		}
		if tail >= mark[i][1] {
			tail--
			used[mark[i][0]] = -1
			count += 2
		}
	}
	return count
}
