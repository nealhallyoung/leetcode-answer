func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0
	for {
		mid := (right + left) / 2.0
		i, count := -1, 0
		x, y := 0, 1
		for j := 1; j < n; j++ {
			for mid*float64(arr[j]) > float64(arr[i+1]) {
				i++
				if y*arr[i] > arr[j]*x {
					x, y = arr[i], arr[j]
				}
			}
			count += i + 1
		}
		if count == k {
			return []int{x, y}
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}
}