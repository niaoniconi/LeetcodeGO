package main

import (
	"fmt"
	"github.com/LeetcodeGO/operateArray/remove"
	"sort"
)

func main() {
	nums := []int{18,18,18,15,54,45,45,45}
	//answer := subArray.LongestWPIGreedy(nums)
	answer:= remove.NumberOfPairsOfficial(nums)
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


