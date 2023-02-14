package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0, 0, 0, 0}
	answer := countFairPairs(nums, 0, 0)
	fmt.Println("answer: ", answer)

}

//周赛题
func countFairPairs(nums []int, lower int, upper int) int64 {
	answer := int64(0)
	maxIndex := len(nums) - 1
	minIndex := 1
	sort.Ints(nums)
	for i, value := range nums {
		minIndex = i + 1
		for minIndex < len(nums) && value+nums[minIndex] < lower {
			minIndex++
		}
		for maxIndex > 0 && value+nums[maxIndex] > upper {
			maxIndex--
		}
		if maxIndex == 0 {
			break
		}
		if minIndex <= maxIndex && minIndex < len(nums) {
			answer += int64(maxIndex - minIndex + 1)
		}

	}
	return answer
}


