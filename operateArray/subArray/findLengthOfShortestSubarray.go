package subArray

// FindLengthOfShortestSubarray  https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
//感觉有点逆序对，有点贪心
//思路完全不对，是双指针，滑动窗口，连续子串，应该要想到滑动窗口，注意是连续子串。
//连续子串无非就是前中后三个部位，分三个部分思考一下，切断点是i，j
func FindLengthOfShortestSubarray(arr []int) int {
	i, j := 0, 0
	n := len(arr)
	answer := n
	//从尾巴开始找，先确定删头的那部分没问题
	for j = n - 1; 0 < j && arr[j-1] <= arr[j]; j-- {
	}
	answer = j
	k := arr[0]
	//从头部开始过滤，直到头不递增，确定删尾巴，删尾巴这边要处理一下，当j=n的时候要继续往下走
	//用k来当i的边界，让[1，2，3，2]，这种情况下的i=2过一下校验
	for ; i < j && k <= arr[i]; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		if j-i-1 < answer {
			answer = j - i - 1
		}
		k = arr[i]
	}
	return answer
}
