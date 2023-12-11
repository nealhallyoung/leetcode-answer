// main.go
package main

import (
	"fmt"
)

/*
|      |      | i    |      |      |      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1    | 1    | 1    | 1    | 2    | 2    | 2    | 2    | 3    | 3    | 3    |
|      |      | j    |      |      |      |      |      |      |      |      |

Nuts[0:i] 处理好的数组

i 是该数组的长度
*/

func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3}
	k := removeDuplicates(nums)
	fmt.Println("k:", k)
	fmt.Println(nums[0:k])

}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i := 2
	j := 2
	for j < n {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
