package numberTheory

import "strconv"

// IsPalindrome https://leetcode.cn/problems/palindrome-number/solution/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
