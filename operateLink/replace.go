package operateLink

// MergeInBetween https://leetcode.cn/problems/merge-in-between-linked-lists/
//这个for循环写的比较多，会意识到，for循环之后i的值就是等于那个条件，比如i<k；i++循环，循环完了之后i=k
//官方解和我差不多不过少了对前a个数的循环
func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var preA, afterB *ListNode //a之前的节点，b之后的节点
	preA = list1

	for i := 0; i < a-1; i++ {
		preA = preA.Next
	}
	//这是我的
	//for i := 0; i < b+1; i++ {
	//	afterB = afterB.Next
	//}
	afterB = preA
	//官方少一点循环
	for i := 0; i < b-a+2; i++ {		//从a到b共有b-a+1个点，因为preA在A之前需要算上A，到b+1有b-a+2个点
		afterB = afterB.Next
	}
	preA.Next = list2
	for list2.Next != nil { //list2由头节点变成尾节点
		list2 = list2.Next
	}
	list2.Next = afterB
	return list1
}
