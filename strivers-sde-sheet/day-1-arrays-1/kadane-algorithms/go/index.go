package main

import "fmt"

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums)) // Expected output: 6
	nums = []int{1}
	fmt.Println(maxSubArray(nums)) // Expected output: 1
	nums = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums)) // Expected output: 23
	nums = []int{-2, -3, -4}
	fmt.Println(maxSubArray(nums)) // Expected output: -2
}
