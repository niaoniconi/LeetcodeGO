package subArray

func findSubArrays(nums []int) bool {
	twoSum := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		temp := nums[i] + nums[i+1]
		if twoSum[temp] != 0 {
			return true
		} else {
			twoSum[temp] = 1
		}
	}
	return false
}
