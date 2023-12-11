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
	p := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(p))
}
func maxProfit(prices []int) int {
	ans := 0
	minPrice := prices[0]
	for _, v := range prices {
		ans = max(ans, v-minPrice)
		minPrice = min(minPrice, v)
	}
	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
