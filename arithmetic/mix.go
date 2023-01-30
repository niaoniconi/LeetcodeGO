package arithmetic

//  link:https://leetcode.cn/problems/calculate-amount-paid-in-taxes/submissions/
func CalculateTax(brackets [][]int, income int) float64 {
	//这个思路是存了一个`rest，用来标注还没用来计税的数量，前一档的upper使用income和rest算出来了的，
	//其实不用如此，只要存前一个的upper就行了
	var answer float64
	rest := income
	temp := 0
	for _, v := range brackets {
		if income < v[0] {
			temp += rest * v[1]
			break
		} else {
			temp += (v[0] - income + rest) * v[1]
			rest = income - v[0]
		}

	}
	//fmt.Println("temp is ",temp)
	answer = float64(temp)
	answer = answer / 100
	//todo I have to research why my Sprintf didn't work
	// fmt.Println("answer is ",answer)
	// fmt.Println(fmt.Sprintf(".5f",answer))
	// fmt.Sprintf(".5f",answer)
	// answer1,_:=strconv.ParseFloat(fmt.Sprintf(".5f",answer),64)
	return answer

}

func CalculateTax2(brackets [][]int, income int) float64 {
	//改良版，别的不说，代码可读性高了很多
	var answer float64
	temp := 0
	temp2 := 0
	for _, v := range brackets {
		if income > v[0] {
			temp += (v[0] - temp2) * v[1]
			temp2 = v[0]
		} else {
			temp += (income - temp2) * v[1]
			break
		}
	}
	answer = float64(temp)
	answer = answer / 100
	return answer

}
