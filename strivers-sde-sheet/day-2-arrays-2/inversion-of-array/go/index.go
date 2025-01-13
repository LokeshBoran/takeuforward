package main

import "fmt"

func getInversions(nums []int, n int) int {

	var inversions int = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				inversions++
			}
		}
	}

	return inversions
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(getInversions(nums, 5)) // Expected output: 0
	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(getInversions(nums, 5)) // Expected output: 10
	nums = []int{5, 3, 2, 1, 4}
	fmt.Println(getInversions(nums, 5)) // Expected output: 7
	nums = []int{1, 3, 4, 2, 2}
	fmt.Println(getInversions(nums, 5)) // Expected output: 4
	nums = []int{-2, -3, -4}
	fmt.Println(getInversions(nums, 3)) // Expected output: 3
}
