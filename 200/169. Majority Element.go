// main.go
package main

import "fmt"

/*
Boyer–Moore majority vote algorithm

最终选出的数，票数过半就能用··

*/

func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3}
	fmt.Println("num: ", majorityElement(nums))

}

func majorityElement(nums []int) int {
	major := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count++
			continue
		}
		if major != nums[i] {
			count--
			continue
		}
		count++
	}
	return major
}
