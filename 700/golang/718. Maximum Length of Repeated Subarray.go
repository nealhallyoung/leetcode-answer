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

// 动态规划
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0718.最长重复子数组.md
// 暴力
// s1 长度 m s2 长度 n
// dp[][] 是一个 (m+1)x(n+1) 的二维表
// s(i) s 的索引 0 到 i 的子串
// dp[i][j] 代表 s1(i-1) s2(j-1) 的最长子串的长度
// 如果 s1[i-1]=s2[j-1] 那么 dp[i][j]=dp[i-1][j-1]+1
// 如果 s1[i-1]!=s2[j-1] 那么说明，以 s1[i-1] 为结尾和以 s2[j-1] 为结尾的两个子串没有公共子串

func findLength(nums1 []int, nums2 []int) int {
	// 初始化二维表，所有值为0
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	res := 0
	// 从左到右，按行遍历
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

			}
			if dp[i][j] > 0 {
				res = dp[i][j]
			}
		}
	}
	return res
}
