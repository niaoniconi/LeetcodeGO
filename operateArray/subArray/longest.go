package subArray

// LongestWPI  https://leetcode.cn/problems/longest-well-performing-interval/  今天的题，待会再转移一下位置
/**
我的暴力解
*/
func LongestWPI(hours []int) int {
	//官方前面的处理和我一样,预处理，省略不必要的信息
	interval := 0
	for i, hour := range hours {
		if hour > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	//暴力解
	for i := 0; i < len(hours); i++ {
		count := 0
		for j := i; j < len(hours); j++ {
			count += hours[j]
			if count > 0 && j-i+1 > interval {
				interval = j - i + 1
			}
		}
	}
	return interval
}

// LongestWPIGreedy
/*
处理前缀和，s[i]是前i-1天的hour（处理过的）s前缀和，
那么原问题可以看做求解区间分数和大于 0 的最长区间长度。我们要求的是一个满足要求的最大区间[l,r]，s[r]-s[l]>0
有两个值需要确定，所以暴力解需要n^2的时间
如果我们先固定一个r，然后向左找l，如果有l1，l2都满足要求且l1<l2，那么l1更符合r的要求，且l2不会是任何r的l，因为有更优的l1（但是l2可能是r）
1.我们可以维护一个栈，栈里存s[0]到s[i-1]的递减数列
（栈顶肯定最近一个符合要求的数，栈元素中间的数是一个不符合要求的区间，由于找到是最长，所以需要尽量包括这些不合格的区间),
2.怎么根据r找l，一直弹出栈顶元素，直到栈顶元素是小于s[r]的最小元素
贪心算法：先找到一定满足要求的解，在找到满足要求的最大解

坏了，官方没有提供go的解，太痛啦
*/
func LongestWPIGreedy(hours []int) int {
	interval := 0
	s := make([]int, len(hours)+1)
	var stk []int

	//s[i]  前n个数的hours的和 ，如果s[i]>0，说明 或者s[j]>s[i]，说明j-i这个区间里，是一个好区间
	for i, hour := range hours {
		if hour > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
		s[i+1] = s[i] + hours[i]
	}

	//用数组模拟栈,stk是栈，栈里压着递减的hours
	for i := 0; i < len(hours); i++ {

		if len(stk) == 0 || s[i] < s[stk[len(stk)-1]] {
			if len(stk) != 0 {
				t := stk[len(stk)-1]
				t = t - t
			}

			stk = append(stk, i)
		}
	}
	//从后往前，固定r，找l
	for r := len(hours); len(stk) > 0 && r >= 1; r-- {
		//temp是栈顶元素
		temp := stk[len(stk)-1]
		var l int //使用l:=r,如果len(hours)是7，l会是6，而不是7，是因为r--吗？：=和等于的区别？
		//默认的左边，如果用l当栈顶元素，那么就算 r不大于栈顶元素，也会算上
		l = r
		//如果r大于栈顶元素，那么出栈
		for (r > temp && s[r] > s[temp]) || r == temp {
			l = temp
			stk = stk[:len(stk)-1]

			if len(stk) == 0 {
				break
			}
			//出栈之后更新栈顶元素
			temp = stk[len(stk)-1]
		}
		if r-l > interval {
			interval = r - l
		}

	}
	return interval
}

