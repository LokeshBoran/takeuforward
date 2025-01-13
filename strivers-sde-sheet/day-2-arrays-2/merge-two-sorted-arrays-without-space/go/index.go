package main

import (
	"fmt"
	"sort"
)

func mergeTwoSortedArraysWithoutExtraSpace(nums1 []int, nums2 []int) {
	m, n := len(nums1), len(nums2)
	i, j := m-1, 0
	for i >= 0 && j < n {
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
		}
		i--
		j++
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(nums1, nums2)
}

func main() {
	nums1 := []int{1, 8, 8}
	nums2 := []int{2, 3, 4, 5}
	mergeTwoSortedArraysWithoutExtraSpace(nums1, nums2) // Expected output: [1,2,3] [4,5,8,8]
	nums1 = []int{1, 1, 1, 1}
	nums2 = []int{2, 2, 3, 3, 5}
	mergeTwoSortedArraysWithoutExtraSpace(nums1, nums2) // Expected output: [1,1,1,1] [2,2,3,3,5]

}
