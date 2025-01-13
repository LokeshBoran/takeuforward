package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
	fmt.Println(nums1)

}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3) // Expected output: [1,2,2,3,5,6]
	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0) // Expected output: [1]
}
