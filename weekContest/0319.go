package weekContest

import "sort"

func EvenOddBit(n int) []int {
	mark1, mark2 := 0, 0
	for n > 0 {
		if n&1 == 1 {
			mark1++
		}
		n = n >> 1
		if n > 0 {
			if n&1 == 1 {
				mark2++
			}
			n = n >> 1
		}
	}
	return []int{mark1, mark2}
}

func CheckValidGrid(grid [][]int) bool {
	row := 0
	col := 0
	n := len(grid)
	flag := false
	move := [][]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {-2, 1}, {2, -1}, {-2, -1}}
	for i := 1; i < n*n; i++ {
		flag = false
		for _, v := range move {
			temp1 := row
			temp2 := col
			temp1 += v[0]
			temp2 += v[1]
			if temp1 >= 0 && temp1 < n && temp2 >= 0 && temp2 < n {
				if grid[temp1][temp2] == i {
					row = temp1
					col = temp2
					flag = true
					break
				}
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

//func BeautifulSubsets1(nums []int, k int) int {  //看错题了，看成了不大于k
//	sort.Ints(nums)
//	begin := 0
//	end := 0
//	pre := 0
//	count:=0
//	for end < len(nums) {
//		for end < len(nums)&&nums[end]-nums[begin] < k{ //先找到end，第一个子窗口
//			end++
//		}
//		pre=end
//		count+=1<<(end-pre)     //把这一个窗口的所有子集都加上了
//		//开始滑动
//		for end < len(nums)&&begin<len(nums)&&nums[end]-nums[begin] < k
//
//	}
//	return count
//}

func BeautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	for i, v := range nums {
		tempMap := make(map[int]int)
		tempMap[v] = 1
		temp := len(nums) - 1 - i
		for j := i + 1; j < len(nums); j++ {
			if tempMap[nums[j]-k] == 1 {
				temp--
			} else {
				tempMap[nums[j]] = 1
			}
		}
		count += 1 << temp
	}
	return count
}
