func numJewelsInStones(jewels string, stones string) int {
	ret := 0
	hashtable_jewels := map[rune]bool{}
	for _, v := range jewels {
		hashtable_jewels[v] = true
	}
	for _, v := range stones {
		if hashtable_jewels[v] {
			ret++
		}
	}
	return ret
}
