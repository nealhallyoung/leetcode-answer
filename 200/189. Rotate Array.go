// main.go
package main

import "fmt"

/*
nums = "----->-->"; k =3
result = "-->----->";

reverse "----->-->" we can get "<--<-----"
reverse "<--" we can get "--><-----"
reverse "<-----" we can get "-->----->"
this visualization help me figure it out :)
*/

func main() {
	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3}
	rotate(nums, 5)
	fmt.Println("n", nums)
}

func rotate(nums []int, k int) {
	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(array []int, a int, b int) {
	for i, j := a, b; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
