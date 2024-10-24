func longestPalindrome(s string) string {
	maxlen := 1 // 最长回文子串长度
	ri, rj := 0, 0
	// 定义二维表
	ret := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = make([]int, len(s))
	}
	// 遍历生成结果
	for i := len(s) - 1; i >= 0; i-- {
		// 按照对角线遍历
		for j := i; j < len(s); j++ {
			if s[i] != s[j] { // 如果不相等，一定不是回文串
				ret[i][j] = 0
				continue
			}
			res := 1
			// 判断 s(i+1,j-1) 是否存在，如果存在，那么看其是不是回文串
			if i+1 <= j-1 {
				res = ret[i+1][j-1]
			}
			ret[i][j] = res
			if ret[i][j] > 0 && maxlen < j-i+1 {
				maxlen = j - i + 1
				ri, rj = i, j
			}
		}
	}
	return s[ri : rj+1]
}
