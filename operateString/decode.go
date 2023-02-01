package operateString

// DecodeMessage https://leetcode.cn/problems/decode-the-message/solution/jie-mi-xiao-xi-by-leetcode-solution-wckx/
//官方和我的思路一样，但是他用的map的value是byte，为什么呢？string用的也是byte
func DecodeMessage(key string,message string) string{
	table:=make(map[rune]rune)
	temp:='a'
	table[' ']=' '  //另外，官方解，没有在替换表里存空格，而是在替换的时候对空格进行了过滤，为什么，难道map的key value实际上是需要遵照的吗？
	for _,v:=range key{
		if _,ok:=table[v];!ok{  	//这边还是很精妙的，用了map来确定唯一，不用循环验证，然后用temp存值，保证了按出现顺序编表
			table[v]=temp
			temp=temp+1
		}
	}
	answer:=[]rune(message)		//go也不能更改string
	for i,value:=range answer{
		answer[i]=table[value]
	}
	return string(answer)
}