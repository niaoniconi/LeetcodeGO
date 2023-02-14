package subArray

// LongestWPI  https://leetcode.cn/problems/longest-well-performing-interval/  今天的题，待会再转移一下位置
/**
我的暴力解
 */
func LongestWPI(hours []int) int {
	//官方前面的处理和我一样
	interval:=0
	for i,hour:=range hours{
		if hour>8{
			hours[i]=1
		}else{
			hours[i]=-1
		}
	}
	//暴力解
	for i:=0;i<len(hours);i++{
		count:=0
		for j:=i;j<len(hours);j++{
			count+=hours[j]
			if count>0&&j-i+1>interval{
				interval=j-i+1
			}
		}
	}
	return interval
}
