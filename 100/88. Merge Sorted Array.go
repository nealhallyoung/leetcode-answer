// main.go
package main

import "sort"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)

}
func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m
	// loop through nums2
	for j := 0; j < n; j++ {
		nums1[i] = nums2[j]
		i++
	}
	sort.Ints(nums1)

}
