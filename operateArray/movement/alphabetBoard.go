package movement

//AlphabetBoardPath https://leetcode.cn/problems/alphabet-board-path/submissions/
/*
官方思路和我差不多，但是官方没有特殊处理到让z先走到u，而是先向上，再向左，（这样从z出发就不会有问题）
再向下，然后向右（这样到z就不会有问题）
如果是其他字母到z，就应该先向左，后向下（上不会走）
如果是z到其他字母，就应该先上，后往右，只要满足这两个相对顺序就不会有问题
官方相当于搞成了互不关联的四种情况
官方把上下左右都搞成了if
*/
func AlphabetBoardPath(target string) string {
	var path []rune
	currentRow := int32(0)
	currentColumn := int32(0)
	for _, aim := range target {
		aimRow := (aim - 'a') / 5
		aimColumn := (aim - 'a') % 5

		//对z的特殊处理，从z出发，如果不是到z，先往上到U
		if currentRow == 5 && aimRow != 5 {
			currentRow--
			path = append(path, 'U')
		}

		//因为z的原因，其他位置到z，先平移再上下
		if currentColumn < aimColumn {
			for ; currentColumn < aimColumn; currentColumn++ {
				path = append(path, 'R')
			}
		} else {
			for ; aimColumn < currentColumn; currentColumn-- {
				path = append(path, 'L')
			}
		}

		if currentRow < aimRow {
			for ; currentRow < aimRow; currentRow++ {
				path = append(path, 'D')
			}
		} else {
			for ; aimRow < currentRow; currentRow-- {
				path = append(path, 'U')
			}
		}

		path = append(path, '!')
	}
	return string(path)
}

// AlphabetBoardPathOfficial  官方是上左下右，我觉得左下上右也行，试试，果然只要满足相对顺序就行
func AlphabetBoardPathOfficial(target string) string {
	var path []rune
	currentRow := int32(0)
	currentColumn := int32(0)
	for _, aim := range target {
		aimRow := (aim - 'a') / 5
		aimColumn := (aim - 'a') % 5

		if aimColumn < currentColumn {
			for ; aimColumn < currentColumn; currentColumn-- {
				path = append(path, 'L')
			}
		}

		if currentRow < aimRow {
			for ; currentRow < aimRow; currentRow++ {
				path = append(path, 'D')
			}
		}

		if aimRow < currentRow {
			for ; aimRow < currentRow; currentRow-- {
				path = append(path, 'U')
			}
		}

		if currentColumn < aimColumn {
			for ; currentColumn < aimColumn; currentColumn++ {
				path = append(path, 'R')
			}
		}

		path = append(path, '!')
	}
	return string(path)
}
