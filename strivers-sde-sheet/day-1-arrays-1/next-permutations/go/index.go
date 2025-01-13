package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 1. Find the first decreasing element from the right
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i < 0 {
		// Array is in descending order, reverse it to get the smallest permutation
		reverse(nums)
		return
	}

	// 2. Find the smallest element to the right of nums[i] that is greater than nums[i]
	j := n - 1
	for j > i && nums[j] <= nums[i] {
		j--
	}

	// 3. Swap nums[i] and nums[j]
	nums[i], nums[j] = nums[j], nums[i]

	// 4. Reverse the subarray to the right of index i
	reverse(nums[i+1:])

}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums) // Expected output: [1,3,2]
	fmt.Println(nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums) // Expected output: [1,2,3]
	fmt.Println(nums)
	nums = []int{1, 1, 5}
	nextPermutation(nums) // Expected output: [1,5,1]
	fmt.Println(nums)

}
