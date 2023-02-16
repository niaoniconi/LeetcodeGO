package remove

import "sort"

// NumberOfPairs https://leetcode.cn/problems/maximum-number-of-pairs-in-array/
/**
我自己是两个相同就计数，顺便跨国后一个相同的
 */
func NumberOfPairs(nums []int)[]int{
	answer:=make([]int,2)
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if i+1<len(nums)&&nums[i]==nums[i+1]{
			answer[0]++
			i++
		}
	}
	answer[1]=len(nums)-answer[0]*2
	return answer
}

// NumberOfPairsOfficial
/**
官方用的哈希表，来存当前是奇数个还是偶数个，由于bool默认值是false，所以奇数个是true，偶数个是false
每遇到一个数就先取反，如果是false，就算做一对，这个还是很巧妙的，不过没有key，map反馈的也会是false吗？
会，因为如果key不存在，会返回元素的类型的默认值
 */
func NumberOfPairsOfficial(nums []int)[]int{
	cnt := map[int]bool{}
	res := 0
	for _, num := range nums {
		cnt[num] = !cnt[num]
		if !cnt[num] {
			res++
		}
	}
	return []int{res, len(nums) - 2*res}
}