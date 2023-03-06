package patternMatch

// MinimumDeletions https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/
/**
逆序对问题
返回需要删减的足校字符个数，让字符串符合aaaaabbbb，无ba逆序对
string转成int：
int, err := strconv.Atoi(string)
string转成int64：
int64, err := strconv.ParseInt(string, 10, 64)
int转成string：
string := strconv.Itoa(int)
int64转成string：
string := strconv.FormatInt(int64,10)
————————————————
版权声明：本文为CSDN博主「排骨瘦肉丁」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/iamlihongwei/article/details/79550958
 */
func MinimumDeletions(s string) int {
	answer:=0
	count:=make([]int,len(s))
	aCount,bCount:=0,0
	for i,ch:=range s{
		//从前往后统计a前面的b
		if ch=='a'{
			count[i]=bCount
		}else{
			bCount++
		}
		//从后往前统计b后面的a
		if s[len(s)-1-i]=='b'{
			count[len(s)-1-i]=aCount
		}else{
			aCount++
		}
	}
	//再用一个栈吧，有点没想好应该怎么样，再想想

	return answer
}