
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	result := make([]bool, len(candies))
	for _, v := range candies {
		if v >= max {
			max = v
		}
	}
	for k, v := range candies {
		if v+extraCandies >= max {
			result[k] = true
		} else {
			result[k] = false
		}
	}
	return result
}