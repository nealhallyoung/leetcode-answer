func maximumHappinessSum(happiness []int, k int) int64 {
	ans := 0
	sort.Ints(happiness)
	for i := 0; i < k; i++ {
		x := happiness[len(happiness)-i-1] - i
		ans += max(x, 0)
	}
	return int64(ans)
}