package sum

import "sort"

// GetMaximumConsecutive https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make/
/**
整体思路：如果要算数列能表现的最大数，那就是所有数相乘再相加
但是要连续：所以如果有表现的数有中断的话，就停止。当前能表示的最大数k，找不到k+1
怎么算找不到k+1呢？打个比方，当前所有小于k+2的数相加，结果为k，coin剩下的最临近k的数就是k+2了，k到k+2只能中断了
所以整体程序有两个出口：1：没有k+1，也就是到nextCoin  2.到maxCoin
时间复杂度：max(O(maxiCoin)*O(n))
空间复杂度：O(n)


我这个函数设计的比较复杂，主要是考虑到一些多个重复出现的数，要是重复比较多，我的函数效率更高，要是数很大，重复不高，官方解更优秀
*/
func GetMaximumConsecutive(coins []int) int {
	coinsNum := make(map[int]int)
	coinsNum[0] = 1
	maxCoin := 0
	for _, v := range coins {
		if _, ok := coinsNum[v]; ok {
			coinsNum[v]++
		} else {
			if v > maxCoin {
				maxCoin = v
			}
			coinsNum[v] = 1
		}
	}
	currentCoin := 0 //表示到当前硬币大小
	maximum := 0     //表示到当前硬币大小最大能表示多少个数
	nextCoin := 1    //表示要找到的下一个coin
	for i := currentCoin; i <= nextCoin; i++ {
		if i > maxCoin {
			break
		}

		if numbers, ok := coinsNum[i]; ok {
			maximum += numbers * i
			nextCoin = maximum + 1
		}
	}
	return maximum + 1
}

// GetMaximumConsecutiveOfficial  官方解真的很简单，理论很简单，排序之后，
//如果前k个数能表示到maximum，那么只要下一个数coins[k+1] <=maximum+1，那么到下一个数，能表示的数是maximum+=coins[k+1]
//只要coins[k+1] <=maximum，就可以一直算到最后一个，否则中断
/**
其实这个问题主要是明确什么是中断
中断就是前k个小硬币的能表示的数（加上0），小于k+1大的硬币的值
所以只要，第k+1个硬币小于前k个小硬币的能表示的数maximum就行

另外还有一个点，如果前k个硬币能表示maximum，coins[k+1]<maximum,那么maximum+=coins[k+1]
这个迭代关系必须知道哈，如果两个硬币能表示4个连续数字，那么再加上3，必定能连续表示7，这个应该很好理解
*/
func GetMaximumConsecutiveOfficial(coins []int) int {
	sort.Ints(coins)
	maximum := 1
	for _, v := range coins {
		if v > maximum {
			return maximum
		}
		maximum += v
	}
	return maximum
}
