// main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	val := 3
	k := removeElement(nums, val)
	fmt.Println("k:", k)
	fmt.Println(nums)

}
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 101
			k++
		}
	}
	sort.Ints(nums)
	return len(nums) - k
}
