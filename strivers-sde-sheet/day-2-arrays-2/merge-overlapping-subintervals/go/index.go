package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	/* Sort the intervals */
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	returnInterals := [][]int{}
	/* Merge intervals */
	i := 0
	for i < len(intervals) {
		if len(returnInterals) > 0 && returnInterals[len(returnInterals)-1][1] >= intervals[i][0] {
			returnInterals[len(returnInterals)-1][1] = max(returnInterals[len(returnInterals)-1][1], intervals[i][1])
		} else {
			returnInterals = append(returnInterals, intervals[i])
		}
		i++

	}

	return returnInterals
}

func main() {

	intervals := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // Expected output: [[1,6],[8,10],[15,18]]
	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals)) // Expected output: [[1,5]]
	intervals = [][]int{{1, 4}, {0, 4}}
	fmt.Println(merge(intervals)) // Expected output: [[0,4]]

}
