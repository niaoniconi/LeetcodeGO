package weekContest

import "sort"

func KItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	} else if k <= numOnes+numZeros {
		return numOnes
	} else {
		return numOnes - (k - numOnes - numZeros)
	}
}

// PrimeSubOperation1  只考虑了局部符合条件，全局就符合条件，没想到全局符合条件，但是局部不符合
//或许应该从尾指针开始找，这样的话，可以以一个全局解的范围来看   https://leetcode.cn/problems/prime-subtraction-operation/
//尾指针可能还解决不了，还有一个问题是，一个数只能减一次质数
//这个方法行不通，换一个.突然悟道，这个题目的本质是让前面的数尽可能小，看能不能找到一个合适的序列
func PrimeSubOperation1(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			j := 0
			for ; j < i && nums[j] < nums[i]; j++ {
			}
			p := findPrime1(nums[j], nums[j+1], nums[i])
			if p == -1 {
				return false
			} else {
				for ; j < i; j++ {
					nums[j] -= p
				}
			}

		}
	}
	return true
}

//返回一个最大质数p，使得j-p<x-p<i（j可减可不减）
func findPrime1(j, x, i int) int {
	begin := x - i + 1
	end := x - j - 1
	if j-1 > end {
		end = j - 1
	}
	prime := -1
	for ; begin <= end; begin++ {
		flag := true
		if begin < 2 {
			continue
		}
		for d := 2; d <= begin/d; d++ {
			if begin%d == 0 {
				flag = false
				break
			}
		}
		if flag {
			prime = begin
		}

	}
	return prime
}

// PrimeSubOperation 看清题目的本质是非常难得的问题，也就是做题目之间转换
func PrimeSubOperation(nums []int) bool {
	temp := 0
	for i, v := range nums {
		p := FindPrime(temp, v)
		if p != -1 {
			nums[i] = v - p
		}
		if nums[i] <= temp {
			return false
		}
		temp = nums[i]
	}
	return true
}

// FindPrime 返回y-p>x的最大p，p是质数,没有这个数时返回0
//思路要理清，要注意p的范围
func FindPrime(x, y int) int {
	begin := 2
	end := y - x - 1
	for ; begin <= end; end-- {
		flag := true
		for i := 2; i <= end/i; i++ {
			if end%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			return end
		}
	}
	return -1
}

func MinOperations(nums []int, queries []int) []int64 {
	answer := make([]int64, len(queries))
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	//控制好时间
	for i := 0; i < len(queries); i++ {
		key := sort.SearchInts(nums, queries[i])
		answer[i] = int64(queries[i]*key*2 - sum[key]*2 + sum[len(nums)] - queries[i]*len(nums))
	}
	//
	//超出时间限制
	//for _, v1 := range queries {
	//	temp := int64(0)
	//	for _, v2 := range nums {
	//		if v2 > v1 {
	//			temp += int64(v2 - v1)
	//		} else {
	//			temp += int64(v1 - v2)
	//		}
	//	}
	//	answer = append(answer, temp)
	//}
	return answer
}

//https://leetcode.cn/problems/collect-coins-in-a-tree/ 第四题
