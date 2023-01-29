package operateString

func CountAsterisks(s string) int {
	verticalBars:=0     //这个可以用flag来标识，因为只有两个状态
	answer:=0
	for _,value:=range s{
		if value=='|'{
			verticalBars=(verticalBars+1)%2
		}else if verticalBars==0&&value=='*'{
			answer++
		}

	}

	return answer

}