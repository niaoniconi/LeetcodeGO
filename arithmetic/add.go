package arithmetic

import "sort"

/**link:https://leetcode.cn/problems/two-sum/submissions/
官方解1：暴力解，双层遍历，时间：O(n)^2，空间：O(1) 需要注意不要重复，循环上三角或者下三角就行
官方暴力的方向:针对每一个nums[i]，去寻找nums[j]=target-nums[i]在数组中是否存在
我的暴力解的方向更偏于：针对每一个nums[i]，去寻找nums[i]+nums[j]=target在数组中是否存在

官方解2：使用哈希表，官方分析解法一的时间太长是因为target-nums[i]的重复度太高，会重复寻找已经确认过的结果
用哈希表的key来存储数组的值，这样只要查找哈希表就能确定target-nums[i]是否在数组中；用哈希表的值存数组的index，这样就能返回索引
时间：O(N)，空间：O(2N)
**注意：我和官方暴力解的思路不同导致了我们优化的思路不同

我的解：因为觉得nums[i]+nums[j]有很多重复，nums[i]+nums[i]大小无法决定如何缩小计算规模，
所以先排序，在根据nums[i]+nums[i]相对target的大小，移动i和j
时间：O(NlogN)-O(n)^2，空间：O(2N)

总结：我为什么会多用那么多时间：
1.我计算了很多两个数的和，这不是必须的，我需要的只是找到两个数，找到这个点很适合用哈希表，因为他是key-value的
2.我和官方解2空间上消耗都是一样的，因为我们需要保存一个键值的对应关系，虽然slice本身已经有了，但是不能根据value来排布或者找到key
*/
func TwoSum1(nums []int, target int) []int { //数组不能作为切片的形式传参   //我的解
	//sort.Ints(nums)
	numbers := make([][]int, len(nums), len(nums))
	for i, v := range nums {
		temp := make([]int, 2, 2)
		temp[0] = i
		temp[1] = v
		numbers[i] = temp
	}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i][1] > numbers[j][1]
	})
	head := 0
	answer := make([]int, 2, 2)
	tail := len(nums) - 1
	for tail > head {
		if numbers[tail][1]+numbers[head][1] == target {
			answer[0] = numbers[head][0]
			answer[1] = numbers[tail][0]
			break
		}
		if numbers[tail][1]+numbers[head][1] > target {
			head++
		}
		if numbers[tail][1]+numbers[head][1] < target {
			tail--
		}

	}
	return answer

}

// TwoSum2 官方返回int 切片有更简单的方法，看看 []int{p, i}
func TwoSum2(nums []int, target int) []int {
	hashArray := make(map[int]int)
	//answer := make([]int, 2, 2)
	for i, v := range nums {
		value, ok := hashArray[target-v]
		if ok {
			//answer[0] = i
			//answer[1] = value
			return []int{i, value}
		}
		hashArray[v] = i

	}
	return nil
}
