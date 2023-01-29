package operateString

func CountAsterisks(s string) int {
	//verticalBars:=0     //这个可以用flag来标识，因为只有两个状态
	flag:=true				//换了flag之后时间还变慢了，好怪
	answer:=0
	for _,value:=range s{
		if value=='|'{
			flag=!flag
			//verticalBars=(verticalBars+1)%2
		//}else if verticalBars==0&&value=='*'{
		}else if flag&&value=='*'{
			answer++
		}

	}

	return answer

}