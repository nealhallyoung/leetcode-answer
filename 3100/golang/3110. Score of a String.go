func scoreOfString(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i-1]) - int(s[i]))
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}