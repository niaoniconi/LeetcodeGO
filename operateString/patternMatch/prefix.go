package patternMatch

import "sort"

//RemoveSubfolders https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/
/**
官方思路一和我一样，这次差不多是一遍过，嘻嘻
strings.HasPrefix(a string,b string)bool  string有这个函数，记录一下，明天再看
思路二是字典树，可以再学习一下：就是用一个树形结构来存，之前我也想过，所以差点放到树那边了，但是嫌分割字符串太麻烦，从而没有试
*/
func RemoveSubfolders(folder []string) []string {
	var answer []string
	sort.Strings(folder)
	shortName := folder[0]
	answer = append(answer, shortName)
	for _, folderName := range folder[1:] {
		if MatchFolderPrefix(folderName, shortName) {
			continue
		} else {
			shortName = folderName
			answer = append(answer, shortName)
		}
	}
	return answer
}

// MatchFolderPrefix only for "/abc" folder
func MatchFolderPrefix(folder string, pattern string) bool {
	//没有重复的，等于可以排除
	if len(folder) <= len(pattern) {
		return false
	}
	prefix := folder[0:len(pattern)]
	if prefix != pattern {
		return false
	}
	//这边应该不需要大于，因为没有重复的
	if folder[len(pattern)] != '/' {
		return false
	}
	return true
}
