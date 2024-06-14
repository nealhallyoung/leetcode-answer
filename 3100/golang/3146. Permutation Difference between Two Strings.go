func findPermutationDifference(s string, t string) int {
	res := 0
	for i, cs := range s {
		for j, ct := range t {
			if cs == ct {
				res += abs(i - j)
			}
		}
	}
	return res

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
