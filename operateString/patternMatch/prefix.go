package patternMatch

import (
	"sort"
	"strconv"
)

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

// GetFolderNames https://leetcode.cn/problems/making-file-names-unique/
/**
差不多一遍过，第一次运行的时候发现和题目揣测的意思不一样
改了之后一遍过。另外，记住int转string不要强制转，会出问题，要用函数
 */
func GetFolderNames(names []string) []string {
	var answer []string
	file:=make(map[string]int)
	for _,v :=range names{
		//if file[v]==0{
		//	answer=append(answer,v)
		//	file[v]=1
		//}else{
		//	temp:=v
		//	//temp=v+"("+string(file[v])+")"  //这个有问题，转出来格式不对
		//	temp=v+"("+strconv.Itoa(file[v])+")"
		//	answer=append(answer,temp)
		//	file[v]+=1
		//}
		temp:=v
		for file[temp]!=0{                    //直到不同名为止，防止跳跃性的同名，比如a(10)
			temp=v+"("+strconv.Itoa(file[v])+")"
			file[v]+=1                        //在同名初始文件记录相同文件的个数
		}
		answer=append(answer,temp)
		file[temp]=1                         //记录不同名的文件
	}
	return answer
}
