package patternMatch

// BalancedString https://leetcode.cn/problems/replace-the-substring-for-balanced-string/
func BalancedString(s string) int {
	minSubstrLength := len(s)
	mark := make([][]int, len(s)+1)
	mark[0]=make([]int,4)
	count:=make([]int,4)
	for i, value := range s {
		mark[i+1]=make([]int,4)
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
		for j:=0;j<4;j++{
			mark[i+1][j]=count[j]
		}
	}
	if mark[len(s)][0]==mark[len(s)][1]&&mark[len(s)][1]==mark[len(s)][2]&&mark[len(s)][2]==mark[len(s)][3]{
		return 0
	}
	minLength:=0
	//搞一下minLength
	for i:=0;i<4;i++{
		if mark[len(s)][i]>len(s)/4{
			count[i]-=len(s)/4
			minLength+=mark[len(s)][i]-len(s)/4

		}else{
			count[i]=0
			for j:=0;j<len(s)+1;j++{
				mark[j][i]=0
			}
		}
	}

	//不能简单的二和一，需要都含有
	//现在mark求和，只存了，需要替换的字符的数量
	//现在的目标是在string中找到含字符的最小子串 i,j 是的mark[j]-mark[i]>=count,
	//使用字典
	lengthPosi:=make(map[int]int)  //需要喜欢的字母数量，对应最后面的i
	for i,value:=range mark{
		sum:=value[0]+value[1]+value[2]+value[3]
		if sum<minLength{
			lengthPosi[sum]=i
		}else{
			before:=lengthPosi[sum-minLength]
			if i-before<minSubstrLength{
				minSubstrLength=i-before
			}
			lengthPosi[sum]=i
		}
	}
	return minSubstrLength
}
