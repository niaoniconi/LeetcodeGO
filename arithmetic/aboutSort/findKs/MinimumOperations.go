package findKs

// MinimumOperations 本来想排序了，后来发现有更好的办法，本来只是为了统计不同的个数,做的很棒
//https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/
func MinimumOperations(nums []int) int {
	numbers := make(map[int]int)
	for _, v := range nums {
		if v != 0 && numbers[v] != 0 {
			numbers[v] = 1
		}
	}
	return len(numbers)
}

// QuickSort  快速排序的教程：https://www.cnblogs.com/MOBIN/p/4681369.html
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	key := 0
	mid := nums[0]
	low := key
	high := len(nums) - 1
	for low < high {
		//一定不能让没交换的high<low，所以high和low要谨慎，而且不能同时操作
		//先右向左，确保key会比high小，重点是左右顺序
		for mid <= nums[high] && low < high {
			high--
		}
		if low < high {
			nums[key] = nums[high]
			nums[high] = mid
			key = high
			low++ //一定要确保low之前的都是验证过得
		}
		for nums[low] <= mid && low < high {
			low++
		}
		if low < high {
			//当前low的值就是大于mid。设定low也大于mid
			nums[key] = nums[low]
			nums[low] = mid
			key = low
			high-- //一定要确保high之后的都是验证过得
		}
	}
	QuickSort(nums[:key])
	QuickSort(nums[key+1:])
}
