// main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	k := removeDuplicates(nums)
	fmt.Println("k:", k)

}

func removeDuplicates(nums []int) int {
	// the length of nums equal 1
	if len(nums) == 1 {
		return 1
	}

	i := 0
	j := 1
	k := 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			nums[i] = 101
			k++
		}
		i++
		j++
	}
	sort.Ints(nums)
	return len(nums) - k
}
