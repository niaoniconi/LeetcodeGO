package patternMatch

// BalancedString https://leetcode.cn/problems/replace-the-substring-for-balanced-string/
func BalancedString(s string) int {
	minSubstrLength := len(s)
	count := make([]int, 4)
	mark := make([][]int, len(s))
	for i, value := range s {
		switch value {
		case 'Q':
			count[0]++
		case 'W':
			count[1]++
		case 'E':
			count[2]++
		case 'R':
			count[3]++
		}
		mark[i] = count[:]
	}

	var chars []int
	for i, value := range count {
		if value > len(s)/4 {
			value -= len(s) / 4
			chars = append(chars, i)
		} else {
			value = 0
		}
	}
	if count[0] == 0 && count[1] == 0 && count[2] == 0 && count[3] == 0 {
		return 0
	}
	//现在的目标是在string中找到含字符的最小子串 i,j 是的mark[j]-mark[i]>=count,
	//二分查找
	//minLength := count[0] + count[1] + count[2] + count[3]
	//for i,value:=range mark{
	//	if mark
	//}
	return minSubstrLength
}
