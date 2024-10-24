https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0718.最长重复子数组.md

// 暴力
func findLength(nums1 []int, nums2 []int) int {
	maxlen := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			k := 0
			for i+k < len(nums1) && j+k < len(nums2) && nums1[i+k] == nums2[j+k] {
				k++
			}
			maxlen = max(maxlen, k)
		}

	}
	return maxlen
}
