package main

import (
	"fmt"
	"github.com/LeetcodeGO/aboutSort/findKs"
	"sort"
)

func main() {
	taps :=[]int{3,4,1,1,0,0}
	answer:=findKs.MinTaps(5,taps)
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

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	nums:=append(nums1,nums2...)
	sort.Slice(nums,func(i int,j int)bool{
		return nums[i][0]<nums[j][0]
	})
	for i:=0;i<len(nums)-1;i++{
		if nums[i][0]==nums[i+1][0]{
			nums[i][1]+=nums[i+1][1]
			nums=append(nums[:i+1],nums[i+2:]...)
		}
	}
	return nums
}

func minOperations(n int) int {
	count:=0
	for n>0{
		consecutive:=0
		for n>0&&(n&1==0){
			n=n>>1
			continue
		}
		for n>0&&(n&1==1){
			count++
			consecutive++
			n=n>>1
		}
		if consecutive>1{
			count=count-consecutive+1
			n=n+1
		}
	}
	return count
}

//看错题了，尴尬
func squareFreeSubsets(nums []int) int {
	goodNumber:=0
	//go的^次方操作，是不是自有其问题
	//有时间排查一下
	for _,value:=range nums{
		if isProPer(value){
			goodNumber++
		}
	}
	x:=1
	for i:=0;i<goodNumber;i++{
		x=x<<1
		x=x%(1000000007)
	}
	return x-1
}
func  isProPer(x int)  bool{
	for i:=2;i<=x/i;i++{
		if x%(i*i)==0{
			return false
		}
	}
	return true
}

