func climbStairs(n int) int {
	ret := make([]int, n+1)
	ret[0] = 1 // 站在地面的情况
	ret[1] = 1
	for i := 2; i < n+1; i++ {
		ret[i] = ret[i-2] + ret[i-1]
	}
	return ret[n]
}
