package main

import "fmt"

func repeatAndMissing(nums []int) (int, int) {

	total := len(nums) * (len(nums) + 1) / 2
	var repeat int
	for _, num := range nums {
		total -= num
		if repeat == num {
			continue
		}
		repeat = num
	}
	return repeat, total + repeat
}

func main() {

	nums := []int{3, 1, 2, 5, 3}
	fmt.Println(repeatAndMissing(nums)) // Expected output: 3,4
	nums = []int{3, 1, 2, 5, 4, 6, 7, 5}
	fmt.Println(repeatAndMissing(nums)) // Expected output: 5,8
}
