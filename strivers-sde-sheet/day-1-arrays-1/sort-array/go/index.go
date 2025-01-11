package main

import "fmt"

func sortColors(nums []int) {

	noOfZeroes := 0
	noOfOnes := 0
	noOfTwos := 0
	for _, ele := range nums {
		if ele == 0 {
			noOfZeroes++
		} else if ele == 1 {
			noOfOnes++
		} else {
			noOfTwos++
		}

	}
	for i := 0; i < len(nums); i++ {
		if i < noOfZeroes {
			nums[i] = 0
		} else if i < noOfZeroes+noOfOnes {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	nums = []int{2, 0, 1}
	sortColors(nums)
}
