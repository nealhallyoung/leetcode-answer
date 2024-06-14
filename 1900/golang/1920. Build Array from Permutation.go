func buildArray(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = nums[nums[i]]
	}
	return res
}
