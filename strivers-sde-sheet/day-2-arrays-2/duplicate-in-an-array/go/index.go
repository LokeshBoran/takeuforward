package main

import (
	"fmt"
	"sort"
)

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if len(nums) > i+1 && nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1

}
func main() {

	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums)) // Expected output: 2
	nums = []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums)) // Expected output: 3
	nums = []int{1, 1}
	fmt.Println(findDuplicate(nums)) // Expected output: 1
	nums = []int{3, 3, 3, 3, 3}
	fmt.Println(findDuplicate(nums)) // Expected output: 3
}
