package pattern

//CheckMatrix 题目：link:https://leetcode.cn/problems/check-if-matrix-is-x-matrix/
//没啥难度，官方也是这样写的，判断matrix满足，正反对角线非0，非对角线为0
func CheckMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}

	return true
}
