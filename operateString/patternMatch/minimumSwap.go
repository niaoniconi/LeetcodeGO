package patternMatch

// MinimumSwap 只能说完美，perfect，一眼看穿题目的本质。已排除同样的，在发现xx，yy消除只需要交换一次；xy，yx交换需要两次
//重点是i，j交换不需要考虑距离，只需要考虑有多少对即可
func MinimumSwap(s1 string, s2 string) int {
	countDiffX := 0
	countDiffY := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			switch s1[i] {
			case 'x':
				countDiffX++
			case 'y':
				countDiffY++
			}
		}
	}
	if (countDiffX+countDiffY)%2 != 0 {
		return -1
	}
	if countDiffY%2 == 0 && countDiffX%2 == 0 {
		return countDiffY/2 + countDiffX/2
	}

	return countDiffY/2 + countDiffX/2 + 2
}
