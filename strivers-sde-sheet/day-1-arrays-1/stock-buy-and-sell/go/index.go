package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, p := range prices {
		minPrice = min(minPrice, p)
		maxProfit = max(maxProfit, p-minPrice)
	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // Expected output: 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // Expected output: 0
}
