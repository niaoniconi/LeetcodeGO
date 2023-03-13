package weekContest

import "sort"

func VowelStrings(words []string, left int, right int) int {
	answer := 0
	yuanyin := make(map[rune]int)
	yuanyin['a'] = 1
	yuanyin['e'] = 1
	yuanyin['i'] = 1
	yuanyin['o'] = 1
	yuanyin['u'] = 1
	for i := left; i <= right; i++ {
		ss := words[i]
		if yuanyin[rune(ss[0])] != 0 && yuanyin[rune(ss[len(ss)-1])] != 0 {
			answer++
		}
	}
	return answer
}

func MaxScore(nums []int) int {
	sort.Ints(nums)
	prefix := 0
	score := 0
	for i := len(nums) - 1; i >= 0; i-- {
		prefix += nums[i]
		if prefix > 0 {
			score++
		}
	}
	return score
}

//FindMinimumTime https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
//func FindMinimumTime(tasks [][]int) int {
//	sort.Slice(tasks, func(i, j int) bool {
//		return tasks[i][1]-tasks[i][0]+1-tasks[i][2] < tasks[j][1]-tasks[j][0]+1-tasks[j][2]
//	})
//
//}

// BeautifulSubarrays  连续非空的子数组的个数
func BeautifulSubarrays(nums []int) int64 {
	count := 0
	prefix := make(map[int]int)
	current := 0
	prefix[0] = 1
	for _, v := range nums {
		current = current ^ v
		if prefix[current] != 0 {
			count += prefix[current]
			prefix[current]++
		} else {
			prefix[current] = 1
		}
	}
	return int64(count)
}
