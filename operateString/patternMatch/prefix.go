package patternMatch

import "sort"

//RemoveSubfolders https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/
/**
官方思路一和我一样，这次差不多是一遍过，嘻嘻
思路二是字典树，可以再学习一下
 */
func RemoveSubfolders(folder []string)[]string{
	var answer []string
	sort.Strings(folder)
	shortName:=folder[0]
	answer=append(answer,shortName)
	for _,folderName:=range folder[1:]{
		if MatchFolderPrefix(folderName,shortName){
			continue
		}else{
			shortName=folderName
			answer=append(answer,shortName)
		}
	}
	return answer
}

func MatchFolderPrefix(folder string,pattern string) bool{
	//没有重复的，等于可以排除
	if len(folder)<=len(pattern){
		return false
	}
	prefix:=folder[0:len(pattern)]
	if prefix!=pattern{
		return false
	}
	//这边应该不需要大于，因为没有重复的
	if folder[len(pattern)]!='/'{
		return false
	}
	return true
}