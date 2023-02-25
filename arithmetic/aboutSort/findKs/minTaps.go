package findKs

// MinTaps  https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/
func MinTaps(n int, ranges []int) int {
	tapsRange := make(map[int]int)
	//统计有效的水龙头,及其能覆盖的范围,同一个起点，只存最大的
	//思考过用map还是用int，最后用了map，官方贪心算法用的是int array。我这样，空水龙头会很省力
	//坏消息，还是不会动态规划
	for i, value := range ranges {
		if value != 0 {
			begin := i - value
			end := i + value
			if begin < 0 {
				begin = 0
			}
			if end > n {
				end = n
			}
			if end > tapsRange[begin] {
				tapsRange[begin] = end
			}
		}
	}

	/**
	last, ret, pre := 0, 0, 0
	    for i := 0; i < n; i++ {
	        last = max(last, rightMost[i])
	        if i == last {
	            return -1
	        }
	        if i == pre {
	            ret++
	            pre = last
	        }
	    }

	官方的第二部分要流畅很多，有一个pre来记录上一个用完的区间
	坏消息：我仍然不会动态规划
	好消息：我学会了贪心算法
	*/

	//来一个for循环，从前向后，选水龙头
	end := 0
	count := 0
	for i := 0; i < n; i++ {
		flag := false //没找到水龙头
		temp := end
		for ; i <= temp; i++ { //原本是end出错了，这种按一个参数迭代的，迭代还会更新参数的，最后用一个temp
			if tapsRange[i] > end {
				end = tapsRange[i]
				flag = true
			}
		}
		if flag {
			count++
			i = temp
			if end >= n { //说好要在end这边做限制，结果忘了，浪费了机会
				break
			}
		} else {
			return -1
		}
	}
	return count
}
